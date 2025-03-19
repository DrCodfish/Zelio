package wasmtime

import "runtime"

// Linker is used to link WebAssembly modules together.
type Linker struct {
	_ptr *C.wasmtime_linker_t
}

// NewLinker creates a new linker.
func NewLinker(store *Store) *Linker {
	ptr := C.wasmtime_linker_new(store._ptr())
	runtime.SetFinalizer(ptr, func(ptr *C.wasmtime_linker_t) {
		C.wasmtime_linker_delete(ptr)
	})
	return &Linker{_ptr: ptr}
}