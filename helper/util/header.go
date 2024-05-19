package util

import (
	"net/http"
	"strings"
)

func StripBearerFromToken(tok string) string {
	if len(tok) > 6 && strings.ToUpper(tok[0:7]) == "BEARER " {
		return tok[7:]
	}
	return tok
}

func GetTokenFromHeader(r *http.Request) string {
	return StripBearerFromToken(r.Header.Get("Authorization"))
}
