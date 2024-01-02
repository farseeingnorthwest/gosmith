package gosmith

func FilterSlice[T any](xs []T, f func(T) bool) []T {
	if xs == nil {
		return nil
	}

	ys := make([]T, 0, len(xs))
	for _, x := range xs {
		if f(x) {
			ys = append(ys, x)
		}
	}

	return ys
}

func MapSlice[T, U any](xs []T, f func(T) U) []U {
	if xs == nil {
		return nil
	}

	ys := make([]U, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}

	return ys
}

func FlatMapSlice[T, U any](xs []T, f func(T) []U) []U {
	if xs == nil {
		return nil
	}

	ys := make([]U, 0, len(xs))
	for _, x := range xs {
		ys = append(ys, f(x)...)
	}

	return ys
}

func MapFilterSlice[T any, U comparable](xs []T, f func(T) U) []U {
	if xs == nil {
		return nil
	}

	var z U
	ys := make([]U, 0, len(xs))
	for _, x := range xs {
		if y := f(x); y != z {
			ys = append(ys, y)
		}
	}

	return ys
}

func MapSliceResult[T, U any](xs []T, f func(T) (U, error)) ([]U, error) {
	if xs == nil {
		return nil, nil
	}

	ys := make([]U, len(xs))
	for i, x := range xs {
		if y, err := f(x); err != nil {
			return nil, err
		} else {
			ys[i] = y
		}
	}

	return ys, nil
}

func Unique[T comparable](xs []T) []T {
	if xs == nil {
		return nil
	}

	ys := make([]T, 0, len(xs))
	m := make(map[T]Never, len(xs))
	for _, x := range xs {
		if _, ok := m[x]; !ok {
			m[x] = Never{}
			ys = append(ys, x)
		}
	}

	return ys
}

func IndexBy[T any, U comparable](xs []T, f func(T) U) map[U]T {
	if xs == nil {
		return nil
	}

	m := make(map[U]T, len(xs))
	for _, x := range xs {
		m[f(x)] = x
	}

	return m
}

func All[T any](xs []T, f func(T) bool) bool {
	for _, x := range xs {
		if !f(x) {
			return false
		}
	}

	return true
}

func Some[T any](xs []T, f func(T) bool) bool {
	for _, x := range xs {
		if f(x) {
			return true
		}
	}

	return false
}
