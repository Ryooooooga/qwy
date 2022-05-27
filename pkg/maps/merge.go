package maps

func Merge[K comparable, V any](a ...map[K]V) map[K]V {
	cap := 0
	for _, m := range a {
		cap += len(m)
	}

	dest := make(map[K]V, cap)
	for _, m := range a {
		for k, v := range m {
			dest[k] = v
		}
	}
	return dest
}
