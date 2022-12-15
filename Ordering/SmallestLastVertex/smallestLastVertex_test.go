package SmallestLastVertex

import (
	"Program2/Graph"
	"testing"
)

func TestRunSmallestLastVertex(t *testing.T) {
	list := Graph.New(10, false)
	list.InitializeComplete()
	list.AddEdge(10, 11, nil)
	list.AddEdge(11, 12, nil)
	list.AddEdge(12, 13, nil)
	list.AddEdge(13, 11, nil)
	Run(list)
}
