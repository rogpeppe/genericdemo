// +build ignore

package main

func main() {
	fmt.Println(foo(Int)(34, 56))
	fmt.Println(foo(Str)("a", "b"))
	fmt.Println(foo(Flag)(Flag{2}, Flag{3})
}

type Adder(type T) interface {
	Add() T
}

contract AdderC(t T) {
	Adder(T)(t)
}

func foo(type T AdderC)(a, b T) T {
	f := sum(T)
	return f([]T{a, b})
}

func sum(type T AdderC)(ts []T) (x T) {
	if len(ts) == 0 {
		return
	}
	x = ts[0]
	for _, t := range ts[1:] {
		x = x.Add(t)
	}
	return x
}
