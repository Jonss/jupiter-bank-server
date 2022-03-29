package db

import (
	"context"
	"database/sql"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"testing"
)

func TestQueries_FetchAppClient(t *testing.T) {
	key := uuid.New().String()
	secret := uuid.New().String()
	appName := faker.Name()
	createAppClient(t, appName, key, secret)

	tests := []struct {
		name string
		arg  FetchAppClientParams
		want *AppClient
		err  error
	}{
		{
			name: "should fetch app client",
			arg: FetchAppClientParams{
				key, secret,
			},
			want: &AppClient{
				Name:      appName,
				APISecret: secret,
				APIKey:    key,
			},
		},
		{
			name: "should get error when app client does not exists",
			err:  sql.ErrNoRows,
			arg: FetchAppClientParams{
				"non-existing-key", "non-existing-secret",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := testQueries.FetchAppClient(context.Background(), tc.arg)
			if err != tc.err {
				t.Errorf("FetchAppClient() error = %v, wantErr %v", err, tc.err)
				return
			}
			if tc.err != nil {
				return
			}

			if got.Name != tc.want.Name {
				t.Errorf("FetchAppClient() got = %v, want %v", got.Name, tc.want.Name)
			}

			if got.APIKey != tc.want.APIKey {
				t.Errorf("FetchAppClient() got = %v, want %v", got.APIKey, tc.want.APIKey)
			}

			if got.APISecret != tc.want.APISecret {
				t.Errorf("FetchAppClient() got = %v, want %v", got.APISecret, tc.want.APISecret)
			}
		})
	}
}

func createAppClient(t *testing.T, appName, key, secret string) {
	insertQuery := `
	INSERT INTO app_clients(
		name,
		api_key,
		secret
	) VALUES (
		$1,
		$2,
		$3
	)`

	_, err := testQueries.db.ExecContext(context.Background(), insertQuery, appName, key, secret)
	if err != nil {
		t.Fatalf("unexpected error inserting existing app_client. error=(%v)", err)
	}
}
