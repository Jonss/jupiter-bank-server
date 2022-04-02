package db

import "context"

type AppClient struct {
	Name      string
	APIKey    string
	APISecret string
}

type FetchAppClientParams struct {
	APIKey    string
	APISecret string
}

const fetchAppClientQuery = `
	SELECT name, api_key, secret FROM app_clients
	WHERE api_key = $1 
	AND secret = $2 
`

func (q *Queries) FetchAppClient(ctx context.Context, arg FetchAppClientParams) (*AppClient, error) {
	row := q.db.QueryRowContext(ctx, fetchAppClientQuery, arg.APIKey, arg.APISecret)
	var ac AppClient
	err := row.Scan(
		&ac.Name,
		&ac.APIKey,
		&ac.APISecret)
	if err != nil {
		return nil, err
	}
	return &ac, nil
}
