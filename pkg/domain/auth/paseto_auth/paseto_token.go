package paseto_auth

import (
	"aidanwoods.dev/go-paseto"
	"errors"
	"strconv"
	"time"
)

var (
	ErrExpiredToken = errors.New("token has expired")
)

type PasetoToken struct {
	SignedKey string
	PublicHex string
}

func CreateToken(userID uint64) (PasetoToken, error) {
	userIDstr := strconv.FormatUint(userID, 10)

	token := paseto.NewToken()
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(2 * time.Hour))
	token.SetString("user-id", userIDstr)
	token.SetSubject(userIDstr)

	secretKey := paseto.NewV4AsymmetricSecretKey()
	public := secretKey.Public()

	signed := token.V4Sign(secretKey, nil)
	return PasetoToken{SignedKey: signed, PublicHex: public.ExportHex()}, nil
}

func DecryptToken(signedKey, publicHex string) (string, error) {
	publicKey, err := paseto.NewV4AsymmetricPublicKeyFromHex(publicHex)
	if err != nil {
		return "", err
	}

	parser := paseto.NewParser()
	token, err := parser.ParseV4Public(publicKey, signedKey, nil)
	if err != nil {
		return "", err
	}
	expiration, err := token.GetExpiration()
	if err != nil {
		return "", err
	}
	if expiration.Before(time.Now()) {
		return "", ErrExpiredToken
	}
	subject, err := token.GetSubject()
	if err != nil {
		return "", ErrExpiredToken
	}
	return subject, err
}
