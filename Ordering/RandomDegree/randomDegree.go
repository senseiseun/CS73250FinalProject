package RandomDegree

import (
	"Program2/Graph"
	"Program2/LinkedList"
	"Program2/Util"
	"math/rand"
	"reflect"
	"time"
)

func Run(list *Graph.AdjacencyList) *LinkedList.LinkedList[*Util.Pair[int, int]] {
	// list of vertices
	randomized := make([]int, list.GetSize())
	// the resulting queue for the ordering, linked list of pairs, the pairs are [first: vertex number, second: degree when deleted]
	finalQueue := LinkedList.New[*Util.Pair[int, int]]()
	// iterates through all the vertices in the graph
	for vertex := 1; vertex <= list.GetSize(); vertex++ {
		randomized[vertex-1] = vertex
	}
	// seed new randomizer
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	// swap vertices 2n times, n being the number of vertices in graph
	for i := 0; i < 2*list.GetSize(); i++ {
		// get 2 random vertices and swap
		vertex1 := r1.Intn(list.GetSize())
		vertex2 := r1.Intn(list.GetSize())
		swapF := reflect.Swapper(randomized)
		swapF(vertex2, vertex1)
	}
	// fill the final queue with the randomized array values
	for i := 1; i <= list.GetSize(); i++ {
		finalQueue.Push(&Util.Pair[int, int]{First: i, Second: list.GetDegree(i)})
	}
	return finalQueue
}
