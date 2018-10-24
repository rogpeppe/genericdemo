package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var gf GenericFuncs
	for _, r := range []func(*GenericFuncs){
		register_foo_Int,
		register_foo_Flag,
		register_foo_Str,
		register_sum_Int,
		register_sum_Flag,
		register_sum_Str,
	} {
		r(&gf)
	}
	gf.Start()
	_foo_Int := gf.Get("foo", Types(new(Int))).(func(Int, Int) Int)
	_foo_Flag := gf.Get("foo", Types(new(Flag))).(func(Flag, Flag) Flag)
	_foo_Str := gf.Get("foo", Types(new(Str))).(func(Str, Str) Str)

	fmt.Println(_foo_Int(34, 56))
	fmt.Println(_foo_Str("hello ", "world"))
	fmt.Println(_foo_Flag(Flag{2}, Flag{3}))
}

// The following types are shared between all generic
// functions that use a particular pointer/value layout.
// The type name represents the pointer layout:
// "p" represents a pointer. "vN" represents N bytes.
//
// So _p would be a single pointer, and _v16ppv8
// might be used to represent a struct type like:
// struct {x, y int64; foo *Bar; z struct {a, b int32}}
//
// When doing this properly, we'd need to take alignment
// into account too.

type _v8 struct {
	_ int64
}

type _pv8 struct {
	_ unsafe.Pointer
	_ [8]byte
}

// foo(Int)
func register_foo_Int(gf *GenericFuncs) {
	var inst _fooData_v8
	gf.Add("foo", Types(new(Int)), asType(
		func(p0, p1 _v8) _v8 {
			return foo_v8(&inst, p0, p1)
		},
		(func(a, b Int) Int)(nil),
	))
	gf.AddCompleter(func() {
		unsafeSet(&inst.sum, gf.Get("sum", Types(new(Int))))
	})
}

func register_foo_Int_inline(gf *GenericFuncs) {
	gf.Add("foo", Types(new(Int)), foo_Int_inline)
}

func register_foo_Int_generic(gf *GenericFuncs) {
	var inst _fooData_generic
	gf.Add("foo", Types(new(Int)), reflect.MakeFunc(
		reflect.TypeOf((func(Int, Int) Int)(nil)),
		func(args []reflect.Value) []reflect.Value {
			return foo_generic(&inst, args)
		},
	).Interface())
	gf.AddCompleter(func() {
		inst.sum = reflect.ValueOf(gf.Get("sum", Types(new(Int))))
		inst.slice = reflect.SliceOf(reflect.TypeOf(new(Int)).Elem())
	})
}

// foo(Flag)
func register_foo_Flag(gf *GenericFuncs) {
	var inst _fooData_v8
	gf.Add("foo", Types(new(Flag)), asType(
		func(p0, p1 _v8) _v8 {
			return foo_v8(&inst, p0, p1)
		},
		(func(a, b Flag) Flag)(nil),
	))
	gf.AddCompleter(func() {
		unsafeSet(&inst.sum, gf.Get("sum", Types(new(Flag))))
	})
}

// foo(Str)
func register_foo_Str(gf *GenericFuncs) {
	var inst _fooData_pv8
	gf.Add("foo", Types(new(Str)), asType(
		func(p0, p1 _pv8) _pv8 {
			return foo_pv8(&inst, p0, p1)
		},
		(func(a, b Str) Str)(nil),
	))
	gf.AddCompleter(func() {
		unsafeSet(&inst.sum, gf.Get("sum", Types(new(Str))))
	})
}

type _fooData_v8 struct {
	sum func([]_v8) _v8
}

func foo_v8(_inst *_fooData_v8, a, b _v8) _v8 {
	f := _inst.sum
	return f([]_v8{a, b})
}

func foo_Int_inline(a, b Int) Int {
	return sum_Int([]Int{a, b})
}

type _fooData_pv8 struct {
	sum func([]_pv8) _pv8
}

func foo_pv8(_inst *_fooData_pv8, a, b _pv8) _pv8 {
	f := _inst.sum
	return f([]_pv8{a, b})
}

type _fooData_generic struct {
	sum   reflect.Value // func(T, T) T
	slice reflect.Type
}

func foo_generic(_inst *_fooData_generic, args []reflect.Value) []reflect.Value {
	a, b := args[0], args[1]
	f := _inst.sum
	_t0 := reflect.MakeSlice(_inst.slice, 2, 2)
	_t0.Index(0).Set(a)
	_t0.Index(1).Set(b)
	return f.Call([]reflect.Value{_t0})
}

// sum(Int)
func register_sum_Int(gf *GenericFuncs) {
	var inst _sumData_v8
	gf.Add("sum", Types(new(Int)), asType(
		func(p0 []_v8) _v8 {
			return sum_v8(&inst, p0)
		},
		(func([]Int) Int)(nil),
	))
	gf.AddCompleter(func() {
		unsafeSet(&inst.add, Int.Add)
	})
}

func register_sum_Int_inline(gf *GenericFuncs) {
	gf.Add("sum", Types(new(Int)), sum_Int)
}

// sum(Flag)
func register_sum_Flag(gf *GenericFuncs) {
	var inst _sumData_v8
	gf.Add("sum", Types(new(Flag)), asType(
		func(p0 []_v8) _v8 {
			return sum_v8(&inst, p0)
		},
		(func([]Flag) Flag)(nil),
	))
	gf.AddCompleter(func() {
		unsafeSet(&inst.add, Flag.Add)
	})
}

// sum(Str)
func register_sum_Str(gf *GenericFuncs) {
	var inst _sumData_pv8
	gf.Add("sum", Types(new(Str)), asType(
		func(p0 []_pv8) _pv8 {
			return sum_pv8(&inst, p0)
		},
		(func([]Str) Str)(nil),
	))
	gf.AddCompleter(func() {
		unsafeSet(&inst.add, Str.Add)
	})
}

type _sumData_v8 struct {
	add func(_v8, _v8) _v8
}

type _sumData_pv8 struct {
	add func(_pv8, _pv8) _pv8
}

func sum_v8(_inst *_sumData_v8, ts []_v8) _v8 {
	var x _v8
	for _, t := range ts {
		x = _inst.add(x, t)
	}
	return x
}

func sum_pv8(_inst *_sumData_pv8, ts []_pv8) _pv8 {
	var x _pv8
	for _, t := range ts {
		x = _inst.add(x, t)
	}
	return x
}

func sum_Int(ts []Int) Int {
	var x Int
	for _, t := range ts {
		x = x.Add(t)
	}
	return x
}
