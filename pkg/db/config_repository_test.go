package db

import (
	"context"
	"testing"
)

func TestGetConfigByKey(t *testing.T) {
	insertQuery := `INSERT INTO configs (key, value) VALUES ('my_key', 'my_value');`
	_, err := testQueries.db.ExecContext(context.Background(), insertQuery)
	if err != nil {
		t.Fatalf("unexpected error inserting config. error=(%v)", err)
	}

	got, err := testQueries.GetConfigByKey(context.Background(), GetConfigByKeyParams{"my_key"})
	if err != nil {
		t.Fatalf("unexpected error fetching config. error=(%v)", err)
	}

	want := Config{Key: "my_key", Value: "my_value"}

	if got.Key != want.Key {
		t.Errorf("Config.Key: want %s got %s", got.Key, want.Key)
	}

	if got.Value != want.Value {
		t.Errorf("Config.Key: want %s got %s", got.Value, want.Value)
	}
}
