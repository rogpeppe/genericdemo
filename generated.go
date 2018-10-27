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
	//	generic.Types(new(Vec__0)): 4,
	//	generic.Types(new(Vec__4)): 5,
}

func main() {
	var gf generic.Funcs
	for _, r := range []func(*generic.Funcs){
		_register_addPair__0,
		_register_addPair__1,
		_register_addPair__2,
		_register_addPair__generic(generic.Types(new(AdderI))),
		_register_sum__0,
		_register_sum__1,
		_register_sum__2,
		_register_sum__generic(generic.Types(new(AdderI))),
	} {
		r(&gf)
	}
	_addPair__0 := gf.Get("addPair", generic.Types(new(Int))).(func(Int, Int) Int)
	_addPair__1 := gf.Get("addPair", generic.Types(new(Flag))).(func(Flag, Flag) Flag)
	_addPair__2 := gf.Get("addPair", generic.Types(new(Str))).(func(Str, Str) Str)
	_addPair__3 := gf.Get("addPair", generic.Types(new(AdderI))).(func(AdderI, AdderI) AdderI)

	fmt.Println(_addPair__0(34, 56))
	fmt.Println(_addPair__2("hello ", "world"))
	fmt.Println(_addPair__1(Flag{2}, Flag{3}))
	fmt.Println(_addPair__3(StrAdderI{"hello "}, StrAdderI{"world"}))
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

type _generic__v8 struct {
	_ int64
}

type _generic__pv8 struct {
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
		func(p0, p1 _generic__v8) _generic__v8 {
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
		func(p0, p1 _generic__v8) _generic__v8 {
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
		func(p0, p1 _generic__pv8) _generic__pv8 {
			inst.once.Do(func() {
				generic.UnsafeSet(&inst.sum, gf.Get("sum", generic.Types(new(Str))))
			})
			return addPair__pv8(&inst, p0, p1)
		},
		(func(a, b Str) Str)(nil),
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
	sum  func([]_generic__v8) _generic__v8
}

func addPair__v8(_inst *_inst_addPair__v8, a, b _generic__v8) _generic__v8 {
	f := _inst.sum
	return f([]_generic__v8{a, b})
}

func addPair__0_inline(a, b Int) Int {
	return sum__0([]Int{a, b})
}

type _inst_addPair__pv8 struct {
	once sync.Once
	sum  func([]_generic__pv8) _generic__pv8
}

func addPair__pv8(_inst *_inst_addPair__pv8, a, b _generic__pv8) _generic__pv8 {
	f := _inst.sum
	return f([]_generic__pv8{a, b})
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
		func(p0 []_generic__v8) _generic__v8 {
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
		func(p0 []_generic__v8) _generic__v8 {
			return sum__v8(&inst, p0)
		},
		(func([]Flag) Flag)(nil),
	))
}

func _register_sum__2(gf *generic.Funcs) {
	var inst _inst_sum__pv8
	generic.UnsafeSet(&inst.add, Str.Add)
	gf.Add("sum", generic.Types(new(Str)), generic.AsType(
		func(p0 []_generic__pv8) _generic__pv8 {
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
	add func(_generic__v8, _generic__v8) _generic__v8
}

func sum__v8(_inst *_inst_sum__v8, ts []_generic__v8) (x _generic__v8) {
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
	add func(_generic__pv8, _generic__pv8) _generic__pv8
}

func sum__pv8(_inst *_inst_sum__pv8, ts []_generic__pv8) (x _generic__pv8) {
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

//
//func register_Vec_Vec__0(
//	var inst _VecData_pv16
//	generic.UnsafeSet(&inst.add, Int.Add)
//	t1 := gf.GetType("Vec", generic.Types(new(Int)))
//	STOP IT!
//	we hadn't thought of types of types...
//	use Instantiate (from generictypes.go)
//
//	gf.AddType("Vec", generic.Types(
//	gf.Add("sum", generic.Types(new(Str)), generic.AsType(
//		func(p0 []_generic__pv8) _generic__pv8 {
//			return sum__pv8(&inst, p0)
//		},
//		(func([]Str) Str)(nil),
//	))
//}

//var _VecData_
//
//type _VecData_Vec__0 []Vec__0
//
//func (v1 Vec__0) Add(v2 Vec__0) Vec__0 {
//	return Vec__0(((Vec_pv16)(v1)).Add(&_VecData_pv16, (Vec_pv16)(v2)))
//}
//
//type Vec_pv16 []_generic_pv16
//
//type _VecData_pv16 struct {
//	add func(_generic_pv16, _generic_pv16) _generic_pv16
//}
//
//func (v1 Vec_pv16) Add(inst *_VecData_pv16, v2 Vec_pv16) Vec_pv16 {
//	if len(v2) > len(v1) {
//		v1, v2 = v2, v1
//	}
//	r := make(Vec(T), len(v1))
//	for i, x := range v2 {
//		r[i] = inst.add(v1[i], v2[i])
//	}
//	for i, x := range v1[len(v2):] {
//		r[i+len(v2)] = x
//	}
//	return r
//}
