package LinkedList

import (
	"Program2/Iterator"
)

type Node[T any] struct {
	next *Node[T]
	prev *Node[T]
	data T
}

type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
	iter *Iterator.Iterator[T]
}

func (list *LinkedList[T]) Push(elem T) *Node[T] {
	if list.size == 0 {
		list.head = &Node[T]{data: elem}
		list.tail = list.head
	} else {
		list.tail.next = &Node[T]{data: elem, prev: list.tail}
		list.tail = list.tail.next
	}
	list.size++
	return list.tail
}

func (list *LinkedList[T]) PushFront(elem T) *Node[T] {
	if list.size == 0 {
		list.head = &Node[T]{data: elem}
		list.tail = list.head
	} else {
		list.head.prev = &Node[T]{data: elem, next: list.head}
		list.head = list.head.prev
	}
	list.size++
	return list.head
}

func (list *LinkedList[T]) Insert(elem T, index int) {
	if list.size == 0 {
		list.head = &Node[T]{data: elem}
		list.tail = list.head
	} else {
		currIndex := 0
		var prev *Node[T]
		for curr := list.head; curr != nil; curr = curr.next {
			if currIndex == index {
				newEntry := &Node[T]{data: elem, next: curr}
				newEntry.prev = prev
				if curr == list.head {
					list.head = newEntry
				}
				if prev != nil {
					prev.next = newEntry
				}
				break
			}
			prev = curr
			currIndex++
		}
	}
	list.size++
}

func (list *LinkedList[T]) Remove(index int) {
	if index >= list.size || index < 0 {
		panic("Out of bounds")
	}
	currIndex := 0
	for curr := list.head; curr != nil; curr = curr.next {
		if currIndex == index-1 {
			if curr.next != nil {
				curr.next = curr.next.next
				if curr.next != nil {
					curr.next.prev = curr
				}
				list.size--
			}
			break
		}
		currIndex++
	}
}

func (list *LinkedList[T]) RemoveSpecific(spec *Node[T]) {
	prev := spec.prev
	if prev != nil {
		prev.next = spec.next
		if spec.next != nil {
			spec.next.prev = prev
		} else {
			list.tail = prev
			list.tail.next = nil
		}
	} else {
		if spec.next != nil {
			list.head = spec.next
			list.head.prev = nil
		} else {
			list.head = nil
			list.tail = nil
		}
	}
	list.size--
}

func (list *LinkedList[T]) Get(index int) T {
	if index >= list.size || index < 0 {
		panic("Out of bounds")
	}
	var zeroValue T
	currIndex := 0
	for curr := list.head; curr != nil; curr = curr.next {
		if currIndex == index {
			return curr.data
		}
		currIndex++
	}
	return zeroValue
}

func (list *LinkedList[T]) GetSize() int {
	return list.size
}

func (list *LinkedList[T]) GetTail() T {
	var zeroVal T
	if list.tail == nil {
		return zeroVal
	}
	return list.tail.data
}

func (list *LinkedList[T]) RemoveTail() T {
	var endVal T
	if list.tail != nil {
		endVal = list.tail.data
		if list.tail.prev != nil {
			list.tail.prev.next = nil
			list.tail = list.tail.prev
		} else {
			list.tail = nil
			list.head = nil
		}
		list.size--
	}
	return endVal
}

func (list *LinkedList[T]) RemoveHead() T {
	var firstVal T
	if list.head != nil {
		firstVal = list.head.data
		if list.head.next != nil {
			list.head.next.prev = nil
			list.head = list.head.next
		} else {
			list.head = nil
			list.tail = nil
		}
		list.size--
	}
	return firstVal
}

func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (list *LinkedList[T]) GetIter() *Iterator.Iterator[T] {
	list.Start()
	return list.iter
}

func (list *LinkedList[T]) Start() {
	if list.iter == nil {
		list.iter = &Iterator.Iterator[T]{}
	}
	list.iter.SetReference(list)
	list.iter.SetCurrent(list.head)
}

func (list *LinkedList[T]) Next() {
	if list.iter.GetCurrent() != nil {
		list.iter.SetCurrent(list.iter.GetCurrent().(*Node[T]).next)
	}
}

func (list *LinkedList[T]) GetIterVal() T {
	var zeroValue T
	if list.iter.GetCurrent() != nil {
		return list.iter.GetCurrent().(*Node[T]).data
	}
	return zeroValue
}

func (list *LinkedList[T]) Done() bool {
	return list.iter.GetCurrent().(*Node[T]) == nil
}
