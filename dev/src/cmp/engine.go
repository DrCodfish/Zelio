package wasmtime

import (
	"runtime"
)

// Engine is a global context for compiling and managing WebAssembly modules.
type Engine struct {
	_ptr *C.wasmtime_engine_t
}

// NewEngine creates a new Engine using the default configuration.
func NewEngine() *Engine {
	ptr := C.wasmtime_engine_new()
	runtime.SetFinalizer(ptr, func(ptr *C.wasmtime_engine_t) {
		C.wasmtime_engine_delete(ptr)
	})
	return &Engine{_ptr: ptr}
}