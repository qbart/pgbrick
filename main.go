package main

import (
	"context"
	"os"

	"github.com/qbart/pgbrick/pgbrick"
)

func main() {
	driver := pgbrick.New()
	err := driver.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	defer driver.Close(context.Background())
}
