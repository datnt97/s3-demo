package logify

import (
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/timex"
)

const (
	durationKey = "duration"
)

type logSvc struct {
	ctx        context.Context
	fields     []logx.LogField
	callerSkip int
}

func New() Logger {
	return &logSvc{
		callerSkip: 1,
	}
}

func (l *logSvc) writer() logx.Logger {
	log := logx.WithCallerSkip(l.callerSkip)
	if l.ctx != nil {
		log = log.WithContext(l.ctx)
	}
	if len(l.fields) > 0 {
		log = log.WithFields(l.fields...)
	}
	return log
}

func (l *logSvc) Error(v ...any) {
	l.writer().Error(v...)
}

func (l *logSvc) ErrorCtx(ctx context.Context, v ...any) {
	l.writer().WithContext(ctx).Error(v...)
}

func (l *logSvc) Errorf(format string, v ...any) {
	l.writer().Errorf(format, v...)
}

func (l *logSvc) Errorw(msg string, fields ...logx.LogField) {
	l.writer().Errorw(msg, fields...)
}

func (l *logSvc) Info(v ...any) {
	l.writer().Info(v...)
}

func (l *logSvc) InfoCtx(ctx context.Context, v ...any) {
	l.writer().WithContext(ctx).Info(v...)
}

func (l *logSvc) Infof(format string, v ...any) {
	l.writer().Infof(format, v...)
}

func (l *logSvc) Infow(msg string, fields ...logx.LogField) {
	l.writer().Infow(msg, fields...)
}

func (l *logSvc) WithCallerSkip(skip int) Logger {
	if skip <= 0 {
		return l
	}
	l.callerSkip = skip
	return l
}

func (l *logSvc) WithContext(ctx context.Context) Logger {
	l.ctx = ctx
	return l
}

func (l *logSvc) WithDuration(duration time.Duration) Logger {
	l.fields = append(l.fields, logx.Field(durationKey, timex.ReprOfDuration(duration)))
	return l
}

func (l *logSvc) WithFields(fields ...logx.LogField) Logger {
	l.fields = append(l.fields, fields...)
	return l
}

func (l *logSvc) Printf(msg string, attrs ...any) {
	l.Info(context.Background(), fmt.Sprintf(msg, attrs))
}
