package mw

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// DebugConfig конфигурация необходимая для переключения дебаг режима
type DebugConfig interface {
	IsDebug() bool
	EnableDebug()
	DisableDebug()
}

// DebugSwitcher сервис для переключения дебаг режима
type DebugSwitcher struct {
	sync.Mutex
	cfg    DebugConfig
	expire time.Time
	timer  *time.Timer
}

// NewDebugSwitcher конструктор
func NewDebugSwitcher(cfg DebugConfig) *DebugSwitcher {
	var timer *time.Timer
	if !cfg.IsDebug() {
		timer = time.AfterFunc(0, cfg.DisableDebug)
	}

	return &DebugSwitcher{
		cfg:    cfg,
		expire: time.Now().UTC(),
		timer:  timer,
	}
}

// ServeHTTP implements http.Handler
func (de *DebugSwitcher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if de.cfg.IsDebug() {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, "service already in debug mode")

		return
	}
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)

		if de.getExpire().Before(time.Now().UTC()) {
			_, _ = fmt.Fprint(w, "debug level off")
		} else {
			_, _ = fmt.Fprintf(w, "debug level enabled until %s", de.getExpire())
		}

		return
	}
	if r.Method != http.MethodPut {
		http.Error(w, "only GET and PUT method allowed", http.StatusMethodNotAllowed)
		return
	}

	dRaw := r.URL.Query().Get("duration")
	if dRaw == "" {
		de.setExpire(0)
		http.Error(w, "debug level off", http.StatusOK)
		return
	}

	d, err := time.ParseDuration(dRaw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if d > time.Hour {
		d = time.Hour
	}

	if de.getExpire().After(time.Now().Add(d)) {
		w.WriteHeader(http.StatusAlreadyReported)
		_, _ = fmt.Fprintf(w, "debug level is already enabled until %s", de.getExpire())
		return
	}

	de.setExpire(d)
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "debug level enabled until %s", de.getExpire())
}

func (de *DebugSwitcher) getExpire() time.Time {
	de.Lock()

	defer de.Unlock()

	return de.expire
}

func (de *DebugSwitcher) setExpire(d time.Duration) {
	de.Lock()

	defer de.Unlock()

	if !de.timer.Reset(d) {
		de.cfg.EnableDebug()
	}
	de.expire = time.Now().Add(d).UTC()
}
