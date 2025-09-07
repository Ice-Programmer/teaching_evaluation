package utils

func MapStructList[T any, R any](list []T, mapper func(T) R) []R {
	result := make([]R, 0, len(list))
	for _, v := range list {
		result = append(result, mapper(v))
	}
	return result
}

// Diff 返回 source 中有但 exclude 中没有的元素
func Diff[T comparable](source, exclude []T) []T {
	m := make(map[T]struct{}, len(exclude))
	for _, v := range exclude {
		m[v] = struct{}{}
	}

	var diff []T
	for _, v := range source {
		if _, ok := m[v]; !ok {
			diff = append(diff, v)
		}
	}
	return diff
}

func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
