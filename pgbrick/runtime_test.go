package pgbrick

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestRuntimeInsert(t *testing.T) {
	ctx := context.Background()
	rt := NewRuntime(ctx)
	defer rt.Close(ctx)

	wasm, err := os.ReadFile("testdata/uppercase.wasm")
	if err != nil {
		t.Fatal(err)
	}
	err = rt.Register(ctx, "test", wasm)
	if err == nil {
		instance := rt.instances["test"]
		data := []byte(`{"name":"hello world"}`)
		result, err := instance.BeforeInsert(ctx, data)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(string(result))
	} else {
		t.Error(err)
	}

}
