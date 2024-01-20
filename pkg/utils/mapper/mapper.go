package mapper

func Map[T, U any](a []T, f func(T) U) []U {
	res := make([]U, len(a))
	for i := range a {
		res[i] = f(a[i])
	}
	return res
}
