package pgbrick

import (
	"context"
	"testing"
)

func TestPing(t *testing.T) {
	withConnection(t, func(driver *Driver) {
		if err := driver.Ping(context.Background()); err != nil {
			t.Error(err)
		}
	})
}
