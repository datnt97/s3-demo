package logify

import (
	"context"
	"fmt"
	"os"
	"time"
	"tronglv_upload_svc/helper/define"
	"tronglv_upload_svc/helper/errors"
	"tronglv_upload_svc/helper/httpc"
	"tronglv_upload_svc/helper/util"

	"go.opentelemetry.io/otel/trace"
)

type CollectorRequest struct {
	Level      string    `json:"level"`
	Type       string    `json:"type"`
	Service    string    `json:"service"`
	TraceID    string    `json:"trace_id"`
	StatusCode int       `json:"status_code"`
	Method     string    `json:"method"`
	Path       string    `json:"path"`
	Message    string    `json:"message"`
	StackTrace string    `json:"stack_trace"`
	DateTime   time.Time `json:"time"`
}

type Report struct {
	serviceName string
	gatewayUrl  string
	token       string
	restClient  httpc.Service
}

func NewReport() *Report {
	return &Report{
		serviceName: os.Getenv("SERVICE_NAME"),
		gatewayUrl:  os.Getenv("INAPP_GATEWAY_URL"),
		token:       os.Getenv("CENTRAL_LOG_TOKEN"),
		restClient:  httpc.New(os.Getenv("SERVICE_NAME")),
	}
}

func (r *Report) isAllowed(err errors.Error) bool {
	if len(r.token) > 0 && err.HasReport() {
		return true
	}
	return false
}

func (r *Report) Send(ctx context.Context, err errors.Error) {
	if !r.isAllowed(err) {
		return
	}

	var method, path string
	if attrs, ok := ctx.Value(define.SvcAttributesKey).(map[string]string); ok {
		if v, ok := attrs["method"]; ok {
			method = v
		}
		if v, ok := attrs["path"]; ok {
			path = v
		}
	}

	span := trace.SpanFromContext(ctx)
	req := CollectorRequest{
		Level:      "error",
		Service:    r.serviceName,
		Type:       fmt.Sprintf("%T", err.GetCause()),
		TraceID:    span.SpanContext().TraceID().String(),
		StatusCode: err.GetCode(),
		Method:     method,
		Path:       path,
		Message:    err.GetCause().Error(),
		StackTrace: err.GetMeta(errors.StackKey),
		DateTime:   util.TimeNow(),
	}

	go func() {
		_, e := r.restClient.Post(
			context.Background(),
			fmt.Sprintf("%s/central-log-svc/api/v1/collectors", r.gatewayUrl),
			req,
			httpc.WithHeaders(map[string]string{
				"Authorization": r.token,
			}),
		)
		if e != nil {
			New().ErrorCtx(ctx, e)
		}
	}()
}
