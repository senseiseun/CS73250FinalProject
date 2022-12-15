package SmallestLastVertex

import (
	"Program2/Graph"
	"Program2/LinkedList"
	"Program2/Util"
	"math"
)

func Run(list *Graph.AdjacencyList) *LinkedList.LinkedList[*Util.Pair[int, int]] {
	// array of linked list pointers, linked list stores vertices, the index symbolizes the degree count of the vertex
	/*
		1: 5->3->->4 (vertices 5, 3 and 4 have degrees of 1)
		2: 1-> 7-> 6 (vertices 1, 7, and 6 have degrees of 2)
	*/
	degrees := make([]*LinkedList.LinkedList[int], list.GetSize())
	// array of Node pointers so we can directly access the node within a linked list, the index of the array symbolizes the vertex
	listNodes := make([]*LinkedList.Node[int], list.GetSize())
	// array of degrees, the index of the array maps to the vertex number, the value at the index is the degree
	actualDegrees := make([]int, list.GetSize())
	// the resulting queue for the ordering, linked list of pairs, the pairs are [first: vertex number, second: degree when deleted]
	finalQueue := LinkedList.LinkedList[*Util.Pair[int, int]]{}
	// to see if vertex has been "taken"
	taken := make([]bool, list.GetSize())
	// initialize degrees with each vertex's original degree
	for i := 1; i <= list.GetSize(); i++ {
		if degrees[list.GetDegree(i)] == nil {
			degrees[list.GetDegree(i)] = LinkedList.New[int]()
		}
		// get node pointer from insertion in linked list
		node := degrees[list.GetDegree(i)].Push(i)
		// set original degree
		actualDegrees[i-1] = list.GetDegree(i)
		// set current pointer
		listNodes[i-1] = node
	}
	// iterate through all degrees
	for i := 0; i < list.GetSize(); i++ {
		if i < 0 {
			continue
		}
		// if the degree list is not nil and the linked list is not empty
		if degrees[i] != nil && degrees[i].GetSize() > 0 {
			// while there are elements in the linked list
			for degrees[i].GetSize() > 0 {
				// get the tail
				end := degrees[i].RemoveTail()
				// this element has been taken
				taken[end-1] = true
				// get the degree before it is removed
				degreeBeforeRemove := actualDegrees[end-1]
				minDegrees := degreeBeforeRemove
				// iterate over the deleted vertex's edges
				iter := list.GetEdgeIterator(end)
				if iter == nil {
					continue
				}
				iter.Start()
				for !iter.Done() {
					// the chosen adjacent vertex
					curr := iter.Get()
					// if the edge has been removed then skip
					if taken[curr-1] {
						iter.Next()
						continue
					}
					// the previous degree before decrementing
					oldDegree := actualDegrees[curr-1]
					// the new degree after decrementing
					newDegree := oldDegree - 1
					// if degrees does not have a linked list at the new degree then create one
					if degrees[newDegree] == nil {
						degrees[newDegree] = LinkedList.New[int]()
					}
					// remove the stored node pointer in the linked list for where curr used to be
					degrees[oldDegree].RemoveSpecific(listNodes[curr-1])
					// add curr to the new degree linked list and update the node pointer in the array
					listNodes[curr-1] = degrees[newDegree].Push(curr)
					// update its degrees
					actualDegrees[curr-1]--
					// if the new degree of the adjacent vertex is less than the current i, then set i to it so we can make sure the algorithm addresses it
					minDegrees = int(math.Min(float64(minDegrees), float64(actualDegrees[curr-1])))
					iter.Next()
				}
				// add the deleted vertex to the final queue
				finalQueue.PushFront(&Util.Pair[int, int]{First: end, Second: degreeBeforeRemove})
				// address the changed vertices
				i = minDegrees - 1
				break
			}
		}
	}
	return &finalQueue
}
