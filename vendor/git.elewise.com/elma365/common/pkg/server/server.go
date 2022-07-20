package server

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/soheilhy/cmux"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

//nolint: gochecknoglobals // метрики должны быть глобальными
var (
	cmuxPortCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "elma365",
		Subsystem: "common",
		Name:      "cmux_port_share",
		Help:      "Count of services used cmux shared port for grpc and http services",
	}, []string{"service"})
)

// Server — это обёртка над gRPC и HTTP серверами для запуска их на одном порту и корректной остановки
type Server struct {
	m        cmux.CMux
	group    *sync.WaitGroup
	grpc     *grpc.Server
	http     *http.Server
	starting sync.RWMutex
	stop     int64
}

// New создаёт новую обёртку
func New(grpcServer *grpc.Server, httpServer *http.Server) *Server {
	return &Server{
		group: new(sync.WaitGroup),
		grpc:  grpcServer,
		http:  httpServer,
	}
}

// Config конфигурация необходимая для запуска сервиса
type Config interface {
	Name() string
	GetBind() string
	GetGRPCBind() string
	GetHTTPBind() string
	GetCMUXBind() bool
}

// Serve начинает слушать адрес из cfg.Bind, одновременно для обоих серверов в своих горутинах
func (srv *Server) Serve(cfg Config) error {
	if err := srv.init(cfg); err != nil {
		return err
	}
	if err := srv.grpcListen(cfg); err != nil {
		return err
	}
	if err := srv.httpListen(cfg); err != nil {
		return err
	}
	srv.listen()
	return nil
}

func (srv *Server) init(cfg Config) error {
	if cfg.GetCMUXBind() {
		cmuxPortCounter.WithLabelValues(cfg.Name()).Add(1)
		lis, err := net.Listen("tcp", cfg.GetBind())
		if err != nil {
			return errors.Wrapf(err, "fail to listen %q", cfg.GetBind())
		}
		srv.m = cmux.New(lis)
	}
	return nil
}

func (srv *Server) grpcListen(cfg Config) (err error) {
	if srv.grpc == nil {
		return nil
	}
	listener, err := srv.getListener(cfg.GetGRPCBind(), cmux.HTTP2HeaderField("content-type", "application/grpc"))
	if err != nil {
		return err
	}
	srv.group.Add(1)
	srv.starting.RLock()
	go func() {
		srv.starting.RUnlock()
		err := srv.grpc.Serve(listener)
		if err != nil {
			zap.L().Fatal("grpc serve fail", zap.Error(err))
		}
		srv.group.Done()
	}()
	return nil
}

func (srv *Server) httpListen(cfg Config) (err error) {
	if srv.http == nil {
		return nil
	}
	listener, err := srv.getListener(cfg.GetHTTPBind(), cmux.HTTP1Fast())
	if err != nil {
		return err
	}
	srv.group.Add(1)
	srv.starting.RLock()
	go func() {
		srv.starting.RUnlock()
		err := srv.http.Serve(listener)
		if err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("http serve fail", zap.Error(err))
		}
		srv.group.Done()
	}()
	return nil
}

func (srv *Server) getListener(bind string, matchers ...cmux.Matcher) (net.Listener, error) {
	if srv.m != nil {
		return srv.m.Match(matchers...), nil
	}
	listener, err := net.Listen("tcp", bind)
	return listener, errors.Wrapf(err, "fail to create listener on %q", bind)
}

func (srv *Server) listen() {
	if srv.m != nil {
		go func() {
			if err := srv.m.Serve(); err != nil && atomic.LoadInt64(&srv.stop) == 0 {
				zap.L().Error("cmux serve fail", zap.Error(err))
			}
		}()
	}
}

// Stop останавливает сервера с учётом завершения контекста
func (srv *Server) Stop(ctx context.Context) {
	srv.starting.Lock()
	atomic.StoreInt64(&srv.stop, 1)
	done := make(chan struct{})

	go func() {
		srv.group.Wait()
		close(done)
	}()

	if srv.grpc != nil {
		go srv.grpc.GracefulStop()
	}
	if srv.http != nil {
		go func() {
			_ = srv.http.Shutdown(ctx)
		}()
	}

	select {
	case <-ctx.Done():
		if srv.grpc != nil {
			srv.grpc.Stop()
		}
		if srv.http != nil {
			_ = srv.http.Close()
		}

	case <-done:
		// pass
	}
}

// WaitSig быстрый метод для сервисов в которых есть только gRPC сервис
//
// Слушает SIGINT и SIGTERM, после чего завершается. При повторном сигнале
// или по истечению таймаута делает force stop.
func (srv *Server) WaitSig(timeout time.Duration) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	<-signals

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	go func() {
		<-signals
		cancel()
	}()

	srv.Stop(ctx)
}
