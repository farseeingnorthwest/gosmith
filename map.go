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

func MapKeyValues[K comparable, U, V any](m map[K]U, f func(K, U) V) []V {
	if m == nil {
		return nil
	}

	vs := make([]V, 0, len(m))
	for k, v := range m {
		vs = append(vs, f(k, v))
	}

	return vs
}

func MapFilterKeyValues[K comparable, U any, V comparable](m map[K]U, f func(K, U) V) []V {
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

func MapValues[K comparable, U, V any](m map[K]U, f func(U) V) map[K]V {
	if m == nil {
		return nil
	}

	n := make(map[K]V, len(m))
	for k, v := range m {
		n[k] = f(v)
	}

	return n
}

func MapFilterValues[K comparable, U any, V comparable](m map[K]U, f func(U) V) map[K]V {
	if m == nil {
		return nil
	}

	var z V
	n := make(map[K]V, len(m))
	for k, v := range m {
		if y := f(v); y != z {
			n[k] = y
		}
	}

	return n
}
