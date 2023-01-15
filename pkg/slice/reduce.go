package slice

func Reduce[T, E any](s []T, fn func(E, T) E, initial E) E {
	reduced := initial
	for _, el := range s {
		reduced = fn(reduced, el)
	}

	return reduced
}
