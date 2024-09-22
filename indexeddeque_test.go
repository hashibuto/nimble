package nimble

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexedDequeBasic(t *testing.T) {
	idq := NewIndexedDequeue("s", "test", "hello", "hello world")
	idq.Push("rhino")

	assert.Equal(t, 5, idq.Size())

	v := idq.Pop()
	assert.Equal(t, "s", v)

	idq.Push("q")

	items := idq.Find("h")
	assert.Equal(t, 3, len(items))

	items = idq.Find("he")
	assert.Equal(t, 2, len(items))

	idq.RemoveItem(items...)

	assert.Equal(t, 3, idq.Size())

	for idq.Size() > 0 {
		idq.Pop()
	}

	assert.Nil(t, idq.head)
	assert.Nil(t, idq.tail)
}
