// +build ignore

package main

func main() {
	fmt.Println(addPair(Int)(34, 56))
	fmt.Println(addPair(Str)("a", "b"))
	fmt.Println(addPair(Flag)(Flag{2}, Flag{3})
	x := sum(sum(sum(Vec(Vec(Vec(Int))){
		{
			{1,2,3,4},
			{3,4},
		},
		{
			{2,3},
			{},
			{4,5,7,787}
		},
	})))
	fmt.Println(x)
}

type Adder(type T) interface {
	Add() T
}

contract AdderC(t T) {
	Adder(T)(t)
}

func addPair(type T AdderC)(a, b T) T {
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

type Vec(type T adder) []T

func (v1 Vec(T)) Add(v2 Vec(T)) Vec(T) {
	if len(v2) > len(v1) {
		v1, v2 = v2, v1
	}
	r := make(Vec(T), len(v1))
	for i, x := range v2 {
		r[i] = v1[i].Add(v2[i])
	}
	for i, x := range v1[len(v2):] {
		r[i+len(v2)] = x
	}
	return r
}
