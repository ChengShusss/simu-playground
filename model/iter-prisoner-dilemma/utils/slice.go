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
