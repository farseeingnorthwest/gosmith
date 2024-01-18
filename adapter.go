package gosmith

type Never struct{}

func Void[T any, U any](f func(T) U) func(T) {
	return func(t T) {
		f(t)
	}
}

func Return[T any](t T) func() T {
	return func() T {
		return t
	}
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func Must1[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func Zero[T any]() T {
	var t T
	return t
}

func Ref[T any](t T) *T {
	return &t
}

func Compose[T any, U any, V any](f func(T) U, g func(U) V) func(T) V {
	return func(t T) V {
		return g(f(t))
	}
}
