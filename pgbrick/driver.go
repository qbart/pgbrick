package pgbrick

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Driver struct {
	conn *pgx.Conn
}

func New() *Driver {
	return &Driver{}
}

func (d *Driver) Connect(ctx context.Context, uri string) error {
	conn, err := pgx.Connect(ctx, uri)
	d.conn = conn
	return err
}

func (d *Driver) Ping(ctx context.Context) error {
	return d.conn.Ping(ctx)
}

func (d *Driver) Close(ctx context.Context) error {
	return d.conn.Close(ctx)
}
