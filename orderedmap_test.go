package nimble

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderedMapInsertion(t *testing.T) {
	om := NewOrderedMap[string, int]()
	om.Put("first", 10)
	om.Put("second", 9)
	om.Put("third", 8)

	assert.Equal(t, om.Length(), 3)
	iter := om.GetIter()
	i := 0
	for iter.Next() {
		switch i {
		case 0:
			assert.Equal(t, "first", iter.Key())
			assert.Equal(t, 10, iter.Value())
		case 1:
			assert.Equal(t, "second", iter.Key())
			assert.Equal(t, 9, iter.Value())
		case 2:
			assert.Equal(t, "third", iter.Key())
			assert.Equal(t, 8, iter.Value())
		}

		i++
	}
}

func TestOrderedMapRemoval(t *testing.T) {
	om := NewOrderedMap[string, int]()
	om.Put("first", 10)
	om.Put("second", 9)
	om.Put("third", 8)

	om.Delete("second")
	assert.Equal(t, 2, om.Length())

	i := 0
	iter := om.GetIter()
	for iter.Next() {
		switch i {
		case 0:
			assert.Equal(t, "first", iter.Key())
			assert.Equal(t, 10, iter.Value())
		case 1:
			assert.Equal(t, "third", iter.Key())
			assert.Equal(t, 8, iter.Value())
		}

		i++
	}
}

func TestFirstNodeRemoval(t *testing.T) {
	om := NewOrderedMap[string, int]()
	om.Put("first", 10)
	om.Put("second", 9)
	om.Put("third", 8)

	om.Delete("first")
	assert.Equal(t, 2, om.Length())

	i := 0
	iter := om.GetIter()
	for iter.Next() {
		switch i {
		case 0:
			assert.Equal(t, "second", iter.Key())
			assert.Equal(t, 9, iter.Value())
		case 1:
			assert.Equal(t, "third", iter.Key())
			assert.Equal(t, 8, iter.Value())
		}

		i++
	}
}

func TestLastNodeRemoval(t *testing.T) {
	om := NewOrderedMap[string, int]()
	om.Put("first", 10)
	om.Put("second", 9)
	om.Put("third", 8)

	om.Delete("third")
	assert.Equal(t, 2, om.Length())

	i := 0
	iter := om.GetIter()
	for iter.Next() {
		switch i {
		case 0:
			assert.Equal(t, "first", iter.Key())
			assert.Equal(t, 10, iter.Value())
		case 1:
			assert.Equal(t, "second", iter.Key())
			assert.Equal(t, 9, iter.Value())
		}

		i++
	}
}

func TestFirstAndLastNodeRemoval(t *testing.T) {
	om := NewOrderedMap[string, int]()
	om.Put("first", 10)
	om.Put("second", 9)
	om.Put("third", 8)

	om.Delete("third")
	om.Delete("first")
	assert.Equal(t, 1, om.Length())

	i := 0
	iter := om.GetIter()
	for iter.Next() {
		switch i {
		case 0:
			assert.Equal(t, "second", iter.Key())
			assert.Equal(t, 9, iter.Value())
		}

		i++
	}
}

func TestAllodeRemoval(t *testing.T) {
	om := NewOrderedMap[string, int]()
	om.Put("first", 10)
	om.Put("second", 9)
	om.Put("third", 8)

	om.Delete("third")
	om.Delete("first")
	om.Delete("second")
	assert.Equal(t, 0, om.Length())

	i := 0
	iter := om.GetIter()
	for iter.Next() {
		i++
	}

	assert.Equal(t, 0, i)

	om.Put("fourth", 7)
	assert.Equal(t, 1, om.Length())

	i = 0
	iter = om.GetIter()
	for iter.Next() {
		switch i {
		case 0:
			assert.Equal(t, "fourth", iter.Key())
			assert.Equal(t, 7, iter.Value())
		}

		i++
	}
}
