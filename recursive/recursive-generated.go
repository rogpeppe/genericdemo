package main
import (
	"unsafe"
	"fmt"
)

type _generic_v8a8 struct {
	t0 uint64
}

func main() {
	fmt.Println(Foo__0(3, 5))
}

type _inst_Foo__v8a8 struct {
	_func_Bar func(_generic_v8a8, int) _generic_v8a8
}

var _inst_Foo__0 = &_inst_Foo__v8a8{
	_func_Bar: (func(_generic_v8a8, int) _generic_v8a8)(unsafe.Pointer(Bar__0)),
}

func Foo__0(t int64, i int) int64 {
	return (func(*_inst_Foo__v8a8, int64, int) int64)(unsafe.Pointer(Foo__v8a8))(_inst_Foo__0, t, i)
}

func Foo__v8a8(_inst *_inst_Foo__v8a8, t _generic_v8a8, i int) _generic_v8a8 {
	if i <= 0 {
		return t
	}
	return _inst._func_Bar(t, i-1)
}

type _inst_Bar__v8a8 struct {
	_func_Foo func(_generic_v8a8, int) _generic_v8a8
}

var _inst_Bar__0 = &_inst_Bar__v8a8{
	_func_Foo: (func(_generic_v8a8, int) _generic_v8a8)(unsafe.Pointer(Foo__0)),
}

func Bar__0(t int64, i int) int64 {
	return (func(*_inst_Bar__v8a8, int64, int) int64)(unsafe.Pointer(Bar__v8a8))(_inst_Bar__0, t, i)
}

func Bar__v8a8(_inst *_inst_Bar__v8a8, t _generic_v8a8, i int) _generic_v8a8 {
	if i <= 0 {
		return t
	}
	return _inst._func_Foo(t, i-1)
}
