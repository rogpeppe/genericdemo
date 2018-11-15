package generic

import (
	"reflect"
	"testing"
	"fmt"
	"go/types"

	qt "github.com/frankban/quicktest"
)

var typeMapTests = []struct {
	about  string
	t      interface{}
	expect string
}{{
	about:  "single int",
	t:      new(int64),
	expect: "v8a8",
}, {
	about:  "interface",
	t:      new(interface{}),
	expect: "pp",
}, {
	about: "struct",
	t: new(struct {
		_ [3]byte
		_ *int
		_ [1]byte
		_ *struct{}
	}),
	expect: "v8pv8p",
}, {
	about:  "array",
	t:      new([4]*int),
	expect: "pppp",
}, {
	about:  "slice",
	t:      new([]int),
	expect: "pv16",
}, {
	about:  "value array",
	t:      new([3]byte),
	expect: "v3a1",
}, {
	about: "short int",
	t: new([4]int16),
	expect: "v8a2",
}}

func TestTypeMap(t *testing.T) {
	if ptrSize != 8 {
		t.Skip("type map tests only work on 64 bit architectures")
	}
	c := qt.New(t)
	for _, test := range typeMapTests {
		test := test
		c.Run(test.about, func(c *qt.C) {
			typ := reflect.TypeOf(test.t).Elem()
			c.Assert(typeMapOf(typ).typeName(), qt.Equals, test.expect)
		})
	}
}

var typesTypeTests = []struct{
	t interface{}
	want string
}{{
	t: new(int),
	want: "int",
}, {
	t: new(struct{
		a int
		b int
		c []byte
	}),
	want: "struct{a int; b int; c []uint8}",
}, {
	t: new(interface{
		A(string, []bool) (int, string)
		B()
	}),
	want: "interface{A(string, []bool) (int, string); B()}",
}}

func TestTypesType(t *testing.T) {
	for _, test := range typesTypeTests {
		t.Run(fmt.Sprintf("%T", test.t), func(t *testing.T) {
			typ := typesType(reflect.TypeOf(test.t).Elem())
			got := types.TypeString(typ, nil)
			if got != test.want {
				t.Fatalf("bad type; got %s want %s", got, test.want)
			}
		})
	}
}
