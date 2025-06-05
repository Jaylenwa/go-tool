package slice

// SliceDiff returns a new slice containing elements that are in slice 'a' but not in slice 'b'.
// Both slices must contain comparable elements.
// Example: a = [1,2,3,4,5], b = [3,4,5,6,7] => returns [1,2]
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

// SliceIntersection returns a new slice containing elements that are present in both slice 'a' and slice 'b'.
// Both slices must contain comparable elements.
// Example: a = [1,2,3,4,5], b = [3,4,5,6,7] => returns [3,4,5]
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

// Map transforms each element of the source slice using the provided mapping function.
// It returns a new slice containing the transformed elements.
// The mapping function takes an element of type S and returns an element of type T.
// Example: source = [1,2,3], m = func(x int) string { return strconv.Itoa(x) } => returns ["1","2","3"]
func Map[S any, T any](source []S, m func(S) T) []T {
	res := make([]T, 0, len(source))
	for _, item := range source {
		res = append(res, m(item))
	}
	return res
}

// ToMap converts a slice to a map using the provided mapping function.
// The mapping function takes an element of type S and returns a key-value pair (K,V).
// Returns a map where keys are of type K and values are of type V.
// Example: source = [{id:1,name:"a"}, {id:2,name:"b"}], m = func(x Item) (int,string) { return x.id, x.name }
// => returns map[int]string{1:"a", 2:"b"}
func ToMap[S any, K comparable, V any](source []S, m func(S) (K, V)) map[K]V {
	res := map[K]V{}
	for _, item := range source {
		k, v := m(item)
		res[k] = v
	}
	return res
}
