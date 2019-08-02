// +build ignore

func Foo(type T)(t T, i int) T {
	if i <= 0 {
		return t
	}
	return Bar(T)(t, i-1)
}

func Bar(type T)(t T, i int) T {
	if i <= 0 {
		return t
	}
	return Foo(T)(t, i-1)
}
