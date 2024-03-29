package gosmith

func Keys[K comparable, V any](m map[K]V) []K {
	if m == nil {
		return nil
	}

	ks := make([]K, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

func Values[K comparable, V any](m map[K]V) []V {
	if m == nil {
		return nil
	}

	vs := make([]V, 0, len(m))
	for _, v := range m {
		vs = append(vs, v)
	}

	return vs
}

func CollectKeyValues[K comparable, U, V any](m map[K]U, f func(K, U) V) []V {
	if m == nil {
		return nil
	}

	vs := make([]V, 0, len(m))
	for k, v := range m {
		vs = append(vs, f(k, v))
	}

	return vs
}

func CollectFilterKeyValues[K comparable, U any, V comparable](m map[K]U, f func(K, U) V) []V {
	if m == nil {
		return nil
	}

	var z V
	ys := make([]V, 0, len(m))
	for k, v := range m {
		if y := f(k, v); y != z {
			ys = append(ys, y)
		}
	}

	return ys
}

func MapKeyValues[K comparable, U, V any](m map[K]U, f func(K, U) V) map[K]V {
	if m == nil {
		return nil
	}

	n := make(map[K]V, len(m))
	for k, v := range m {
		n[k] = f(k, v)
	}

	return n
}

func MapFilterKeyValues[K comparable, U any, V comparable](m map[K]U, f func(K, U) V) map[K]V {
	if m == nil {
		return nil
	}

	var z V
	n := make(map[K]V, len(m))
	for k, v := range m {
		if y := f(k, v); y != z {
			n[k] = y
		}
	}

	return n
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

func IndexFilterBy[T any, U comparable](xs []T, f func(T) U) map[U]T {
	if xs == nil {
		return nil
	}

	var z U
	m := make(map[U]T, len(xs))
	for _, x := range xs {
		i := f(x)
		if i != z {
			m[i] = x
		}
	}

	return m
}

func IndexKeyValues[T any, K comparable, V any](xs []T, f func(T) (K, V)) map[K]V {
	if xs == nil {
		return nil
	}

	m := make(map[K]V)
	for _, x := range xs {
		k, v := f(x)
		m[k] = v
	}

	return m
}

func IndexFilterKeyValues[T any, K, V comparable](xs []T, f func(T) (K, V)) map[K]V {
	if xs == nil {
		return nil
	}

	var z V
	m := make(map[K]V)
	for _, x := range xs {
		k, v := f(x)
		if v != z {
			m[k] = v
		}
	}

	return m
}

func CollectBy[T any, U comparable](xs []T, f func(T) U) map[U][]T {
	if xs == nil {
		return nil
	}

	m := make(map[U][]T)
	for _, x := range xs {
		i := f(x)
		m[i] = append(m[i], x)
	}

	return m
}

func CollectFilterBy[T any, U comparable](xs []T, f func(T) U) map[U][]T {
	if xs == nil {
		return nil
	}

	var z U
	m := make(map[U][]T)
	for _, x := range xs {
		i := f(x)
		if i != z {
			m[i] = append(m[i], x)
		}
	}

	return m
}
