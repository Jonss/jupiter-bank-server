package paseto_auth

import (
	"context"
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

func (s service) VerifyUser(ctx context.Context, token, hex string) error {
	userIDstr, err := DecryptToken(token, hex)
	if err != nil {
		return err
	}
	userID, err := strconv.ParseUint(userIDstr, 10, 64)
	if err != nil {
		return err
	}
	_, err = s.userService.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}
