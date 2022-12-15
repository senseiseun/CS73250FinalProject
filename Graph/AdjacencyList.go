package Graph

import (
	"Program2/Iterator"
	"Program2/LinkedList"
	"fmt"
	"strconv"
	"strings"
)

type AdjacencyList struct {
	list              []*LinkedList.LinkedList[int]
	size              int
	vertexAssociative []any
	edgeAssociative   []map[int]*EdgeBaggage
	isDirected        bool
}

type EdgeBaggage struct {
	listNode *LinkedList.Node[int]
	extra    any
}

func (bag *EdgeBaggage) ListNode() *LinkedList.Node[int] {
	return bag.listNode
}

func (bag *EdgeBaggage) Extra() any {
	return bag.extra
}

func (bag *EdgeBaggage) SetListNode(listNode *LinkedList.Node[int]) {
	bag.listNode = listNode
}

func (bag *EdgeBaggage) SetExtra(extra any) {
	bag.extra = extra
}

func New(size int, isDirected bool) *AdjacencyList {
	list := &AdjacencyList{}
	list.list = make([]*LinkedList.LinkedList[int], size)
	list.vertexAssociative = make([]any, size)
	list.edgeAssociative = make([]map[int]*EdgeBaggage, size)
	list.size = size
	list.isDirected = isDirected
	return list
}

func (list *AdjacencyList) InitializeComplete() {
	for i := 1; i <= list.size; i++ {
		if list.list[i-1] == nil {
			list.list[i-1] = LinkedList.New[int]()
		}
		for j := 1; j <= list.size; j++ {
			if j == i {
				continue
			}
			list.list[i-1].Push(j)
		}
	}
}

func (list *AdjacencyList) InitializeCycle() {
	for i := 1; i <= list.size; i++ {
		if list.list[i-1] == nil {
			list.list[i-1] = LinkedList.New[int]()
		}
		if i == list.size {
			list.list[i-1].Push(1)
		} else {
			list.list[i-1].Push(i + 1)
		}
	}
}

func (list *AdjacencyList) AddEdgeUtil(source, target int, baggage any) {
	if source >= list.size {
		if source > list.size {
			newList := make([]*LinkedList.LinkedList[int], source+10)
			newEdgeAssoc := make([]map[int]*EdgeBaggage, source+10)
			priorLen := len(list.list)
			for i, edges := range list.list {
				if i < priorLen {
					newList[i] = edges
					newEdgeAssoc[i] = list.edgeAssociative[i]
				}
			}
			list.size = source + 10
			list.list = newList
			list.edgeAssociative = newEdgeAssoc
		}
	}
	if list.list[source-1] == nil {
		list.list[source-1] = LinkedList.New[int]()
	}
	node := list.list[source-1].Push(target)
	edgeBags := &EdgeBaggage{listNode: node, extra: baggage}
	if list.edgeAssociative[source-1] == nil {
		list.edgeAssociative[source-1] = map[int]*EdgeBaggage{}
	}
	list.edgeAssociative[source-1][target] = edgeBags
}

func (list *AdjacencyList) AddEdge(source, target int, baggage any) {
	list.AddEdgeUtil(source, target, baggage)
	if !list.isDirected {
		list.AddEdgeUtil(target, source, baggage)
	}
}

func (list *AdjacencyList) RemoveEdge(source, target int) {
	if source >= list.size {
		return
	}
	if list.edgeAssociative[source-1][target] != nil {
		list.list[source-1].RemoveSpecific(list.edgeAssociative[source][target].listNode)
		list.edgeAssociative[source-1][target] = nil
	}
}

func (list *AdjacencyList) GetEdgeBaggage(source, target int) *EdgeBaggage {
	return list.edgeAssociative[source-1][target]
}

func (list *AdjacencyList) GetSize() int {
	return list.size
}

func (list *AdjacencyList) GetEdgeIterator(vertex int) *Iterator.Iterator[int] {
	if list.list[vertex-1] != nil {
		return list.list[vertex-1].GetIter()
	}
	return nil
}

func (list *AdjacencyList) GetDegree(vertex int) int {
	if list.list[vertex-1] != nil {
		return list.list[vertex-1].GetSize()
	}
	return 0
}

func (list *AdjacencyList) SetEdgeBaggage(source, target int, baggage *EdgeBaggage) {
	list.edgeAssociative[source-1][target] = baggage
}

func (list *AdjacencyList) ToString() string {
	bldr := strings.Builder{}
	bldr.WriteString(strconv.Itoa(list.size) + "\n")
	count := list.size + 1
	for i := 1; i <= list.size; i++ {
		bldr.WriteString(strconv.Itoa(count) + "\n")
		count += list.list[i-1].GetSize()
	}
	for i := 1; i <= list.size; i++ {
		iter := list.list[i-1].GetIter()
		iter.Start()
		for !iter.Done() {
			curr := iter.Get()
			bldr.WriteString(strconv.Itoa(curr) + "\n")
			iter.Next()
		}
	}
	return bldr.String()
}

func (list *AdjacencyList) PrintAdj() {
	for i := 1; i <= list.GetSize(); i++ {
		fmt.Print(i, ": ")
		if list.list[i-1] == nil {
			continue
		}
		iter := list.list[i-1].GetIter()
		iter.Start()
		for !iter.Done() {
			curr := iter.Get()
			fmt.Print(curr, ",")
			iter.Next()
		}
		fmt.Println("")
	}
}
