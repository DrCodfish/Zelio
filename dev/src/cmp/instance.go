package wasmtime

// Instance represents an instantiated WebAssembly module.
type Instance struct {
	_ptr *C.wasmtime_instance_t
}

// NewInstance instantiates a WebAssembly module.
func NewInstance(store *Store, module *Module, imports []*Extern) (*Instance, error) {
	var instance Instance
	var trap *Trap
	trap = C.wasmtime_instance_new(
		store._ptr(),
		module._ptr(),
		(*C.wasmtime_extern_t)(unsafe.Pointer(&imports[0])),
		C.size_t(len(imports)),
		&instance._ptr,
	)
	if trap != nil {
		return nil, trap.toError()
	}
	runtime.SetFinalizer(&instance, func(i *Instance) {
		C.wasmtime_instance_delete(i._ptr)
	})
	return &instance, nil
}