package errors

import (
	"fmt"
	"tronglv_upload_svc/helper/locale"

	"gorm.io/gorm"
)

type Option func(f *errorSvc)

type Status struct {
	Code     int               `json:"code,omitempty"`
	Reason   string            `json:"reason,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

type errorSvc struct {
	Status
	cause  error
	report bool
}

func (e *errorSvc) Error() string {
	return fmt.Sprintf("error: code = %d reason = %s metadata = %v cause = %v", e.Code, e.Reason, e.Metadata, e.cause)
}

func (e *errorSvc) GetCause() error {
	return e.cause
}

func (e *errorSvc) GetCode() int {
	return e.Code
}

func (e *errorSvc) GetReason() string {
	if len(e.Reason) > 0 {
		return e.Reason
	}
	if e.cause != nil {
		return e.cause.Error()
	}
	return DefaultMsg
}

func (e *errorSvc) GetMetaData() map[string]string {
	return e.Metadata
}

func (e *errorSvc) GetMeta(key string) string {
	if v, ok := e.Metadata[key]; ok {
		return v
	}
	return ""
}

func (e *errorSvc) GetMetaCode() string {
	key := e.GetMeta(CodeKey)
	if len(key) > 0 {
		return key
	}
	return fmt.Sprintf("%d", e.Code)
}

func (e *errorSvc) HasReport() bool {
	return e.report
}

func From(err error) Error {
	if err == nil {
		return InternalServerReason(locale.FailedMsg.Message)
	}
	if Is(err, gorm.ErrRecordNotFound) {
		return DataNotFound()
	}
	if e := IsError(err); e != nil {
		return e
	}
	return InternalServer(err)
}

func IsError(err error) Error {
	var r Error
	if As(err, &r) {
		return r
	}
	return nil
}

func ToError(result any) error {
	var err error
	switch e := result.(type) {
	default:
		err = fmt.Errorf("%s", result)
	case error:
		err = e
	}
	return err
}

func WithStack(stack string) Option {
	return func(f *errorSvc) {
		f.Metadata[StackKey] = stack
	}
}

func WithReport() Option {
	return func(f *errorSvc) {
		f.report = true
	}
}

func WithMetas(keyvals ...string) Option {
	return func(f *errorSvc) {
		var metas = make(map[string]string)
		for i := 0; i < len(keyvals); i = i + 2 {
			if (i + 1) < len(keyvals) {
				metas[keyvals[i]] = keyvals[i+1]
			}
		}
		f.Metadata = metas
	}
}
