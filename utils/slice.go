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

func GroupBy[T any, K comparable, V any](items []T, keySelector func(T) K, valueSelector func(T) V) map[K][]V {
	result := make(map[K][]V, len(items))
	for _, item := range items {
		key := keySelector(item)
		value := valueSelector(item)
		result[key] = append(result[key], value)
	}
	return result
}

func DistinctStringArray(slice []string) []string {
	if len(slice) == 0 {
		return slice
	}

	seen := make(map[string]bool)
	result := make([]string, 0, len(slice))

	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}

func ToMap[T any, K comparable, V any](
	items []T,
	keySelector func(T) K,
	valueSelector func(T) V,
) map[K]V {
	result := make(map[K]V, len(items))
	for _, item := range items {
		key := keySelector(item)
		value := valueSelector(item)
		result[key] = value
	}
	return result
}

func DistinctIntArray(slice []int) []int {
	if len(slice) == 0 {
		return slice
	}

	seen := make(map[int]bool)
	result := make([]int, 0, len(slice))

	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}
