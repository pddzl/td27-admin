package utils

// IsContain slice contain
func IsContain(list []string, item string) bool {
	for _, value := range list {
		if value == item {
			return true
		}
	}
	return false
}
