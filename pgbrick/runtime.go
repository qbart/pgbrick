package pgbrick

import (
	"context"
	"fmt"
	"unsafe"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

type Runtime struct {
	modules     map[string]api.Module
	instances   map[string]*Instance
	wasmRuntime wazero.Runtime
}

type Instance struct {
	beforeInsert api.Function
}

func NewRuntime(ctx context.Context) *Runtime {
	rt := &Runtime{
		wasmRuntime: wazero.NewRuntime(ctx),
		modules:     make(map[string]api.Module),
		instances:   make(map[string]*Instance),
	}
	return rt
}

func (rt *Runtime) Close(ctx context.Context) {
	for _, module := range rt.modules {
		module.Close(ctx)
	}
	rt.modules = nil
	rt.instances = nil
}

func (rt *Runtime) Register(ctx context.Context, name string, data []byte) error {
	module, err := rt.wasmRuntime.Instantiate(ctx, data)
	if err != nil {
		return fmt.Errorf("failed to instantiate module: %w", err)
	}
	rt.modules[name] = module

	instance := &Instance{}
	instance.beforeInsert = module.ExportedFunction("BeforeInsert")
	rt.instances[name] = instance

	return nil
}

func (i *Instance) BeforeInsert(ctx context.Context, in []byte) ([]byte, error) {
	if i.beforeInsert == nil {
		return nil, fmt.Errorf("beforeInsert not found")
	}
	ptr := unsafe.Pointer(&in)
	result, err := i.beforeInsert.Call(ctx, api.EncodeExternref(uintptr(ptr)))
	for _, v := range result {
		fmt.Printf("result: %v\n", v)
	}
	if err != nil {
		return nil, fmt.Errorf("beforeInsert failed: %w", err)
	}
	// // data, ok := result.([]byte)
	// if !ok {
	// 	return nil, fmt.Errorf("beforeInsert returned unexpected type")
	// }
	// return data, nil
	return nil, fmt.Errorf("not implemented")
}
