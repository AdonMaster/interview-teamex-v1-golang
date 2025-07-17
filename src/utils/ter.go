package utils

func Coalesce[T comparable](list ...T) T {
	var empty T
	for _, element := range list {
		if element != empty {
			return element
		}
	}
	return empty
}
