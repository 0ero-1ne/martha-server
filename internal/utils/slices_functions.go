package utils

func Filter[T any](slice []T, test func(T) bool) (result []T) {
	for _, value := range slice {
		if test(value) {
			result = append(result, value)
		}
	}
	return
}
