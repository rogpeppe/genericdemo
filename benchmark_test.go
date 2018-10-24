package main

import "testing"

func BenchmarkInline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo_Int_inline(34, 56)
	}
}

func BenchmarkDirect(b *testing.B) {
	var gf GenericFuncs
	for _, r := range []func(*GenericFuncs){
		register_foo_Int_inline,
		register_foo_Flag,
		register_foo_Str,
		register_sum_Int_inline,
		register_sum_Flag,
		register_sum_Str,
	} {
		r(&gf)
	}
	gf.Start()
	_foo_Int := gf.Get("foo", Types(new(Int))).(func(Int, Int) Int)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_foo_Int(34, 56)
	}
}

func BenchmarkDirectInner(b *testing.B) {
	var gf GenericFuncs
	for _, r := range []func(*GenericFuncs){
		register_foo_Int,
		register_foo_Flag,
		register_foo_Str,
		register_sum_Int_inline,
		register_sum_Flag,
		register_sum_Str,
	} {
		r(&gf)
	}
	gf.Start()
	_foo_Int := gf.Get("foo", Types(new(Int))).(func(Int, Int) Int)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_foo_Int(34, 56)
	}
}

func BenchmarkReuse(b *testing.B) {
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

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_foo_Int(34, 56)
	}
}

func BenchmarkGenericImpl(b *testing.B) {
	var gf GenericFuncs
	for _, r := range []func(*GenericFuncs){
		register_foo_generic(Types(new(Int))),
		register_foo_generic(Types(new(Flag))),
		register_foo_generic(Types(new(Str))),
		register_sum_generic(Types(new(Int))),
		register_sum_generic(Types(new(Flag))),
		register_sum_generic(Types(new(Str))),
	} {
		r(&gf)
	}
	gf.Start()
	_foo_Int := gf.Get("foo", Types(new(Int))).(func(Int, Int) Int)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_foo_Int(34, 56)
	}
}
