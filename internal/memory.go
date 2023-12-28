package internal

import "encoding/binary"

func Load(p *Pointer, buf []byte) {
	length := len(buf)

	for i := 0; i < length; i++ {
		if length-i >= 8 {
			x := p.LoadU64(Ptr(uint64(i)))
			binary.LittleEndian.PutUint64(buf[i:i+8], x)
			i += 7
			continue
		}
		buf[i] = p.LoadU8(uint8(i))
	}
}

func LoadInput() []byte {
	length := int(InputLength())
	buf := make([]byte, length)

	for i := 0; i < length; i++ {
		if length-i >= 8 {
			x := Ptr(uint64(i)).InputLoadU64()
			binary.LittleEndian.PutUint64(buf[i:i+8], x)
			i += 7
			continue
		}
		buf[i] = Ptr(uint64(i)).InputLoadU8()
	}

	return buf
}

func Store(p *Pointer, buf []byte) {
	length := len(buf)

	for i := 0; i < length; i++ {
		if length-i >= 8 {
			x := binary.LittleEndian.Uint64(buf[i : i+8])

			p.AddPtr(Ptr(uint64(i))).StoreU64(x)
			i += 7
			continue
		}

		p.AddPtr(Ptr(uint64(i))).StoreU8(buf[i])
	}
}

func AllocBytes(data []byte) (*Pointer, uint64) {
	clength := uint64(len(data))
	offset := Alloc(clength)

	Store(offset, data)

	return offset, clength
}
