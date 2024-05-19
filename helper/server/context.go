package server

import (
	"context"
)

type (
	CtxOption struct {
		vals map[string]any
	}

	ContextOption func(s *CtxOption)
)

func WithValue(key string, val any) ContextOption {
	return func(r *CtxOption) {
		r.vals[key] = val
	}
}

type contextSvc struct {
	*CtxOption
}

func NewContext(opts ...ContextOption) context.Context {
	r := &contextSvc{
		CtxOption: &CtxOption{},
	}
	for _, opt := range opts {
		opt(r.CtxOption)
	}
	return r.Build()
}

func (r *contextSvc) Build() context.Context {
	ctx := context.Background()
	for k, v := range r.vals {
		ctx = context.WithValue(ctx, k, v)
	}
	return ctx
}
