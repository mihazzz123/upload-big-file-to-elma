package syncutils

import (
	"context"
	"time"
)

// Promise обёртка над OnceWithReset для удобного конкурентного кеширования
type Promise struct {
	once   OnceWithReset
	fn     func(context.Context) (interface{}, error)
	result interface{}
	err    error
	ttl    time.Duration
}

// PromiseOption опции промиса
type PromiseOption func(*Promise)

// WithTTL добавить время жизни закешированного результата
func WithTTL(ttl time.Duration) PromiseOption {
	return func(p *Promise) {
		p.ttl = ttl
	}
}

// NewPromise конструктор
func NewPromise(fn func(context.Context) (interface{}, error), opts ...PromiseOption) *Promise {
	p := &Promise{
		fn: fn,
	}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

// Await получить результат промиса
func (p *Promise) Await(ctx context.Context) (interface{}, error) {
	p.once.Do(func() {
		p.result, p.err = p.fn(ctx)
		if p.ttl > 0 {
			time.AfterFunc(p.ttl, p.once.Reset)
		}
	})
	if p.err != nil {
		p.once.Reset()
	}
	return p.result, p.err
}

// Reset сбросить промис, чтобы в следующий раз он разрешился заново
func (p *Promise) Reset() {
	p.once.Reset()
}
