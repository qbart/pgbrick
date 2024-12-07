.PHONY: build
build:
	GOOS=wasip1 GOARCH=wasm tinygo build -scheduler=none -no-debug -target=wasi -o uppercase.wasm ./main.go
	mkdir -p ../../../pgbrick/testdata
	mv uppercase.wasm ../../../pgbrick/testdata/uppercase.wasm