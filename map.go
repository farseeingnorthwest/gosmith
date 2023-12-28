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

func MapKeyValues[K comparable, U, V any](f func(K, U) V, m map[K]U) []V {
	if m == nil {
		return nil
	}

	vs := make([]V, 0, len(m))
	for k, v := range m {
		vs = append(vs, f(k, v))
	}

	return vs
}
