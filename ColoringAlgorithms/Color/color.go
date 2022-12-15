package Color

import (
	"Program2/Graph"
	"Program2/Iterator"
	"Program2/LinkedList"
	"Program2/Util"
	"time"
)

type COLOR int

type Stats struct {
	Color             COLOR
	DegreeWhenDeleted int
	placeInOrder      int
}

type Payload struct {
	OrderingTime int64
	ColorTime    int64
	Stats        *[]*Stats
}

func Color(ordering func(*Graph.AdjacencyList) *LinkedList.LinkedList[*Util.Pair[int, int]], list *Graph.AdjacencyList) *Payload {
	orderingT := time.Unix(0, time.Now().UnixNano())
	orderedList := ordering(list)
	orderingElapsed := time.Since(orderingT)
	var iter *Iterator.Iterator[*Util.Pair[int, int]]
	if orderedList != nil {
		iter = orderedList.GetIter()
	}

	coloringT := time.Unix(0, time.Now().UnixNano())

	colors := make([]COLOR, list.GetSize()+1)

	stats := make([]*Stats, list.GetSize()+1)

	count := 0
	iter.Start()
	for !iter.Done() {
		curr := iter.Get()
		iter2 := list.GetEdgeIterator(curr.First)
		iter2.Start()
		toColor := LinkedList.New[int]()

		takenColors := map[COLOR]bool{}
		stats[curr.First] = &Stats{DegreeWhenDeleted: curr.Second, placeInOrder: count}
		if colors[curr.First] == 0 {
			toColor.Push(curr.First)
		} else {
			stats[curr.First].Color = colors[curr.First]
			takenColors[colors[curr.First]] = true
		}
		for !iter2.Done() {
			edge := iter2.Get()
			if colors[edge] == 0 {
				toColor.Push(edge)
			} else {
				takenColors[colors[edge]] = true
			}
			iter2.Next()
		}

		for color := COLOR(1); color <= COLOR(list.GetSize()); color++ {
			if toColor.GetSize() == 0 {
				break
			}
			if _, has := takenColors[color]; !has {
				first := toColor.RemoveHead()
				colors[first] = color
				if first == curr.First {
					stats[curr.First].Color = color
				}
				takenColors[color] = true
			}
		}
		count++
		iter.Next()
	}
	coloringElapsed := time.Since(coloringT)
	return &Payload{
		OrderingTime: orderingElapsed.Nanoseconds(),
		ColorTime:    coloringElapsed.Nanoseconds(),
		Stats:        &stats,
	}
}
