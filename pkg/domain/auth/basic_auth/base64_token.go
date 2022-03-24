package basic_auth

import (
	"encoding/base64"
	"strings"
)

type Base64Token struct {
	APIKey string
	Secret string
}

func DecodeBase64Token(token string) (*Base64Token, error) {
	decodeString, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}
	base64Token := strings.Split(string(decodeString), ":")
	return &Base64Token{
		APIKey: base64Token[0],
		Secret: base64Token[1],
	}, nil
}
