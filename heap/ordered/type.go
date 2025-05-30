package ordered

type Type interface {
	LessThan(other any) bool
}

type Int int

func (i Int) LessThan(other any) bool {
	return i < other.(Int)
}

type Int8 int8

func (i Int8) LessThan(other any) bool {
	return i < other.(Int8)
}

type Int16 int16

func (i Int16) LessThan(other any) bool {
	return i < other.(Int16)
}

type Int32 int32

func (i Int32) LessThan(other any) bool {
	return i < other.(Int32)
}

type Int64 int64

func (i Int64) LessThan(other any) bool {
	return i < other.(Int64)
}

type Uint uint

func (i Uint) LessThan(other any) bool {
	return i < other.(Uint)
}

type Uint8 uint8

func (i Uint8) LessThan(other any) bool {
	return i < other.(Uint8)
}

type Uint16 uint16

func (i Uint16) LessThan(other any) bool {
	return i < other.(Uint16)
}

type Uint32 uint

func (i Uint32) LessThan(other any) bool {
	return i < other.(Uint32)
}

type Uint64 uint64

func (i Uint64) LessThan(other any) bool {
	return i < other.(Uint64)
}

type Uintptr uintptr

func (i Uintptr) LessThan(other any) bool {
	return i < other.(Uintptr)
}

type Float32 float32

func (i Float32) LessThan(other any) bool {
	return i < other.(Float32)
}

type Float64 float64

func (i Float64) LessThan(other any) bool {
	return i < other.(Float64)
}

type String string

func (i String) LessThan(other any) bool {
	return i < other.(String)
}
