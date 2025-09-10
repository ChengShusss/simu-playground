package utils

func Sum[T int | float64](arr []T) T {
	total := *new(T)

	for _, ele := range arr {
		total += ele
	}

	return total
}

func Avg[T int | float64](arr []T) float64 {
	if len(arr) == 0 {
		return 0.0
	}

	return float64(Sum(arr)) / float64(len(arr))
}

func ToSlice[T, K int | float64 | string](in map[K]T) []T {
	s := make([]T, 0, len(in))
	for _, v := range in {
		s = append(s, v)
	}
	return s
}

func Contain[T int | float64 | string](s []T, ele T) bool {
	for _, k := range s {
		if k == ele {
			return true
		}
	}

	return false
}
