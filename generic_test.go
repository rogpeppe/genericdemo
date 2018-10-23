package main

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestSanity(t *testing.T) {
	c := qt.New(t)
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

	c.Assert(_foo_Int(34, 67), qt.Equals, Int(34+67))
	c.Assert(_foo_Flag(Flag{1}, Flag{3}), qt.Equals, Flag{3})
	c.Assert(_foo_Str(Str("hello "), Str("world")), qt.Equals, Str("hello world"))
}
