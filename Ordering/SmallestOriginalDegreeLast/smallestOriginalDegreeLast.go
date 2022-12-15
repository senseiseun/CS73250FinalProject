package SmallestOriginalDegreeLast

import (
	"Program2/Graph"
	"Program2/LinkedList"
	"Program2/Util"
)

func Run(list *Graph.AdjacencyList) *LinkedList.LinkedList[*Util.Pair[int, int]] {
	// array of linked list pointers, linked list stores vertices, the index symbolizes the degree count of the vertex
	/*
		1: 5->3->->4 (vertices 5, 3 and 4 have degrees of 1)
		2: 1-> 7-> 6 (vertices 1, 7, and 6 have degrees of 2)
	*/
	degrees := make([]*LinkedList.LinkedList[int], list.GetSize())
	// the resulting queue for the ordering, linked list of pairs, the pairs are [first: vertex number, second: degree when deleted]
	finalQueue := LinkedList.LinkedList[*Util.Pair[int, int]]{}
	// initialize degrees with each vertex's original degree
	for i := 1; i <= list.GetSize(); i++ {
		if degrees[list.GetDegree(i)] == nil {
			degrees[list.GetDegree(i)] = LinkedList.New[int]()
		}
		degrees[list.GetDegree(i)].Push(i)
	}
	// iterate through all degrees
	for i := 0; i < list.GetSize(); i++ {
		// if the degree list is not nil
		if degrees[i] != nil {
			iter := degrees[i].GetIter()
			iter.Start()
			for !iter.Done() {
				curr := iter.Get()
				finalQueue.PushFront(&Util.Pair[int, int]{First: curr, Second: i})
				iter.Next()
			}
		}
	}
	return &finalQueue
}
