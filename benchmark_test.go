package main

import (
	"testing"

	"github.com/rogpeppe/genericdemo/generic"
)

func BenchmarkInline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addPair_Int_inline(34, 56)
	}
}

func BenchmarkDirect(b *testing.B) {
	var gf generic.Funcs
	for _, r := range []func(*generic.Funcs){
		register_addPair_Int_inline,
		register_addPair_Flag,
		register_addPair_Str,
		register_sum_Int_inline,
		register_sum_Flag,
		register_sum_Str,
	} {
		r(&gf)
	}
	_addPair_Int := gf.Get("addPair", generic.Types(new(Int))).(func(Int, Int) Int)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_addPair_Int(34, 56)
	}
}

func BenchmarkDirectInner(b *testing.B) {
	var gf generic.Funcs
	for _, r := range []func(*generic.Funcs){
		register_addPair_Int,
		register_addPair_Flag,
		register_addPair_Str,
		register_sum_Int_inline,
		register_sum_Flag,
		register_sum_Str,
	} {
		r(&gf)
	}
	_addPair_Int := gf.Get("addPair", generic.Types(new(Int))).(func(Int, Int) Int)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_addPair_Int(34, 56)
	}
}

func BenchmarkReuse(b *testing.B) {
	var gf generic.Funcs
	for _, r := range []func(*generic.Funcs){
		register_addPair_Int,
		register_addPair_Flag,
		register_addPair_Str,
		register_sum_Int,
		register_sum_Flag,
		register_sum_Str,
	} {
		r(&gf)
	}
	_addPair_Int := gf.Get("addPair", generic.Types(new(Int))).(func(Int, Int) Int)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_addPair_Int(34, 56)
	}
}

func BenchmarkGenericImpl(b *testing.B) {
	var gf generic.Funcs
	for _, r := range []func(*generic.Funcs){
		register_addPair_generic(generic.Types(new(Int))),
		register_addPair_generic(generic.Types(new(Flag))),
		register_addPair_generic(generic.Types(new(Str))),
		register_sum_generic(generic.Types(new(Int))),
		register_sum_generic(generic.Types(new(Flag))),
		register_sum_generic(generic.Types(new(Str))),
	} {
		r(&gf)
	}
	_addPair_Int := gf.Get("addPair", generic.Types(new(Int))).(func(Int, Int) Int)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_addPair_Int(34, 56)
	}
}
