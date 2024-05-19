package handler

import (
	"net/http"
	"os"
	"regexp"

	"github.com/go-openapi/runtime/middleware"
)

type SwaggerHandler interface {
	File(prefix string) http.HandlerFunc
	Docs(prefix string) http.HandlerFunc
}

type swaggerHandler struct {
}

func NewSwaggerHandler() SwaggerHandler {
	return &swaggerHandler{}
}

func (p *swaggerHandler) File(prefix string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fileBytes, err := os.ReadFile("swagger.yaml")
		if err != nil {
			panic(err)
		}

		regex, _ := regexp.Compile(`^basePath\s*:\s+.*`)
		fileBytes = regex.ReplaceAll(fileBytes, []byte("basePath: "+prefix))

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/yaml")
		_, _ = w.Write(fileBytes)
	}
}

func (p *swaggerHandler) Docs(prefix string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		opts := middleware.SwaggerUIOpts{SpecURL: "swagger.yaml", BasePath: prefix}
		h := middleware.SwaggerUI(opts, nil)
		h.ServeHTTP(w, r)
	}
}
