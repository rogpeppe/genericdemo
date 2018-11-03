package main

import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"

	"github.com/rogpeppe/genericdemo/generic"
)

type TypeId int

var _allTypeParams = map[generic.TypeTuple]TypeId{
	generic.Types(new(Int)):    0,
	generic.Types(new(Flag)):   1,
	generic.Types(new(Str)):    2,
	generic.Types(new(AdderI)): 3,
	generic.Types(new(Vec__0)): 4,
	generic.Types(new(Vec__4)): 5,
	generic.Types(new(Vec__5)): 6,
}

func main() {
	var gf generic.Funcs
	for _, r := range []func(*generic.Funcs){
		_register_addPair__0,
		_register_addPair__1,
		_register_addPair__2,
		_register_addPair__generic(generic.Types(new(AdderI))),
		_register_addPair__generic(generic.Types(new(Vec__0))),
		_register_addPair__generic(generic.Types(new(Vec__4))),
		_register_addPair__generic(generic.Types(new(Vec__5))),
		_register_sum__0,
		_register_sum__1,
		_register_sum__2,
		_register_sum__generic(generic.Types(new(AdderI))),
		_register_sum__generic(generic.Types(new(Vec__0))),
		_register_sum__generic(generic.Types(new(Vec__4))),
		_register_sum__generic(generic.Types(new(Vec__5))),
	} {
		r(&gf)
	}
	_addPair__0 := gf.Get("addPair", generic.Types(new(Int))).(func(Int, Int) Int)
	_addPair__1 := gf.Get("addPair", generic.Types(new(Flag))).(func(Flag, Flag) Flag)
	_addPair__2 := gf.Get("addPair", generic.Types(new(Str))).(func(Str, Str) Str)
	_addPair__3 := gf.Get("addPair", generic.Types(new(AdderI))).(func(AdderI, AdderI) AdderI)
	_addPair__4 := gf.Get("addPair", generic.Types(new(Vec__0))).(func(Vec__0, Vec__0) Vec__0)
	_addPair__5 := gf.Get("addPair", generic.Types(new(Vec__4))).(func(Vec__4, Vec__4) Vec__4)
	_addPair__6 := gf.Get("addPair", generic.Types(new(Vec__5))).(func(Vec__5, Vec__5) Vec__5)

	fmt.Println(_addPair__0(34, 56))
	fmt.Println(_addPair__2("hello ", "world"))
	fmt.Println(_addPair__1(Flag{2}, Flag{3}))
	fmt.Println(_addPair__3(StrAdderI{"hello "}, StrAdderI{"world"}))
	fmt.Println(_addPair__4(Vec__0{1, 3, 9}, Vec__0{2, 20}))
	fmt.Println(_addPair__5(Vec__4{{1, 2}, {3}, {9, 100}}, Vec__4{{9, 2, 200}, {20}}))
	fmt.Println(_addPair__6(Vec__5{{{1, 2}, {3}, {9, 100}}}, Vec__5{{{9, 2, 200}, {20}}, {{1000, 3000}}}))
}

// The following types are shared between all generic
// functions that use a particular pointer/value layout.
// The type name represents the pointer layout:
// "p" represents a pointer. "vN" represents N bytes.
//
// So _p would be a single pointer, and _v16ppv8
// might be used to represent a struct type like:
// struct {x, y int64; addPair *Bar; z struct {a, b int32}}
//
// When doing this properly, we'd need to take alignment
// into account too.

type _generic_v8 struct {
	_ int64
}

type _generic_pv8 struct {
	_ unsafe.Pointer
	_ [8]byte
}

type _generic_pv16 struct {
	_ unsafe.Pointer
	_ [16]byte
}

func _register_addPair__0(gf *generic.Funcs) {
	var inst _inst_addPair__v8
	gf.Add("addPair", generic.Types(new(Int)), generic.AsType(
		func(p0, p1 _generic_v8) _generic_v8 {
			inst.once.Do(func() {
				generic.UnsafeSet(&inst.sum, gf.Get("sum", generic.Types(new(Int))))
			})
			return addPair__v8(&inst, p0, p1)
		},
		(func(a, b Int) Int)(nil),
	))
}

// addPair(Int)
func _register_addPair__0_inline(gf *generic.Funcs) {
	gf.Add("addPair", generic.Types(new(Int)), addPair__0_inline)
}

// addPair(Flag)
func _register_addPair__1(gf *generic.Funcs) {
	var inst _inst_addPair__v8
	gf.Add("addPair", generic.Types(new(Flag)), generic.AsType(
		func(p0, p1 _generic_v8) _generic_v8 {
			inst.once.Do(func() {
				generic.UnsafeSet(&inst.sum, gf.Get("sum", generic.Types(new(Flag))))
			})
			return addPair__v8(&inst, p0, p1)
		},
		(func(a, b Flag) Flag)(nil),
	))
}

// addPair(Str)
func _register_addPair__2(gf *generic.Funcs) {
	var inst _inst_addPair__pv8
	gf.Add("addPair", generic.Types(new(Str)), generic.AsType(
		func(p0, p1 _generic_pv8) _generic_pv8 {
			inst.once.Do(func() {
				generic.UnsafeSet(&inst.sum, gf.Get("sum", generic.Types(new(Str))))
			})
			return addPair__pv8(&inst, p0, p1)
		},
		(func(a, b Str) Str)(nil),
	))
}

func _register_addPair__4(gf *generic.Funcs) {
	var inst _inst_addPair__pv16
	gf.Add("addPair", generic.Types(new(Str)), generic.AsType(
		func(p0, p1 _generic_pv16) _generic_pv16 {
			inst.once.Do(func() {
				generic.UnsafeSet(&inst.sum, gf.Get("sum", generic.Types(new(Vec__0))))
			})
			return addPair__pv16(&inst, p0, p1)
		},
		(func(a, b Vec__0) Vec__0)(nil),
	))
}

// addPair generic
func _register_addPair__generic(t generic.TypeTuple) func(gf *generic.Funcs) {
	return func(gf *generic.Funcs) {
		var inst _inst_addPair__generic
		t0 := t.At(0)
		gf.Add("addPair", t, reflect.MakeFunc(
			reflect.FuncOf([]reflect.Type{t0, t0}, []reflect.Type{t0}, false),
			func(args []reflect.Value) []reflect.Value {
				inst.once.Do(func() {
					inst.sum = reflect.ValueOf(gf.Get("sum", t))
					inst.slice = reflect.SliceOf(t0)
				})
				return addPair__generic(&inst, args)
			},
		).Interface())
	}
}

type _inst_addPair__v8 struct {
	once sync.Once
	sum  func([]_generic_v8) _generic_v8
}

func addPair__v8(_inst *_inst_addPair__v8, a, b _generic_v8) _generic_v8 {
	f := _inst.sum
	return f([]_generic_v8{a, b})
}

func addPair__0_inline(a, b Int) Int {
	return sum__0([]Int{a, b})
}

type _inst_addPair__pv8 struct {
	once sync.Once
	sum  func([]_generic_pv8) _generic_pv8
}

func addPair__pv8(_inst *_inst_addPair__pv8, a, b _generic_pv8) _generic_pv8 {
	f := _inst.sum
	return f([]_generic_pv8{a, b})
}

type _inst_addPair__pv16 struct {
	once sync.Once
	sum  func([]_generic_pv16) _generic_pv16
}

func addPair__pv16(_inst *_inst_addPair__pv16, a, b _generic_pv16) _generic_pv16 {
	f := _inst.sum
	return f([]_generic_pv16{a, b})
}

type _inst_addPair__generic struct {
	once  sync.Once
	sum   reflect.Value // func(T, T) T
	slice reflect.Type
}

func addPair__generic(_inst *_inst_addPair__generic, args []reflect.Value) []reflect.Value {
	a, b := args[0], args[1]
	f := _inst.sum
	_t0 := reflect.MakeSlice(_inst.slice, 2, 2)
	_t0.Index(0).Set(a)
	_t0.Index(1).Set(b)
	return f.Call([]reflect.Value{_t0})
}

func _register_sum__0(gf *generic.Funcs) {
	var inst _inst_sum__v8
	generic.UnsafeSet(&inst.add, Int.Add)
	gf.Add("sum", generic.Types(new(Int)), generic.AsType(
		func(p0 []_generic_v8) _generic_v8 {
			return sum__v8(&inst, p0)
		},
		(func([]Int) Int)(nil),
	))
}

func _register_sum__0_inline(gf *generic.Funcs) {
	gf.Add("sum", generic.Types(new(Int)), sum__0)
}

func _register_sum__1(gf *generic.Funcs) {
	var inst _inst_sum__v8
	generic.UnsafeSet(&inst.add, Flag.Add)
	gf.Add("sum", generic.Types(new(Flag)), generic.AsType(
		func(p0 []_generic_v8) _generic_v8 {
			return sum__v8(&inst, p0)
		},
		(func([]Flag) Flag)(nil),
	))
}

func _register_sum__2(gf *generic.Funcs) {
	var inst _inst_sum__pv8
	generic.UnsafeSet(&inst.add, Str.Add)
	gf.Add("sum", generic.Types(new(Str)), generic.AsType(
		func(p0 []_generic_pv8) _generic_pv8 {
			return sum__pv8(&inst, p0)
		},
		(func([]Str) Str)(nil),
	))
}

func _register_sum__generic(t generic.TypeTuple) func(gf *generic.Funcs) {
	return func(gf *generic.Funcs) {
		var inst _inst_sum__generic
		t0 := t.At(0)
		inst.t0 = t0
		m, ok := t0.MethodByName("Add")
		if !ok {
			panic(fmt.Errorf("%s has no Add method", t0))
		}
		if t0.Kind() == reflect.Interface {
			inst.add = reflect.MakeFunc(
				reflect.FuncOf([]reflect.Type{t0, t0}, []reflect.Type{t0}, false),
				func(args []reflect.Value) []reflect.Value {
					return args[0].Method(m.Index).Call(args[1:])
				},
			)
		} else {
			inst.add = m.Func
		}
		gf.Add("sum", t, reflect.MakeFunc(
			reflect.FuncOf([]reflect.Type{reflect.SliceOf(t0)}, []reflect.Type{t0}, false),
			func(args []reflect.Value) []reflect.Value {
				return sum__generic(&inst, args)
			},
		).Interface())
	}
}

type _inst_sum__v8 struct {
	add func(_generic_v8, _generic_v8) _generic_v8
}

func sum__v8(_inst *_inst_sum__v8, ts []_generic_v8) (x _generic_v8) {
	if len(ts) == 0 {
		return
	}
	x = ts[0]
	for _, t := range ts[1:] {
		x = _inst.add(x, t)
	}
	return x
}

type _inst_sum__pv8 struct {
	add func(_generic_pv8, _generic_pv8) _generic_pv8
}

func sum__pv8(_inst *_inst_sum__pv8, ts []_generic_pv8) (x _generic_pv8) {
	if len(ts) == 0 {
		return
	}
	x = ts[0]
	for _, t := range ts[1:] {
		x = _inst.add(x, t)
	}
	return x
}

func sum__0(ts []Int) (x Int) {
	if len(ts) == 0 {
		return
	}
	x = ts[0]
	for _, t := range ts[1:] {
		x = x.Add(t)
	}
	return x
}

type _inst_sum__generic struct {
	t0  reflect.Type
	add reflect.Value
}

func sum__generic(_inst *_inst_sum__generic, args []reflect.Value) []reflect.Value {
	ts := args[0]
	if ts.Len() == 0 {
		return []reflect.Value{reflect.Zero(_inst.t0)}
	}
	x := _copyVal(ts.Index(0))
	{
		_n := ts.Len()
		for _i := 1; _i < _n; _i++ {
			t := ts.Index(_i)
			x.Set(_inst.add.Call([]reflect.Value{x, t})[0])
		}
	}
	return []reflect.Value{x}
}

func _copyVal(v reflect.Value) reflect.Value {
	v1 := reflect.New(v.Type()).Elem()
	v1.Set(v)
	return v1
}

type Vec__0 []Int

var _inst_Vec__0 = _inst_Vec__v8{
	add: (func(_generic_v8, _generic_v8) _generic_v8)(unsafe.Pointer(Int.Add)),
}

func (v1 Vec__0) Add(v2 Vec__0) Vec__0 {
	return (func(*_inst_Vec__v8, Vec__0, Vec__0) Vec__0)(unsafe.Pointer(Vec__v8_Add))(&_inst_Vec__0, v1, v2)
}

type Vec__4 []Vec__0

var _inst_Vec__4 = _inst_Vec__pv16{
	add: (func(_generic_pv16, _generic_pv16) _generic_pv16)(unsafe.Pointer(Vec__0.Add)),
}

func (v1 Vec__4) Add(v2 Vec__4) Vec__4 {
	return (func(*_inst_Vec__pv16, Vec__4, Vec__4) Vec__4)(unsafe.Pointer(Vec__pv16_Add))(&_inst_Vec__4, v1, v2)
}

type Vec__5 []Vec__4

var _inst_Vec__5 = _inst_Vec__pv16{
	add: (func(_generic_pv16, _generic_pv16) _generic_pv16)(unsafe.Pointer(Vec__4.Add)),
}

func (v1 Vec__5) Add(v2 Vec__5) Vec__5 {
	return (func(*_inst_Vec__pv16, Vec__5, Vec__5) Vec__5)(unsafe.Pointer(Vec__pv16_Add))(&_inst_Vec__5, v1, v2)
}

type Vec__v8 []_generic_v8

type _inst_Vec__v8 struct {
	add  func(_generic_v8, _generic_v8) _generic_v8
}

func Vec__v8_Add(_inst *_inst_Vec__v8, v1 Vec__v8, v2 Vec__v8) Vec__v8 {
	if len(v2) > len(v1) {
		v1, v2 = v2, v1
	}
	r := make(Vec__v8, len(v1))
	for i, x := range v2 {
		r[i] = _inst.add(v1[i], x)
	}
	for i, x := range v1[len(v2):] {
		r[i+len(v2)] = x
	}
	return r
}

type Vec__pv16 []_generic_pv16

type _inst_Vec__pv16 struct {
	once sync.Once
	add  func(_generic_pv16, _generic_pv16) _generic_pv16
}

func Vec__pv16_Add(inst *_inst_Vec__pv16, v1 Vec__pv16, v2 Vec__pv16) Vec__pv16 {
	if len(v2) > len(v1) {
		v1, v2 = v2, v1
	}
	r := make(Vec__pv16, len(v1))
	for i, x := range v2 {
		r[i] = inst.add(v1[i], x)
	}
	for i, x := range v1[len(v2):] {
		r[i+len(v2)] = x
	}
	return r
}
