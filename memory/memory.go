package memory

import (
	"github.com/extism/go-pdk/internal"
)

type Memory struct {
	Off *internal.Pointer
	Len uint64
}

func (m *Memory) Load(buffer []byte) {
	internal.Load(m.Off, buffer)
}

func (m *Memory) Store(data []byte) {
	internal.Store(m.Off, data)
}

func (m *Memory) Free() {
	m.Off.Free()
}

func (m *Memory) Length() uint64 {
	return m.Len
}

func (m *Memory) Offset() uint64 {
	return uint64(m.Off.Ptr)
}

func FindMemory(offset uint64) *Memory {
	length := internal.Ptr(offset).Length()

	return &Memory{
		Off: internal.Ptr(offset),
		Len: length,
	}
}

func (m *Memory) Output() {
	m.Off.OutputSet(m.Len)
}

type LogLevel int
