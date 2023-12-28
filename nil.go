package gosmith

func MapNullable[T, U any](t *T, f func(*T) *U) *U {
	if t == nil {
		return nil
	}

	return f(t)
}

func OrZero[T any](t *T) T {
	if t == nil {
		return Zero[T]()
	}

	return *t
}

func OrElse[T any](t *T, f func() T) T {
	if t == nil {
		return f()
	}

	return *t
}
