package pgbrick

import (
	"context"
	"fmt"
	"testing"

	"github.com/docker/docker/api/types/container"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func withConnection(t *testing.T, fn func(driver *Driver)) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:      "postgres:16.6-bookworm",
		WaitingFor: wait.ForListeningPort("5432/tcp"),
		Env: map[string]string{
			"POSTGRES_USER":     "pgbrick",
			"POSTGRES_PASSWORD": "pgbrick",
			"POSTGRES_DB":       "pgbrick",
		},
		HostConfigModifier: func(hc *container.HostConfig) {
			hc.AutoRemove = true
		},
	}
	pgC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer testcontainers.CleanupContainer(t, pgC)

	endpoint, err := pgC.Endpoint(ctx, "")
	if err != nil {
		t.Fatal(err)
	}
	driver := New()
	err = driver.Connect(ctx, fmt.Sprintf("postgres://pgbrick:pgbrick@%s/pgbrick?sslmode=disable", endpoint))
	if err != nil {
		t.Log("Failed to connect to the database")
		t.Fatal(err)
	}
	defer driver.Close(ctx)

	fn(driver)
}
