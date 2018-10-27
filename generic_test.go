package main

import (
	"testing"

	qt "github.com/frankban/quicktest"

	"github.com/rogpeppe/genericdemo/generic"
)

func TestSanity(t *testing.T) {
	c := qt.New(t)
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
	_addPair__1 := gf.Get("addPair", generic.Types(new(Flag))).(func(Flag, Flag) Flag)
	_addPair__2 := gf.Get("addPair", generic.Types(new(Str))).(func(Str, Str) Str)

	c.Assert(_addPair__0(34, 67), qt.Equals, Int(34+67))
	c.Assert(_addPair__1(Flag{1}, Flag{3}), qt.Equals, Flag{3})
	c.Assert(_addPair__2(Str("hello "), Str("world")), qt.Equals, Str("hello world"))
}

func TestGenericImpl(t *testing.T) {
	c := qt.New(t)
	var gf generic.Funcs
	for _, r := range []func(*generic.Funcs){
		_register_addPair__generic(generic.Types(new(Int))),
		_register_addPair__generic(generic.Types(new(Flag))),
		_register_addPair__generic(generic.Types(new(Str))),
		_register_addPair__generic(generic.Types(new(AdderI))),
		_register_sum__generic(generic.Types(new(Int))),
		_register_sum__generic(generic.Types(new(Flag))),
		_register_sum__generic(generic.Types(new(Str))),
		_register_sum__generic(generic.Types(new(AdderI))),
	} {
		r(&gf)
	}
	_addPair__0 := gf.Get("addPair", generic.Types(new(Int))).(func(Int, Int) Int)
	_addPair__1 := gf.Get("addPair", generic.Types(new(Flag))).(func(Flag, Flag) Flag)
	_addPair__2 := gf.Get("addPair", generic.Types(new(Str))).(func(Str, Str) Str)
	_addPair_Adder := gf.Get("addPair", generic.Types(new(AdderI))).(func(AdderI, AdderI) AdderI)

	c.Assert(_addPair__0(34, 67), qt.Equals, Int(34+67))
	c.Assert(_addPair__1(Flag{1}, Flag{3}), qt.Equals, Flag{3})
	c.Assert(_addPair__2(Str("hello "), Str("world")), qt.Equals, Str("hello world"))
	c.Assert(_addPair_Adder(StrAdderI{"hello "}, StrAdderI{"world"}), qt.Equals, StrAdderI{"hello world"})
}
