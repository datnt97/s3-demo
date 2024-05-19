package httpc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"tronglv_upload_svc/helper/define"

	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/rest/httpc"
)

const DefaultTimeout int = 10

type (
	Option func(r *http.Request) *http.Request

	Service interface {
		Head(ctx context.Context, url string, opts ...Option) (*http.Response, error)
		Get(ctx context.Context, url string, opts ...Option) (*http.Response, error)
		Post(ctx context.Context, url string, payload any, opts ...Option) (*http.Response, error)
		Put(ctx context.Context, url string, payload any, opts ...Option) (*http.Response, error)
		Patch(ctx context.Context, url string, payload any, opts ...Option) (*http.Response, error)
		Delete(ctx context.Context, url string, payload any, opts ...Option) (*http.Response, error)
		PostForm(ctx context.Context, url string, payload map[string]string, opts ...Option) (*http.Response, error)
		httpc.Service
	}

	clientService struct {
		name string
		svc  httpc.Service
		cli  *http.Client
		opts []Option
	}
)

func Client(timeout int) *http.Client {
	return &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
}

func WithFormUrlEncode() Option {
	return func(r *http.Request) *http.Request {
		r.Header.Set(ContentType, FormUrlEncodeContentType)
		return r
	}
}

func WithJsonContentType() Option {
	return func(r *http.Request) *http.Request {
		r.Header.Set(ContentType, ApplicationJson)
		return r
	}
}

func WithHeaders(headers map[string]string) Option {
	return func(r *http.Request) *http.Request {
		for k, v := range headers {
			r.Header.Set(k, v)
		}
		return r
	}
}

func WithAuthToken(token string) Option {
	return func(r *http.Request) *http.Request {
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		return r
	}
}

func WithBasicAuth(user, pass string) Option {
	return func(r *http.Request) *http.Request {
		r.SetBasicAuth(user, pass)
		return r
	}
}

func WithQueryParams(params map[string]string) Option {
	return func(r *http.Request) *http.Request {
		q := r.URL.Query()
		for k, v := range params {
			q.Add(k, v)
		}
		r.URL.RawQuery = q.Encode()
		return r
	}
}

func New(name string, opts ...Option) Service {
	return NewWithClient(name, Client(DefaultTimeout), opts...)
}

func NewWithClient(name string, cli *http.Client, opts ...Option) Service {
	return clientService{
		name: name,
		cli:  cli,
		opts: opts,
	}
}

func (s clientService) new(opts ...Option) httpc.Service {
	var reqOpts []httpc.Option
	for _, opt := range append(s.opts, opts...) {
		reqOpts = append(reqOpts, httpc.Option(opt))
	}
	return httpc.NewServiceWithClient(s.name, s.cli, reqOpts...)
}

func (s clientService) Do(ctx context.Context, method, url string, data any) (*http.Response, error) {
	return s.new().Do(ctx, method, url, data)
}

func (s clientService) DoRequest(r *http.Request) (*http.Response, error) {
	return s.new().DoRequest(r)
}

func (s clientService) Head(ctx context.Context, url string, opts ...Option) (*http.Response, error) {
	return s.new(opts...).Do(ctx, http.MethodHead, url, nil)
}

func (s clientService) Get(ctx context.Context, url string, opts ...Option) (*http.Response, error) {
	return s.new(opts...).Do(ctx, http.MethodGet, url, nil)
}

func (s clientService) Post(ctx context.Context, url string, payload any, opts ...Option) (*http.Response, error) {
	req, err := s.request(ctx, http.MethodPost, url, payload, opts...)
	if err != nil {
		return nil, err
	}
	return s.new().DoRequest(req)
}

func (s clientService) Put(ctx context.Context, url string, payload any, opts ...Option) (*http.Response, error) {
	req, err := s.request(ctx, http.MethodPut, url, payload, opts...)
	if err != nil {
		return nil, err
	}
	return s.new().DoRequest(req)
}

func (s clientService) Patch(ctx context.Context, url string, payload any, opts ...Option) (*http.Response, error) {
	req, err := s.request(ctx, http.MethodPatch, url, payload, opts...)
	if err != nil {
		return nil, err
	}
	return s.new().DoRequest(req)
}

func (s clientService) Delete(ctx context.Context, url string, payload any, opts ...Option) (*http.Response, error) {
	req, err := s.request(ctx, http.MethodDelete, url, payload, opts...)
	if err != nil {
		return nil, err
	}
	return s.new().DoRequest(req)
}

func (s clientService) PostForm(ctx context.Context, u string, payload map[string]string, opts ...Option) (*http.Response, error) {
	var f = url.Values{}
	for k, v := range payload {
		f.Add(k, v)
	}
	req, err := s.request(ctx, http.MethodPost, u, f, append(opts, WithFormUrlEncode())...)
	if err != nil {
		return nil, err
	}
	return s.new().DoRequest(req)
}

func (s clientService) request(ctx context.Context, method string, url string, payload any, opts ...Option) (*http.Request, error) {
	b, t, e := s.payload(payload)
	if e != nil {
		return nil, e
	}
	r, err := http.NewRequestWithContext(ctx, method, url, b)
	if err != nil {
		return nil, err
	}
	r.Header.Set(ContentType, t)
	for _, opt := range opts {
		r = opt(r)
	}
	return r, nil
}

func (s clientService) payload(payload any) (io.Reader, string, error) {
	if p, ok := payload.(url.Values); ok {
		return strings.NewReader(p.Encode()), FormUrlEncodeContentType, nil
	}
	if p, ok := payload.(io.Reader); ok {
		return p, FormUrlEncodeContentType, nil
	}
	if p, ok := payload.(string); ok {
		return bytes.NewReader([]byte(p)), ApplicationJson, nil
	}
	p, err := json.Marshal(payload)
	if err != nil {
		return nil, define.EmptyString, err
	}
	return bytes.NewReader(p), ApplicationJson, nil
}
