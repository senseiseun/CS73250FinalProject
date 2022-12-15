package LinkedList

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	list := New[int]()
	for i := 1; i <= 1000; i++ {
		list.Push(i)
	}
	assert.Equal(t, 1000, list.GetSize())
	assert.Equal(t, 1000, list.Get(999))
	assert.Equal(t, 1, list.Get(0))
	assert.Equal(t, 500, list.Get(499))

	for i := 999; i >= 500; i-- {
		list.Remove(i)
	}

	assert.Equal(t, 500, list.GetSize())
	list.Insert(100, 0)
	list.Insert(40, 50)
	list.Insert(126, 126)
	assert.Equal(t, 100, list.Get(0))
	assert.Equal(t, 40, list.Get(50))
	assert.Equal(t, 126, list.Get(126))

	iter := list.GetIter()
	iter.Start()
	for !iter.Done() {
		curr := iter.Get()
		fmt.Println(curr)
		iter.Next()
	}
}
