package utils

func MapStructList[T any, R any](list []T, mapper func(T) R) []R {
	result := make([]R, 0, len(list))
	for _, v := range list {
		result = append(result, mapper(v))
	}
	return result
}
