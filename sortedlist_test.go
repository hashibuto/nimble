package nimble

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedListInsertEven(t *testing.T) {
	sl := NewSortedList[int](4, 1)
	sl.Extend(2)

	assert.Equal(t, 3, sl.Length())
	v := sl.PopSmallest()
	assert.Equal(t, 1, v)
	v = sl.PopSmallest()
	assert.Equal(t, 2, v)
	v = sl.PopSmallest()
	assert.Equal(t, 4, v)
}

func TestSortedListInsertOdd(t *testing.T) {
	sl := NewSortedList[int](4, 1, 2)
	sl.Extend(2)

	assert.Equal(t, 4, sl.Length())
	v := sl.PopSmallest()
	assert.Equal(t, 1, v)
	v = sl.PopSmallest()
	assert.Equal(t, 2, v)
	v = sl.PopSmallest()
	assert.Equal(t, 2, v)
	v = sl.PopSmallest()
	assert.Equal(t, 4, v)
}

func TestSortedListLargestOut(t *testing.T) {
	sl := NewSortedList[int](4, 1, 2)
	sl.Extend(2)

	assert.Equal(t, 4, sl.Length())
	v := sl.PopLargest()
	assert.Equal(t, 4, v)
	v = sl.PopLargest()
	assert.Equal(t, 2, v)
	v = sl.PopLargest()
	assert.Equal(t, 2, v)
	v = sl.PopLargest()
	assert.Equal(t, 1, v)
}

func TestSortedListEmpty(t *testing.T) {
	sl := NewSortedList[int]()
	sl.Extend(2)

	assert.Equal(t, 1, sl.Length())
	v := sl.PopLargest()
	assert.Equal(t, 2, v)
}

func TestSortedListSingle(t *testing.T) {
	sl := NewSortedList[int](1)
	sl.Extend(2)

	assert.Equal(t, 2, sl.Length())
	v := sl.PopLargest()
	assert.Equal(t, 2, v)
	v = sl.PopLargest()
	assert.Equal(t, 1, v)
}
