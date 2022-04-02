package paseto_auth

import (
	"context"
	"github.com/Jonss/jupiter-bank-server/pkg/db"
	"strconv"
)

func (s service) Login(ctx context.Context, email, password string) (PasetoToken, error) {
	user, err := s.userService.FetchUserByEmailAndPassword(ctx, email, password)
	if err != nil {
		return PasetoToken{}, err
	}
	token, err := CreateToken(user.ID)
	if err != nil {
		return PasetoToken{}, err
	}
	return PasetoToken{
			SignedKey: token.SignedKey,
			PublicHex: token.PublicHex},
		nil
}

func (s service) VerifyUser(ctx context.Context, token, hex string) (*db.User, error) {
	userIDstr, err := DecryptToken(token, hex)
	if err != nil {
		return nil, err
	}
	userID, err := strconv.ParseUint(userIDstr, 10, 64)
	if err != nil {
		return nil, err
	}
	user, err := s.userService.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
