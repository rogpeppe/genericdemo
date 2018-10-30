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
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: gengen package name type...\n")
		os.Exit(2)
	}
	flag.Parse()
	if flag.NArg() < 3 {
		flag.Usage()
	}
	pkgName := "github.com/genericdemo/heap" // flag.Arg(0)
	instName := "sum"		// flag.Arg(1)
	instTypes := []string{"Int"} // flag.Args()[2:]

	cfg := packages.Config{
		Mode: packages.LoadSyntax,
		Dir:  "/home/rog/src/s/genericdemo/heap",
	}
	pkgs, err := packages.Load(&cfg, pkgName)
	if err != nil {
		log.Printf("%s\n", err)
		os.Exit(1)
	}
	fmt.Printf("ok: %#v", pkgs)
	pkg := pkgs[0]
	for _, err := range pkg.Errors {
		fmt.Printf("error: %s (%v)\n", err.Msg, err.Kind)
	}
	pkg.Types.Scope().Lookup(name)
	
	//prog, ssaPkgs := ssautil.Packages(pkgs, ssa.NaiveForm)
	//fmt.Printf("prog: %#v\n", prog)
	//fmt.Printf("pkgs: %#v\n", ssaPkgs)
	// astutil.PathEnclosingInterval(root *ast.File, start, end token.Pos) (path []ast.Node, exact bool)
	look up _types
	find rtequiv type for new type

	after := []func(*astutil.Cursor) bool
	var definition xxxx

	astutil.Apply(f, func(c *astutil.Cursor) bool {
		if we're at definition of name {
			definition = details of definition
			rewritingEquivTypes = true
			after = push(after, func(c *astutil.Cursor) bool {
				definition = nil
				change NAME to NAME_EQUIVTYPENAMES
			}
		} else {
			after = append(after, nopApply)
		}
		if rewritingEquivTypes && isLocalTypeName(c) {
			c.Replace(replacement)
		}
	}, func(c *astutil.Cursor) bool {
		n := len(after)
		after[n-1](c)
		after = after[0:n-1]
	})
}

func nopApply(c *astutil.Cursor) bool {
	return true
}

for function:

type _ptype_NAME__pINDEX generic.Type

var _inst_NAME__LOCALTYPESID = _inst_NAME__EQUIVID{
	func_FUNCNAME: (func(LOCALTYPE1, []LOCALTYPE2) *LOCALTYPE3)(unsafe.Pointer(FUNCNAME__LOCALTYPESID)),
	method_TYPENAME__METHODNAME: (func(....LOCALTYPE))(unsafe.Pointer(TYPENAME__LOCALTYPESID_METHODNAME))
	type_LOCALTYPE: reflect.TypeOf(...),
}

type _inst_NAME struct {
	func_NAME func(LOCALTYPE1 etc)
	method_TYPENAME__METHODNAME: func(etc),
	type_ID LOCALTYPE
}

func NAME(_inst *_inst_NAME, LOCALTYPE...) LOCALTYPE... {
	etc. LOCALTYPE
}

func NAME__LOCALTYPESID(... LOCALTYPE) (... LOCALTYPE)


for type:

type _inst_NAME struct {
	func_NAME func(LOCALTYPE1 etc)
	method_TYPENAME__METHODNAME: func(etc),
	type_ID LOCALTYPE
}

type NAME__LOCALTYPESID structetc LOCALTYPE

func (r *NAME__LOCALTYPESID) METHOD(LOCALTYPE) LOCALTYPE {
	implementation
	LOCALTYPE
}

etc

------

instantiate NAME TYPE...

look up _types definition, and read types into slice

determine runtime-equivalent type for TYPE...

apply ast
	within NAME definition:
		if it's a function or a type, in the body:
			when there's a name with a "_ptype_NAME__p" prefix.
			check that it refers to a generic.Type type.
			if so, replace node by equivalent instantiated type

		change NAME to NAME_EQUIVTYPE
		
	within NAME.method:
		replace "_ptype_NAME__p*" references by equivalent types.
		change method definition to function definition with extra _inst arg.

	within _inst_NAME definition:
		instantiate type names


for wrapper:

add types to 

apply ast
	within NAME
		save signature, with _inst argument removed
	within NAME__METHOD
		translate to NAME.METHOD

for each saved signature,
	generate:
		func NAME(instantiate(sig, types))
	generate:
		var _inst_NAME__INDEX = _inst_NAME

