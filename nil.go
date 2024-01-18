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

func OrElse[T comparable](t T, fs ...func() T) T {
	for _, f := range fs {
		if !IsZero(t) {
			break
		}

		t = f()
	}

	return t
}

func IsZero[T comparable](t T) bool {
	return t == Zero[T]()
}
