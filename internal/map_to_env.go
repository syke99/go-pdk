package internal

func InputLength() uint64 {
	return extism_input_length()
}

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

func Alloc(i uint64) *Pointer {
	ptr := extism_alloc(i)

	return &Pointer{ptr}
}

func (p *Pointer) Free() {
	extism_free(p.Ptr)
}

func (p *Pointer) InputLoadU8() uint8 {
	return extism_input_load_u8(p.Ptr)
}

func (p *Pointer) InputLoadU64() uint64 {
	return extism_input_load_u64(p.Ptr)
}

func (p *Pointer) OutputSet(o uint64) {
	extism_output_set(p.Ptr, o)
}

func (p *Pointer) ErrorSet() {
	extism_error_set(p.Ptr)
}

func (p *Pointer) ConfigGet() *Pointer {
	ptr := extism_config_get(p.Ptr)

	return &Pointer{ptr}
}

func (p *Pointer) VarGet() *Pointer {
	ptr := extism_var_get(p.Ptr)

	return &Pointer{ptr}
}

func (p *Pointer) StoreU8(v uint8) {
	extism_store_u8(p.Ptr, v)
}

func (p *Pointer) LoadU8(v uint8) uint8 {
	return extism_load_u8(p.Ptr + extismPointer(v))
}

func (p *Pointer) StoreU64(v uint64) {
	extism_store_u64(p.Ptr, v)
}

func (p *Pointer) LoadU64(ptr *Pointer) uint64 {
	return extism_load_u64(p.Ptr + ptr.Ptr)
}

func (p *Pointer) HTTPRequest(ptr *Pointer) *Pointer {
	point := extism_http_request(p.Ptr, ptr.Ptr)

	return &Pointer{point}
}

func (p *Pointer) LogInfo() {
	extism_log_info(p.Ptr)
}

func (p *Pointer) LogDebug() {
	extism_log_debug(p.Ptr)
}

func (p *Pointer) LogWarn() {
	extism_log_warn(p.Ptr)
}

func (p *Pointer) LogError() {
	extism_log_error(p.Ptr)
}
