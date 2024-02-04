package utils

// IsContain slice contain
func IsContain[T comparable](list []T, item T) bool {
	for _, value := range list {
		if value == item {
			return true
		}
	}
	return false
}
