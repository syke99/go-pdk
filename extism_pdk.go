package pdk

import (
	"github.com/extism/go-pdk/internal"
	"github.com/extism/go-pdk/memory"
)

type Pointer struct {
	ptr *internal.Pointer
}

type LogLevel int

const (
	LogInfo LogLevel = iota
	LogDebug
	LogWarn
	LogError
)

func InputString() string {
	return string(Input())
}

func Input() []byte {
	return internal.LoadInput()
}

func Output(data []byte) {
	clength := uint64(len(data))
	offset := internal.Alloc(clength)

	internal.Store(offset, data)

	offset.OutputSet(clength)
}

func OutputString(s string) {
	Output([]byte(s))
}

func Allocate(length int) *memory.Memory {
	clength := uint64(length)
	offset := internal.Alloc(clength)

	return &memory.Memory{
		Off: offset,
		Len: clength,
	}
}

func AllocateBytes(data []byte) *memory.Memory {
	offset, clength := internal.AllocBytes(data)

	return &memory.Memory{
		Off: offset,
		Len: clength,
	}

}

func AllocateString(data string) *memory.Memory {
	return AllocateBytes([]byte(data))
}

func GetConfig(key string) (string, bool) {
	mem := AllocateBytes([]byte(key))
	defer mem.Free()

	offset := mem.Off
	if offset == nil {
		return "", false
	}

	config := offset.ConfigGet()
	if config == nil {
		return "", false
	}

	clength := config.Length()
	if clength == 0 {
		return "", false
	}

	value := make([]byte, clength)
	internal.Load(config, value)

	return string(value), true
}

func LogPDKMemory(level LogLevel, m *memory.Memory) {
	switch level {
	case LogInfo:
		m.Off.LogInfo()
	case LogDebug:
		m.Off.LogDebug()
	case LogWarn:
		m.Off.LogWarn()
	case LogError:
		m.Off.LogError()
	}
}

func Log(level LogLevel, s string) {
	mem := AllocateString(s)
	defer mem.Free()

	LogPDKMemory(level, mem)
}

func GetVar(key string) []byte {
	mem := AllocateBytes([]byte(key))

	offset := mem.Off
	if offset == nil {
		return nil
	}

	v := offset.VarGet()
	if v == nil {
		return nil
	}

	clength := mem.Off.Length()
	if clength == 0 {
		return nil
	}

	value := make([]byte, clength)
	internal.Load(v, value)

	return value
}

func SetVar(key string, value []byte) {
	keyMem := AllocateBytes([]byte(key))
	defer keyMem.Free()

	valMem := AllocateBytes(value)
	defer valMem.Free()

	keyMem.Off.VarSet(valMem.Off)
}

func RemoveVar(key string) {
	mem := AllocateBytes([]byte(key))
	mem.Off.VarSet(nil)
}
