package GraphCreation

import (
	"Program2/Graph"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func CreateTieredGraph(list *Graph.AdjacencyList, maxEdges int) {
	length := list.GetSize()
	set := map[int]*map[int]bool{} // keeping track of the edges that have already been added
	count := 0
	for count < maxEdges {
		for i := 1; i <= length; i++ {
			target := ProduceDisRandom(length, .1)
			// if the randomly generated vertex is the same as the current vertex, then skip
			if i == target {
				continue
			} else {
				if set[i] == nil {
					set[i] = &map[int]bool{}
				} else {
					if (*set[i])[target] {
						continue
					}
				}
				if set[target] == nil {
					set[target] = &map[int]bool{}
				}
				list.AddEdge(i, target, nil)
				(*set[i])[target] = true
				(*set[target])[i] = true
				count++
			}
		}
	}
}

func ProduceDisRandom(length int, factor float64) int {
	// seed new random
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	i := int(r1.Float64() * float64(length))
	// split odd to be 50/50 and based on which side, either first [factor]% or second [1-factor]%
	if i < length/2 {
		return int(r1.Float64()*(float64(length)*factor)) + 1
	} else {
		min := int(float64(length)*factor) + 1
		max := r1.Float64() * float64(length) * (1 - factor)
		return int(max) + min
	}
}

func CreateUniformGraph(list *Graph.AdjacencyList, maxEdges int) {
	length := list.GetSize()
	set := map[int]*map[int]bool{}
	count := 0
	for count < maxEdges {
		for i := 1; i <= length; i++ {
			target := ProduceUniformRandom(length)
			// if the randomly generated vertex is the same as the current vertex, then skip
			if i == target {
				continue
			} else {
				if set[i] == nil {
					set[i] = &map[int]bool{}
				} else {
					if (*set[i])[target] {
						continue
					}
				}
				if set[target] == nil {
					set[target] = &map[int]bool{}
				}
				list.AddEdge(i, target, nil)
				(*set[i])[target] = true
				(*set[target])[i] = true
				count++
			}
		}
	}
}

func ProduceUniformRandom(length int) int {
	// seed new random
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Intn(length) + 1
}

func CreateDoubleTieredGraph(list *Graph.AdjacencyList, maxEdges int) {
	length := list.GetSize()
	set := map[int]*map[int]bool{} // keeping track of the edges that have already been added
	count := 0
	for count < maxEdges {
		for i := 1; i <= length; i++ {
			target := ProduceDisRandomDouble(length, .1)
			// if the randomly generated vertex is the same as the current vertex, then skip
			if i == target {
				continue
			} else {
				if set[i] == nil {
					set[i] = &map[int]bool{}
				} else {
					if (*set[i])[target] {
						continue
					}
				}
				if set[target] == nil {
					set[target] = &map[int]bool{}
				}
				list.AddEdge(i, target, nil)
				(*set[i])[target] = true
				(*set[target])[i] = true
				count++
			}
		}
	}
}

func ProduceDisRandomDouble(length int, factor float64) int {
	// seed new random
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	i := r1.Float64()
	// split odd to be 25/50/25 and based on which side, either first [factor]%, second [1-factor * 2]%, or the last [factor]%
	if i < .25 {
		return int(r1.Float64()*(float64(length)*factor)) + 1
	} else if i < .75 {
		min := int(float64(length)*factor) + 1
		max := r1.Float64() * float64(length) * (1 - factor*2)
		return int(max) + min
	} else {
		sub := int(r1.Float64() * (float64(length) * factor))
		return length - sub
	}
}

func OutputEdgeCount(list *Graph.AdjacencyList, name string) {
	bldr := strings.Builder{}
	for vertex := 1; vertex <= list.GetSize(); vertex++ {
		bldr.WriteString(strconv.Itoa(vertex) + "," + strconv.Itoa(list.GetDegree(vertex)) + "\n")
	}
	newCSV, createErr := os.Create("./Output/Histogram/" + name + "_edgeCount.csv")
	if createErr != nil {
		fmt.Errorf("file output error")
		return
	}
	_, writeErr := newCSV.WriteString(bldr.String())
	if writeErr != nil {
		fmt.Println("edge count write to file errored")
	}
}
