package utils

func Filter[T any](arr []T, filter func(T) bool) []T {
	var nextArr []T
	for _, item := range arr {
		if filter(item) {
			nextArr = append(nextArr, item)
		}
	}

	return nextArr
}
