package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

func main() {
	cfg := packages.Config{
		Mode: packages.LoadSyntax,
		Dir:  "/home/rog/src/s/genericdemo/heap",
	}
	pkgs, err := packages.Load(&cfg, "github.com/genericdemo/heap")
	if err != nil {
		log.Printf("%s\n", err)
		os.Exit(1)
	}
	fmt.Printf("ok: %#v", pkgs)
	pkg := pkgs[0]
	for _, err := range pkg.Errors {
		fmt.Printf("error: %s (%v)\n", err.Msg, err.Kind)
	}
	prog, ssaPkgs := ssautil.Packages(pkgs, ssa.NaiveForm)
	fmt.Printf("prog: %#v\n", prog)
	fmt.Printf("pkgs: %#v\n", ssaPkgs)
}

outer transformation:

var _inst_{{.Name}}_{{.TPIndex}} = _inst_{{.Name}}_{{.Equiv}}{

}

go through function, looking for
- method calls on type parameters
- generic function calls using type parameters.
- type instantations using type parameters (disallowed for now except converting to interface)
- conversions of type parameter-derived types.

OR

look through contract, determining what operations are allowed.
then generate an instance that allows all of those.
However, that doesn't take into account generic function
and type instantiations, so we'll need to go through
and to find those anyway.

each one of the above gets turned into an entry in the instance.

for each call, we want to look up the method and find
its signature

func instantiateFunc(name string, with []types.Type) *ast.File {

	find equivalent type for type parameters 
	astutil.Apply(file, f (c *
	
	if receiver is generic type parameter,
	change