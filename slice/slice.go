package slice

// 找出在 a 切片中存在但在 b 切片中不存在的元素
// 例如：a = [1,2,3,4,5], b = [3,4,5,6,7], 返回 [1,2]
func SliceDiff[T comparable](a, b []T) []T {
	bSet := make(map[T]struct{})
	for _, item := range b {
		bSet[item] = struct{}{}
	}

	var result []T
	for _, item := range a {
		if _, found := bSet[item]; !found {
			result = append(result, item)
		}
	}
	return result
}

// 找出在 a 切片和 b 切片中的交集
func SliceIntersection[T comparable](a, b []T) []T {
	bSet := make(map[T]struct{})
	for _, item := range b {
		bSet[item] = struct{}{}
	}

	var result []T
	for _, item := range a {
		if _, found := bSet[item]; found {
			result = append(result, item)
		}
	}
	return result
}

// Map 转换slice
func Map[S any, T any](source []S, m func(S) T) []T {
	res := make([]T, 0, len(source))
	for _, item := range source {
		res = append(res, m(item))
	}
	return res
}

// ToMap 转换slice到map
func ToMap[S any, K comparable, V any](source []S, m func(S) (K, V)) map[K]V {
	res := map[K]V{}
	for _, item := range source {
		k, v := m(item)
		res[k] = v
	}
	return res
}
