package parse

const (
	Invalid = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Float32
	Float64
	Complex64
	Complex128
	Interface
	Bytes
	Map
	Ptr
	String
	Slice
	Struct
)

var Types = map[string]uint8{
	"bool":        Bool,
	"int":         Int,
	"int8":        Int8,
	"int16":       Int16,
	"int32":       Int32,
	"int64":       Int64,
	"uint":        Uint,
	"uint8":       Uint8,
	"uint16":      Uint16,
	"uint32":      Uint32,
	"uint64":      Uint64,
	"float32":     Float32,
	"float64":     Float64,
	"complex64":   Complex64,
	"complex128":  Complex128,
	"interface{}": Interface,
	"[]byte":      Bytes,
	"string":      String,
}
