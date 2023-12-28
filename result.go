package gosmith

func ResultFunc[T, U any](f func(T) U) func(T, error) (U, error) {
	return func(t T, e error) (u U, err error) {
		if err = e; err == nil {
			u = f(t)
		}

		return
	}
}

func ZeroError[T any](t T, err error) T {
	if err != nil {
		return Zero[T]()
	}

	return t
}
