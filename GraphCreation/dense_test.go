package GraphCreation

import (
	"Program2/Graph"
	"Program2/Ordering/SmallestLastVertex"
	"testing"
)

func TestCreateTieredGraph(t *testing.T) {
	list := Graph.New(10, false)
	CreateTieredGraph(list, 7)
	list.PrintAdj()
	SmallestLastVertex.Run(list)
}

func TestOutputEdgeCount(t *testing.T) {
	list := Graph.New(500, false)
	CreateDoubleTieredGraph(list, 50000)
	OutputEdgeCount(list, "")
}
