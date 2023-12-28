package internal

import "encoding/binary"

func Load(p *Pointer, buf []byte) {
	length := len(buf)

	chunkCount := length >> 3

	for chunkIdx := 0; chunkIdx < chunkCount; chunkIdx++ {
		i := chunkIdx << 3

		binary.LittleEndian.PutUint64(buf[i:i+8], p.LoadU64(Ptr(uint64(i))))
	}

	remainder := length & 7
	remainderOffset := chunkCount << 3
	for index := remainderOffset; index < (remainder + remainderOffset); index++ {
		buf[index] = byte(p.LoadU64(Ptr(uint64(index))))
	}
}

func LoadInput() []byte {
	length := int(InputLength())
	buf := make([]byte, length)

	chunkCount := length >> 3

	for chunkIdx := 0; chunkIdx < chunkCount; chunkIdx++ {
		i := chunkIdx << 3

		binary.LittleEndian.PutUint64(buf[i:i+8], Ptr(uint64(i)).InputLoadU64())
	}

	remainder := length & 7
	remainderOffset := chunkCount << 3
	for index := remainderOffset; index < (remainder + remainderOffset); index++ {
		buf[index] = Ptr(uint64(index)).InputLoadU8()
	}

	return buf
}

func Store(p *Pointer, buf []byte) {
	length := len(buf)

	chunkCount := length >> 3

	for chunkIdx := 0; chunkIdx < chunkCount; chunkIdx++ {
		i := chunkIdx << 3
		x := binary.LittleEndian.Uint64(buf[i : i+8])

		p.AddPtr(Ptr(uint64(i))).StoreU64(x)
	}

	remainder := length & 7
	remainderOffset := chunkCount << 3
	for index := remainderOffset; index < (remainder + remainderOffset); index++ {
		p.AddPtr(Ptr(uint64(index))).StoreU64(uint64(buf[index]))
	}
}

func AllocBytes(data []byte) (*Pointer, uint64) {
	clength := uint64(len(data))
	offset := Alloc(clength)

	Store(offset, data)

	return offset, clength
}
