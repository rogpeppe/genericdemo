package main

type Int int

func (i Int) Add(j Int) Int {
	return i + j
}

type Str string

func (s Str) Add(t Str) Str {
	return s + t
}

type Flag struct {
	Mask int
}

func (f Flag) Add(g Flag) Flag {
	return Flag{f.Mask | g.Mask}
}
