package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var testDB *sql.DB
var testQueries *Queries

func TestMain(m *testing.M) {
	ctx := context.Background()
	container, err := postgresTestContainer(ctx)
	if err != nil {
		log.Fatal("error instantiating postgres container", err)
	}

	port, err := container.MappedPort(ctx, nat.Port("5432/tcp"))
	if err != nil {
		log.Fatal("error getting mapped port", err)
	}
	dbSource := fmt.Sprintf("postgres://postgres:password@localhost:%s/%s?sslmode=disable", port.Port(), "jupiterbank_test")

	testDB, err = NewConnection(dbSource)
	if err != nil {
		log.Fatal("error connecting postgres docker", err)
	}

	err = Migrate(testDB, "jupiterbank_test", "migrations")
	if err != nil {
		log.Fatal("error on migrate:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}

func postgresTestContainer(ctx context.Context) (testcontainers.Container, error) {
	var env = map[string]string{
		"POSTGRES_PASSWORD": "password",
		"POSTGRES_USER":     "postgres",
		"POSTGRES_DB":       "jupiterbank_test",
	}
	var port = "5432/tcp"
	dbURL := func(port nat.Port) string {
		return fmt.Sprintf("postgres://postgres:password@localhost:%s/%s?sslmode=disable", port.Port(), "jupiterbank_test")
	}

	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres:12.3-alpine",
			ExposedPorts: []string{port},
			Cmd:          []string{"postgres", "-c", "fsync=off"},
			Env:          env,
			WaitingFor:   wait.ForSQL(nat.Port(port), "postgres", dbURL).Timeout(time.Second * 5),
		},
		Started: true,
	}

	container, err := testcontainers.GenericContainer(ctx, req)
	if err != nil {
		return container, fmt.Errorf("failed to start container: %s", err)
	}
	return container, nil
}
