package nimble

type Set[T comparable] struct {
	collection map[T]struct{}
}

// NewSet creates a new set, initialized with values
func NewSet[T comparable](values ...T) *Set[T] {
	collection := map[T]struct{}{}
	for _, v := range values {
		collection[v] = struct{}{}
	}
	return &Set[T]{
		collection: collection,
	}
}

func NewSetFromMap[T comparable, X any](values map[T]X) *Set[T] {
	collection := map[T]struct{}{}
	for v := range values {
		collection[v] = struct{}{}
	}
	return &Set[T]{
		collection: collection,
	}
}

// Add adds a single value to the set
func (s *Set[T]) Add(value T) {
	s.collection[value] = struct{}{}
}

// Has returns true if the set already contains the provided value
func (s *Set[T]) Has(value T) bool {
	_, exists := s.collection[value]
	return exists
}

// Remove removes the specified item from the set
func (s *Set[T]) Remove(value T) {
	delete(s.collection, value)
}

// Size returns the size of the set
func (s *Set[T]) Size() int {
	return len(s.collection)
}

// Items returns a slice of the items in the set
func (s *Set[T]) Items() []T {
	items := make([]T, len(s.collection))
	idx := 0
	for item := range s.collection {
		items[idx] = item
		idx++
	}

	return items
}
