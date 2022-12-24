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
