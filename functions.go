package nimble

// Unique returns a slice from the original containing only unique elements
func Unique[T comparable](values ...T) []T {
	filtered := []T{}
	s := NewSet[T]()
	for _, v := range values {
		if !s.Has(v) {
			filtered = append(filtered, v)
			s.Add(v)
		}
	}

	return filtered
}

// Filter returns the filtered items from the input set, using the filter function
func Filter[T comparable](f func(v T) bool, values ...T) []T {
	filtered := []T{}
	for _, v := range values {
		if f(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

// Map returns the original slice passed through the mapping function
func Map[T any, U any](f func(v T) U, values ...T) []U {
	output := make([]U, len(values))
	for idx, v := range values {
		output[idx] = f(v)
	}

	return output
}
