package internal

func InputLength() uint64 {
	return extism_input_length()
}

//go:wasmimport env extism_input_length
func extism_input_length() uint64

type extismPointer uint64

type Pointer struct {
	Ptr extismPointer
}

func Ptr(i uint64) *Pointer {
	return &Pointer{extismPointer(i)}
}

func (p *Pointer) AddPtr(ptr *Pointer) *Pointer {
	point := p.Ptr + ptr.Ptr

	return &Pointer{point}
}

func (p *Pointer) Length() uint64 {
	return extism_length(p.Ptr)
}

//go:wasmimport env extism_length
func extism_length(extismPointer) uint64

func Alloc(i uint64) *Pointer {
	ptr := extism_alloc(i)

	return &Pointer{ptr}
}

//go:wasmimport env extism_alloc
func extism_alloc(uint64) extismPointer

func (p *Pointer) Free() {
	extism_free(p.Ptr)
}

//go:wasmimport env extism_free
func extism_free(extismPointer)

func (p *Pointer) InputLoadU8() uint8 {
	return extism_input_load_u8(p.Ptr)
}

//go:wasmimport env extism_input_load_u8
func extism_input_load_u8_(extismPointer) uint32

func extism_input_load_u8(p extismPointer) uint8 {
	return uint8(extism_input_load_u8_(p))
}

func (p *Pointer) InputLoadU64() uint64 {
	return extism_input_load_u64(p.Ptr)
}

//go:wasmimport env extism_input_load_u64
func extism_input_load_u64(extismPointer) uint64

func (p *Pointer) OutputSet(o uint64) {
	extism_output_set(p.Ptr, o)
}

//go:wasmimport env extism_output_set
func extism_output_set(extismPointer, uint64)

func (p *Pointer) ErrorSet() {
	extism_error_set(p.Ptr)
}

//go:wasmimport env extism_error_set
func extism_error_set(extismPointer)

func (p *Pointer) ConfigGet() *Pointer {
	ptr := extism_config_get(p.Ptr)

	return &Pointer{ptr}
}

//go:wasmimport env extism_config_get
func extism_config_get(extismPointer) extismPointer

func (p *Pointer) VarGet() *Pointer {
	ptr := extism_var_get(p.Ptr)

	return &Pointer{ptr}
}

//go:wasmimport env extism_var_get
func extism_var_get(extismPointer) extismPointer

func (p *Pointer) VarSet(v *Pointer) {
	if v != nil {
		extism_var_set(p.Ptr, v.Ptr)
		return
	}

	extism_var_set(p.Ptr, 0)
}

//go:wasmimport env extism_var_set
func extism_var_set(extismPointer, extismPointer)

func (p *Pointer) StoreU8(v uint8) {
	extism_store_u8(p.Ptr, v)
}

//go:wasmimport env extism_store_u8
func extism_store_u8_(extismPointer, uint32)
func extism_store_u8(p extismPointer, v uint8) {
	extism_store_u8_(p, uint32(v))
}

func (p *Pointer) LoadU8(v uint8) uint8 {
	return extism_load_u8(p.Ptr + extismPointer(v))
}

//go:wasmimport env extism_load_u8
func extism_load_u8_(extismPointer) uint32
func extism_load_u8(p extismPointer) uint8 {
	return uint8(extism_load_u8_(p))
}

func (p *Pointer) StoreU64(v uint64) {
	extism_store_u64(p.Ptr, v)
}

//go:wasmimport env extism_store_u64
func extism_store_u64(extismPointer, uint64)

func (p *Pointer) LoadU64(ptr *Pointer) uint64 {
	return extism_load_u64(p.Ptr + ptr.Ptr)
}

//go:wasmimport env extism_load_u64
func extism_load_u64(extismPointer) uint64

func (p *Pointer) HTTPRequest(ptr *Pointer) *Pointer {
	point := extism_http_request(p.Ptr, ptr.Ptr)

	return &Pointer{point}
}

//go:wasmimport env extism_http_request
func extism_http_request(extismPointer, extismPointer) extismPointer

func HTTPStatusCode() int32 {
	return extism_http_status_code()
}

//go:wasmimport env extism_http_status_code
func extism_http_status_code() int32

func (p *Pointer) LogInfo() {
	extism_log_info(p.Ptr)
}

//go:wasmimport env extism_log_info
func extism_log_info(extismPointer)

func (p *Pointer) LogDebug() {
	extism_log_debug(p.Ptr)
}

//go:wasmimport env extism_log_debug
func extism_log_debug(extismPointer)

func (p *Pointer) LogWarn() {
	extism_log_warn(p.Ptr)
}

//go:wasmimport env extism_log_warn
func extism_log_warn(extismPointer)

func (p *Pointer) LogError() {
	extism_log_error(p.Ptr)
}

//go:wasmimport env extism_log_error
func extism_log_error(extismPointer)
