package db

import (
	"context"
	"database/sql"
	"testing"
)

func TestGetConfigByKey(t *testing.T) {
	insertQuery := `INSERT INTO configs (key, value) VALUES ('my_key', 'my_value');`
	_, err := testQueries.db.ExecContext(context.Background(), insertQuery)
	if err != nil {
		t.Fatalf("unexpected error inserting config. error=(%v)", err)
	}

	testCases := []struct {
		name           string
		expectedErr    error
		expectedConfig Config
		key            string
	}{
		{
			name:           "test existing key",
			expectedErr:    nil,
			key:            "my_key",
			expectedConfig: Config{Key: "my_key", Value: "my_value"},
		},
		{
			name:        "test non existing key",
			expectedErr: sql.ErrNoRows,
			key:         "non_exist_key",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := testQueries.GetConfigByKey(context.Background(), GetConfigByKeyParams{tc.key})

			if err != tc.expectedErr {
				t.Fatalf("unexpected error. want %v got %v", err, tc.expectedErr)
			}

			if got.Key != tc.expectedConfig.Key {
				t.Errorf("Config.Key: want %s got %s", got.Key, tc.expectedConfig.Key)
			}

			if got.Value != tc.expectedConfig.Value {
				t.Errorf("Config.Key: want %s got %s", got.Value, tc.expectedConfig.Value)
			}
		})
	}
}
