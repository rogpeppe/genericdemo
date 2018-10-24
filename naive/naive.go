/*
Naive generic code using reflect.

This is a rough hand-coded translation of the following code, without
any attempt at optimization, by way of demonstrating that such a thing
is possible.

Symbols not in the original code are preceded with underscores.

	package main

	func main() {
		fmt.Println(addPair(Int)(34, 56))
		fmt.Println(addPair(Str)("a", "b"))
	}

	type Adder(type T) interface {
		Add() T
	}

	contract AdderC(t T) {
		Adder(T)(t)
	}

	func addPair(type T AdderC)(a, b T) T {
		f := sum(T)
		return f([]T{a, b})
	}

	func sum(type T AdderC)(ts []T) (x T) {
		if len(ts) == 0 {
			return
		}
		x = ts[0]
		for _, t := range ts[1:] {
			x = x.Add(t)
		}
		return x
	}
*/
package main

import (
	"fmt"
	"reflect"
)

var (
	_IntType = reflect.TypeOf(new(Int)).Elem()
	_StrType = reflect.TypeOf(new(Str)).Elem()
)

func main() {
	fmt.Println(addPair_generic(_Types{_IntType}).Interface().(func(Int, Int) Int)(34, 56))
	fmt.Println(addPair_generic(_Types{_StrType}).Interface().(func(Str, Str) Str)("hello ", "world"))
}

type Int int

func (i Int) Add(j Int) Int {
	return i + j
}

type Str string

func (s Str) Add(t Str) Str {
	return s + t
}

func addPair_generic(t _Types) reflect.Value {
	_t0 := t[0]
	_t1 := reflect.SliceOf(_t0)
	return reflect.MakeFunc(
		reflect.FuncOf([]reflect.Type{_t0, _t0}, []reflect.Type{_t0}, false),
		func(args []reflect.Value) []reflect.Value {
			a, b := args[0], args[1]
			f := sum_generic(_Types{_t0})
			_t2 := reflect.MakeSlice(_t1, 2, 2)
			_t2.Index(0).Set(a)
			_t2.Index(1).Set(b)
			return f.Call([]reflect.Value{_t2})
		},
	)
}

func sum_generic(t _Types) reflect.Value {
	_t0 := t[0]
	// Assume type checker has already checked for method presence.
	_m, _ := _t0.MethodByName("Add")
	var _Add reflect.Value
	if _t0.Kind() == reflect.Interface {
		_Add = reflect.MakeFunc(
			reflect.FuncOf([]reflect.Type{_t0, _t0}, []reflect.Type{_t0}, false),
			func(args []reflect.Value) []reflect.Value {
				return args[0].Method(_m.Index).Call(args[1:])
			},
		)
	} else {
		_Add = _m.Func
	}
	return reflect.MakeFunc(
		reflect.FuncOf([]reflect.Type{reflect.SliceOf(_t0)}, []reflect.Type{_t0}, false),
		func(_args []reflect.Value) []reflect.Value {
			ts := _args[0]
			if ts.Len() == 0 {
				return []reflect.Value{reflect.Zero(_t0)}
			}
			x := _copyVal(ts.Index(0))
			{
				_n := ts.Len()
				for _i := 1; _i < _n; _i++ {
					t := ts.Index(_i)
					x.Set(_Add.Call([]reflect.Value{x, t})[0])
				}
			}
			return []reflect.Value{x}
		},
	)
}

// _Types represents a set of type parameters.
type _Types []reflect.Type

func _copyVal(v reflect.Value) reflect.Value {
	v1 := reflect.New(v.Type()).Elem()
	v1.Set(v)
	return v1
}
