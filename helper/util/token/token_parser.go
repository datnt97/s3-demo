package token

import (
	"encoding/base64"
	"fmt"
	"tronglv_upload_svc/helper/logify"

	"github.com/golang-jwt/jwt"
)

type (
	// ParseOption defines the method to customize a TokenParser.
	ParseOption func(parser *TokenParser)

	// A TokenParser is used to parse tokens.
	TokenParser struct {
		publicKey []byte
		secretKey []byte
	}
)

func NewTokenParser(opts ...ParseOption) *TokenParser {
	parser := &TokenParser{}
	for _, opt := range opts {
		opt(parser)
	}
	return parser
}

func WithPublicKey(key string) ParseOption {
	return func(parser *TokenParser) {
		publicKey, err := base64.StdEncoding.DecodeString(key)
		if err != nil {
			logify.New().Error(err)
		}
		parser.publicKey = publicKey
	}
}

func (tp *TokenParser) Parse(token string) (*MapClaims, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if tp.publicKey != nil {
			return jwt.ParseRSAPublicKeyFromPEM(tp.publicKey)
		}
		return tp.secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !t.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return &MapClaims{t.Claims.(jwt.MapClaims)}, nil
}
