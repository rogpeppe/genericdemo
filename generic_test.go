package main

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestSanity(t *testing.T) {
	c := qt.New(t)
	var gf GenericFuncs
	for _, r := range []func(*GenericFuncs){
		register_addPair_Int,
		register_addPair_Flag,
		register_addPair_Str,
		register_sum_Int,
		register_sum_Flag,
		register_sum_Str,
	} {
		r(&gf)
	}
	_addPair_Int := gf.Get("addPair", Types(new(Int))).(func(Int, Int) Int)
	_addPair_Flag := gf.Get("addPair", Types(new(Flag))).(func(Flag, Flag) Flag)
	_addPair_Str := gf.Get("addPair", Types(new(Str))).(func(Str, Str) Str)

	c.Assert(_addPair_Int(34, 67), qt.Equals, Int(34+67))
	c.Assert(_addPair_Flag(Flag{1}, Flag{3}), qt.Equals, Flag{3})
	c.Assert(_addPair_Str(Str("hello "), Str("world")), qt.Equals, Str("hello world"))
}

func TestGenericImpl(t *testing.T) {
	c := qt.New(t)
	var gf GenericFuncs
	for _, r := range []func(*GenericFuncs){
		register_addPair_generic(Types(new(Int))),
		register_addPair_generic(Types(new(Flag))),
		register_addPair_generic(Types(new(Str))),
		register_addPair_generic(Types(new(AdderI))),
		register_sum_generic(Types(new(Int))),
		register_sum_generic(Types(new(Flag))),
		register_sum_generic(Types(new(Str))),
		register_sum_generic(Types(new(AdderI))),
	} {
		r(&gf)
	}
	_addPair_Int := gf.Get("addPair", Types(new(Int))).(func(Int, Int) Int)
	_addPair_Flag := gf.Get("addPair", Types(new(Flag))).(func(Flag, Flag) Flag)
	_addPair_Str := gf.Get("addPair", Types(new(Str))).(func(Str, Str) Str)
	_addPair_Adder := gf.Get("addPair", Types(new(AdderI))).(func(AdderI, AdderI) AdderI)

	c.Assert(_addPair_Int(34, 67), qt.Equals, Int(34+67))
	c.Assert(_addPair_Flag(Flag{1}, Flag{3}), qt.Equals, Flag{3})
	c.Assert(_addPair_Str(Str("hello "), Str("world")), qt.Equals, Str("hello world"))
	c.Assert(_addPair_Adder(StrAdderI{"hello "}, StrAdderI{"world"}), qt.Equals, StrAdderI{"hello world"})
}
