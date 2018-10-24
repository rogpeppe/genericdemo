package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var gf GenericFuncs
	for _, r := range []func(*GenericFuncs){
		register_addPair_Int,
		register_addPair_Flag,
		register_addPair_Str,
		register_addPair_generic(Types(new(AdderI))),
		register_sum_Int,
		register_sum_Flag,
		register_sum_Str,
		register_sum_generic(Types(new(AdderI))),
	} {
		r(&gf)
	}
	gf.Start()
	_addPair_Int := gf.Get("addPair", Types(new(Int))).(func(Int, Int) Int)
	_addPair_Flag := gf.Get("addPair", Types(new(Flag))).(func(Flag, Flag) Flag)
	_addPair_Str := gf.Get("addPair", Types(new(Str))).(func(Str, Str) Str)
	_addPair_Adder := gf.Get("addPair", Types(new(AdderI))).(func(AdderI, AdderI) AdderI)

	fmt.Println(_addPair_Int(34, 56))
	fmt.Println(_addPair_Str("hello ", "world"))
	fmt.Println(_addPair_Flag(Flag{2}, Flag{3}))
	fmt.Println(_addPair_Adder(StrAdderI{"hello "}, StrAdderI{"world"}))
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

func register_addPair_Int(gf *GenericFuncs) {
	var inst _addPairData_v8
	gf.Add("addPair", Types(new(Int)), asType(
		func(p0, p1 _generic_v8) _generic_v8 {
			return addPair_v8(&inst, p0, p1)
		},
		(func(a, b Int) Int)(nil),
	))
	gf.AddInit(func() {
		unsafeSet(&inst.sum, gf.Get("sum", Types(new(Int))))
	})
}

func register_addPair_Int_inline(gf *GenericFuncs) {
	gf.Add("addPair", Types(new(Int)), addPair_Int_inline)
}

func register_addPair_Flag(gf *GenericFuncs) {
	var inst _addPairData_v8
	gf.Add("addPair", Types(new(Flag)), asType(
		func(p0, p1 _generic_v8) _generic_v8 {
			return addPair_v8(&inst, p0, p1)
		},
		(func(a, b Flag) Flag)(nil),
	))
	gf.AddInit(func() {
		unsafeSet(&inst.sum, gf.Get("sum", Types(new(Flag))))
	})
}

func register_addPair_Str(gf *GenericFuncs) {
	var inst _addPairData_pv8
	gf.Add("addPair", Types(new(Str)), asType(
		func(p0, p1 _generic_pv8) _generic_pv8 {
			return addPair_pv8(&inst, p0, p1)
		},
		(func(a, b Str) Str)(nil),
	))
	gf.AddInit(func() {
		unsafeSet(&inst.sum, gf.Get("sum", Types(new(Str))))
	})
}

func register_addPair_generic(t TypeTuple) func(gf *GenericFuncs) {
	return func(gf *GenericFuncs) {
		var inst _addPairData_generic
		t0 := t.At(0)
		gf.Add("addPair", t, reflect.MakeFunc(
			reflect.FuncOf([]reflect.Type{t0, t0}, []reflect.Type{t0}, false),
			func(args []reflect.Value) []reflect.Value {
				return addPair_generic(&inst, args)
			},
		).Interface())
		gf.AddInit(func() {
			inst.sum = reflect.ValueOf(gf.Get("sum", t))
			inst.slice = reflect.SliceOf(t0)
		})
	}
}

type _addPairData_v8 struct {
	sum func([]_generic_v8) _generic_v8
}

func addPair_v8(_inst *_addPairData_v8, a, b _generic_v8) _generic_v8 {
	f := _inst.sum
	return f([]_generic_v8{a, b})
}

func addPair_Int_inline(a, b Int) Int {
	return sum_Int([]Int{a, b})
}

type _addPairData_pv8 struct {
	sum func([]_generic_pv8) _generic_pv8
}

func addPair_pv8(_inst *_addPairData_pv8, a, b _generic_pv8) _generic_pv8 {
	f := _inst.sum
	return f([]_generic_pv8{a, b})
}

type _addPairData_generic struct {
	sum   reflect.Value // func(T, T) T
	slice reflect.Type
}

func addPair_generic(_inst *_addPairData_generic, args []reflect.Value) []reflect.Value {
	a, b := args[0], args[1]
	f := _inst.sum
	_t0 := reflect.MakeSlice(_inst.slice, 2, 2)
	_t0.Index(0).Set(a)
	_t0.Index(1).Set(b)
	return f.Call([]reflect.Value{_t0})
}

func register_sum_Int(gf *GenericFuncs) {
	var inst _sumData_v8
	gf.Add("sum", Types(new(Int)), asType(
		func(p0 []_generic_v8) _generic_v8 {
			return sum_v8(&inst, p0)
		},
		(func([]Int) Int)(nil),
	))
	gf.AddInit(func() {
		unsafeSet(&inst.add, Int.Add)
	})
}

func register_sum_Int_inline(gf *GenericFuncs) {
	gf.Add("sum", Types(new(Int)), sum_Int)
}

func register_sum_Flag(gf *GenericFuncs) {
	var inst _sumData_v8
	gf.Add("sum", Types(new(Flag)), asType(
		func(p0 []_generic_v8) _generic_v8 {
			return sum_v8(&inst, p0)
		},
		(func([]Flag) Flag)(nil),
	))
	gf.AddInit(func() {
		unsafeSet(&inst.add, Flag.Add)
	})
}

func register_sum_Str(gf *GenericFuncs) {
	var inst _sumData_pv8
	gf.Add("sum", Types(new(Str)), asType(
		func(p0 []_generic_pv8) _generic_pv8 {
			return sum_pv8(&inst, p0)
		},
		(func([]Str) Str)(nil),
	))
	gf.AddInit(func() {
		unsafeSet(&inst.add, Str.Add)
	})
}

func register_sum_generic(t TypeTuple) func(gf *GenericFuncs) {
	return func(gf *GenericFuncs) {
		var inst _sumData_generic
		t0 := t.At(0)
		gf.Add("sum", t, reflect.MakeFunc(
			reflect.FuncOf([]reflect.Type{reflect.SliceOf(t0)}, []reflect.Type{t0}, false),
			func(args []reflect.Value) []reflect.Value {
				return sum_generic(&inst, args)
			},
		).Interface())
		gf.AddInit(func() {
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
		})
	}
}

type _sumData_v8 struct {
	add func(_generic_v8, _generic_v8) _generic_v8
}

func sum_v8(_inst *_sumData_v8, ts []_generic_v8) (x _generic_v8) {
	if len(ts) == 0 {
		return
	}
	x = ts[0]
	for _, t := range ts[1:] {
		x = _inst.add(x, t)
	}
	return x
}

type _sumData_pv8 struct {
	add func(_generic_pv8, _generic_pv8) _generic_pv8
}

func sum_pv8(_inst *_sumData_pv8, ts []_generic_pv8) (x _generic_pv8) {
	if len(ts) == 0 {
		return
	}
	x = ts[0]
	for _, t := range ts[1:] {
		x = _inst.add(x, t)
	}
	return x
}

func sum_Int(ts []Int) (x Int) {
	if len(ts) == 0 {
		return
	}
	x = ts[0]
	for _, t := range ts[1:] {
		x = x.Add(t)
	}
	return x
}

type _sumData_generic struct {
	t0  reflect.Type
	add reflect.Value
}

func sum_generic(_inst *_sumData_generic, args []reflect.Value) []reflect.Value {
	ts := args[0]
	if ts.Len() == 0 {
		return []reflect.Value{reflect.Zero(_inst.t0)}
	}
	x := copyVal(ts.Index(0))
	{
		_n := ts.Len()
		for _i := 1; _i < _n; _i++ {
			t := ts.Index(_i)
			x.Set(_inst.add.Call([]reflect.Value{x, t})[0])
		}
	}
	return []reflect.Value{x}
}
