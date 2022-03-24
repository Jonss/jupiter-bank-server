package basic_auth

import (
	"context"
	"github.com/Jonss/jupiter-bank-server/pkg/db"
)

type FetchAppClientParams struct {
	Token string
}

func (d *service) FetchAppClient(ctx context.Context, arg FetchAppClientParams) (bool, error) {
	token, err := DecodeBase64Token(arg.Token)
	if err != nil {
		return false, err
	}

	client, err := d.queries.FetchAppClient(ctx, db.FetchAppClientParams{APIKey: token.APIKey, APISecret: token.Secret})
	if err != nil {
		return false, err
	}

	return client != nil, nil
}
