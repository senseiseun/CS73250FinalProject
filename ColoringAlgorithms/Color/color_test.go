package Color

import (
	"Program2/Graph"
	"Program2/GraphCreation"
	"Program2/Ordering/SmallestLastVertex"
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {
	list := Graph.New(100, false)
	GraphCreation.CreateTieredGraph(list, 1000)
	res := Color(SmallestLastVertex.Run, list)
	for vertex, color := range *res.Stats {
		if vertex == 0 {
			continue
		}
		fmt.Println(vertex, color)
	}
}
