package generic

import (
	"fmt"
)

type Type struct{}

type Funcs struct {
	funcs map[string]map[TypeTuple]interface{}
}

func (gf *Funcs) Add(name string, types TypeTuple, inst interface{}) {
	if gf.funcs == nil {
		gf.funcs = make(map[string]map[TypeTuple]interface{})
	}
	f := gf.funcs[name]
	if f == nil {
		f = make(map[TypeTuple]interface{})
		gf.funcs[name] = f
	}
	if _, ok := f[types]; ok {
		panic(fmt.Errorf("function %s%s registered twice", name, types))
	}
	f[types] = inst
}

// Get gets an instance of the function with the given name and
// type parameters.
func (gf *Funcs) Get(name string, types TypeTuple) interface{} {
	f := gf.funcs[name]
	if f == nil {
		panic("no function found")
	}
	inst := f[types]
	if inst == nil {
		panic(fmt.Errorf("no implementation found for %s%s", name, types))
	}
	return inst
}
