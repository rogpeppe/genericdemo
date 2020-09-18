package main

import (
	"testing"

	"github.com/rogpeppe/genericdemo/generic"
)

func BenchmarkInline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addPair__0_inline(34, 56)
	}
}

func BenchmarkDirect(b *testing.B) {
	var gf generic.Funcs
	for _, r := range []func(*generic.Funcs){
		_register_addPair__0_inline,
		_register_addPair__1,
		_register_addPair__2,
		_register_sum__0_inline,
		_register_sum__1,
		_register_sum__2,
	} {
		r(&gf)
	}
	_addPair__0 := gf.Get("addPair", generic.Types(new(Int))).(func(Int, Int) Int)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_addPair__0(34, 56)
	}
}

func BenchmarkDirectInner(b *testing.B) {
	var gf generic.Funcs
	for _, r := range []func(*generic.Funcs){
		_register_addPair__0,
		_register_addPair__1,
		_register_addPair__2,
		_register_sum__0_inline,
		_register_sum__1,
		_register_sum__2,
	} {
		r(&gf)
	}
	_addPair__0 := gf.Get("addPair", generic.Types(new(Int))).(func(Int, Int) Int)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_addPair__0(34, 56)
	}
}

func BenchmarkReuse(b *testing.B) {
	var gf generic.Funcs
	for _, r := range []func(*generic.Funcs){
		_register_addPair__0,
		_register_addPair__1,
		_register_addPair__2,
		_register_sum__0,
		_register_sum__1,
		_register_sum__2,
	} {
		r(&gf)
	}
	_addPair__0 := gf.Get("addPair", generic.Types(new(Int))).(func(Int, Int) Int)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_addPair__0(34, 56)
	}
}

func BenchmarkGenericImpl(b *testing.B) {
	var gf generic.Funcs
	for _, r := range []func(*generic.Funcs){
		_register_addPair__generic(generic.Types(new(Int))),
		_register_addPair__generic(generic.Types(new(Flag))),
		_register_addPair__generic(generic.Types(new(Str))),
		_register_sum__generic(generic.Types(new(Int))),
		_register_sum__generic(generic.Types(new(Flag))),
		_register_sum__generic(generic.Types(new(Str))),
	} {
		r(&gf)
	}
	_addPair__0 := gf.Get("addPair", generic.Types(new(Int))).(func(Int, Int) Int)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_addPair__0(34, 56)
	}
}
