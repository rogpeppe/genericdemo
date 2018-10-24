package main

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestBasic(t *testing.T) {
	c := qt.New(t)
	c.Assert(addPair_generic(_Types{_IntType}).Interface().(func(Int, Int) Int)(34, 56), qt.Equals, Int(34+56))
	c.Assert(addPair_generic(_Types{_StrType}).Interface().(func(Str, Str) Str)("hello ", "world"), qt.Equals, Str("hello world"))
}

func BenchmarkAddPair(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addPair_generic(_Types{_IntType}).Interface().(func(Int, Int) Int)(34, 56)
	}
}
