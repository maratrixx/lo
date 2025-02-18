package lo

// Contains returns true if an element is present in a collection.
func Contains[T comparable](collection []T, element T) bool {
	for _, item := range collection {
		if item == element {
			return true
		}
	}

	return false
}

// ContainsBy returns true if predicate function return true.
func ContainsBy[T any](collection []T, predicate func(T) bool) bool {
	for _, item := range collection {
		if predicate(item) {
			return true
		}
	}

	return false
}

// Every returns true if all elements of a subset are contained into a collection.
func Every[T comparable](collection []T, subset []T) bool {
	for _, elem := range subset {
		if !Contains(collection, elem) {
			return false
		}
	}

	return true
}

// Some returns true if at least 1 element of a subset is contained into a collection.
func Some[T comparable](collection []T, subset []T) bool {
	for _, elem := range subset {
		if Contains(collection, elem) {
			return true
		}
	}

	return false
}

// Intersect returns the intersection between two collections.
func Intersect[T comparable](list1 []T, list2 []T) []T {
	result := []T{}
	seen := make(map[T]struct{}, len(list1))

	for _, elem := range list1 {
		seen[elem] = struct{}{}
	}

	for _, elem := range list2 {
		if _, ok := seen[elem]; ok {
			result = append(result, elem)
		}
	}

	return result
}

// Difference returns the difference between two collections.
// The first value is the collection of element absent of list2.
// The second value is the collection of element absent of list1.
func Difference[T comparable](list1 []T, list2 []T) ([]T, []T) {
	left := []T{}
	right := []T{}

	seenLeft := make(map[T]struct{}, len(list1))
	seenRight := make(map[T]struct{}, len(list2))

	for _, elem := range list1 {
		seenLeft[elem] = struct{}{}
	}

	for _, elem := range list2 {
		seenRight[elem] = struct{}{}
	}

	for _, elem := range list1 {
		if _, ok := seenRight[elem]; !ok {
			left = append(left, elem)
		}
	}

	for _, elem := range list2 {
		if _, ok := seenLeft[elem]; !ok {
			right = append(right, elem)
		}
	}

	return left, right
}

// Union returns all distinct elements from both collections.
// result returns will not change the order of elements relatively.
func Union[T comparable](list1 []T, list2 []T) []T {
	total := len(list1) + len(list2)
	result := make([]T, 0, total)
	hasAdd := make(map[T]struct{}, total)

	for _, e := range list1 {
		if _, ok := hasAdd[e]; !ok {
			result = append(result, e)
			hasAdd[e] = struct{}{}
		}
	}

	for _, e := range list2 {
		if _, ok := hasAdd[e]; ok {
			continue
		}

		result = append(result, e)
		hasAdd[e] = struct{}{}
	}

	return result
}
