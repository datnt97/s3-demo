package logify

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type Logger interface {
	Error(...any)
	ErrorCtx(ctx context.Context, v ...any)
	Errorf(string, ...any)
	Errorw(string, ...logx.LogField)
	InfoCtx(ctx context.Context, v ...any)
	Info(...any)
	Infof(string, ...any)
	Infow(string, ...logx.LogField)
	WithCallerSkip(skip int) Logger
	WithContext(ctx context.Context) Logger
	WithDuration(d time.Duration) Logger
	WithFields(fields ...logx.LogField) Logger
	Printf(msg string, attrs ...any)
}
