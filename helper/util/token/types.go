package token

import (
	"fmt"
	"tronglv_upload_svc/helper/define"

	"github.com/golang-jwt/jwt"
)

type MapClaims struct {
	jwt.MapClaims
}

func (m MapClaims) GetInt(claimName string) int {
	value, ok := m.MapClaims[claimName].(int)
	if !ok {
		return 0
	}
	return value
}

func (m MapClaims) GetString(claimName string) string {
	value, ok := m.MapClaims[claimName].(string)
	if !ok {
		return define.EmptyString
	}
	return value
}

func (m MapClaims) GetSliceString(claimName string) []string {
	values, ok := m.MapClaims[claimName].([]any)
	if !ok {
		return nil
	}
	var results []string
	for _, v := range values {
		results = append(results, fmt.Sprintf("%v", v))
	}
	return results
}

func (m MapClaims) GetMap(claimName string) map[string]string {
	value, ok := m.MapClaims[claimName].(map[string]string)
	if !ok {
		return nil
	}
	return value
}
