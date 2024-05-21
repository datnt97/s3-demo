package util

import (
	"regexp"
	"strings"
	"tronglv_upload_svc/helper/util/unicode"
)

func Slug(input string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	s := reg.ReplaceAllString(unicode.ToLatin(strings.ToLower(input)), " ")
	s = strings.TrimSpace(s)
	return strings.ToLower(strings.ReplaceAll(s, " ", "-"))
}
