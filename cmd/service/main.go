package main

import(
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"git.elewise.com/elma365/upload-big-file-elma365/internal/adaptor"
	"git.elewise.com/elma365/upload-big-file-elma365/internal/config"
	"git.elewise.com/elma365/upload-big-file-elma365/internal/service/http"
	"git.elewise.com/elma365/common/pkg/mw"
	"git.elewise.com/elma365/common/pkg/server"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
)

//nolint:golint
var (
	// BRANCH ветка
	BRANCH = "unknown"
	// BUILD_DATE дата
	BUILD_DATE = "unknown"
	// BUILD_HOST хост
	BUILD_HOST = "unknown"
	// COMMIT git commit
	COMMIT = "unknown"
)

func main() {
	cfg, err := config.New("upload-big-file-elma365")
	if err != nil {
		fmt.Println(err.Error())

		os.Exit(1)
	}

	zap.L().Info("version",
		zap.String("branch", BRANCH),
		zap.String("buildDate", BUILD_DATE),
		zap.String("buildHost", BUILD_HOST),
		zap.String("commit", COMMIT),
	)

	di, err := adaptor.NewDIContainer(cfg)
	if err != nil {
		zap.L().Fatal("di", zap.Error(err))
	}

	ctx := ctxzap.ToContext(context.Background(), zap.L())
	httpServer := mw.NewHTTPServer(cfg.Config, http.NewService(di), )

	srv := server.New(nil, httpServer)

	zap.L().Info("starting service")
	err = srv.Serve(cfg.Config)

	if err != nil {
		zap.L().Fatal("server", zap.Error(err))
	}

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	<-c
	zap.L().Warn("stopping service")
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	go func() {
		<-c
		zap.L().Warn("force stopping service")
		cancel()
	}()
	srv.Stop(ctx)

	zap.L().Info("stopped")
}
