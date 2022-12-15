package main

import (
	"Program2/ColoringAlgorithms/Color"
	"Program2/Graph"
	"Program2/GraphCreation"
	"Program2/LinkedList"
	"Program2/Ordering/RandomDegree"
	"Program2/Ordering/SmallestLastVertex"
	"Program2/Ordering/SmallestOriginalDegreeLast"
	"Program2/Util"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
}

func OutputTiming() {
	graphCreationList := []func(*Graph.AdjacencyList, int){GraphCreation.CreateComplete, GraphCreation.CreateCycle, GraphCreation.CreateTieredGraph, GraphCreation.CreateDoubleTieredGraph, GraphCreation.CreateUniformGraph}
	names := []string{"complete_", "cycle_", "tiered_", "doubleTiered_", "uniform_"}
	orderingAlgos := []func(*Graph.AdjacencyList) *LinkedList.LinkedList[*Util.Pair[int, int]]{SmallestLastVertex.Run, SmallestOriginalDegreeLast.Run, RandomDegree.Run}
	orderingAlgoNames := []string{"smallest_last_vertex_", "smallest_original_last_", "random_degree_"}
	for vertices := 1000; vertices <= 10000; vertices += 1000 {
		bldrs := [][]*strings.Builder{
			{&strings.Builder{}, &strings.Builder{}, &strings.Builder{}},
			{&strings.Builder{}, &strings.Builder{}, &strings.Builder{}},
			{&strings.Builder{}, &strings.Builder{}, &strings.Builder{}},
			{&strings.Builder{}, &strings.Builder{}, &strings.Builder{}},
			{&strings.Builder{}, &strings.Builder{}, &strings.Builder{}},
		}
		graphOutput := make([]*strings.Builder, 5)
		fmt.Println("vertex: ", vertices)
		max := (vertices / 2) * (vertices / 10)
		if max > 2000000 {
			max = 2000000
		}
		for edges := vertices * 5; edges <= 2000000 && edges < max; edges += max / 20 {
			fmt.Println("\tedges: ", edges)
			for i, create := range graphCreationList {
				for j, ordering := range orderingAlgos {
					list := Graph.New(vertices, false)
					//fmt.Println("\t", names[i])
					create(list, edges)
					//fmt.Println("\tafter: ", names[i])
					name := names[i] + strconv.Itoa(vertices) + "_" + strconv.Itoa(edges)
					GraphCreation.OutputEdgeCount(list, name)
					if graphOutput[i] == nil {
						graphOutput[i] = &strings.Builder{}
						graphOutput[i].WriteString(list.ToString())
					}
					//fmt.Println("\t\t", orderingAlgoNames[j])
					payload := Color.Color(ordering, list)
					edgeCountStr := strconv.Itoa(edges)
					colorsTimeStr := strconv.Itoa(int(payload.ColorTime))
					orderingTimeStr := strconv.Itoa(int(payload.OrderingTime))
					bldrs[i][j].WriteString(edgeCountStr + "," + colorsTimeStr + "," + orderingTimeStr + "\n")
					//fmt.Println("\t\tafter: ", orderingAlgoNames[j])
				}
			}
		}
		for i := range graphCreationList {
			for j := range orderingAlgos {
				fileName := "Output/Timing/" + names[i] + orderingAlgoNames[j] + strconv.Itoa(vertices) + "_" + "times.csv"
				file, fileErr := os.Create(fileName)
				if fileErr != nil {
					fmt.Println("file creation error for times")
				}
				_, writeErr := file.WriteString(bldrs[i][j].String())
				if writeErr != nil {
					fmt.Println("write to file errored")
				}
			}
			fileName := "Output/Graphs/" + names[i] + strconv.Itoa(vertices) + ".csv"
			file, fileErr := os.Create(fileName)
			if fileErr != nil {
				fmt.Println("file creation error for times")
			}
			_, writeErr := file.WriteString(graphOutput[i].String())
			if writeErr != nil {
				fmt.Println("write to file errored")
			}
		}
	}
}

func OutputGraphEdges() {
	graphCreationList := []func(*Graph.AdjacencyList, int){GraphCreation.CreateComplete, GraphCreation.CreateCycle, GraphCreation.CreateTieredGraph, GraphCreation.CreateDoubleTieredGraph, GraphCreation.CreateUniformGraph}
	names := []string{"complete_", "cycle_", "tiered_", "doubleTiered_", "uniform_"}
	for vertices := 1000; vertices <= 4000; vertices += 1000 {
		fmt.Println("vertex: ", vertices)
		max := (vertices / 2) * (vertices / 10)
		if max > 2000000 {
			max = 2000000
		}
		for edges := vertices * 5; edges <= 2000000 && edges < max; edges += max / 20 {
			fmt.Println("\tedges: ", edges)
			for i, create := range graphCreationList {
				list := Graph.New(vertices, false)
				//fmt.Println("\t", names[i])
				create(list, edges)
				//fmt.Println("\tafter: ", names[i])
				name := names[i] + strconv.Itoa(vertices) + "_" + strconv.Itoa(edges)
				GraphCreation.OutputEdgeCount(list, name)
			}
		}
	}
}

func OutputGraphCreationTiming() {
	graphCreationList := []func(*Graph.AdjacencyList, int){GraphCreation.CreateComplete, GraphCreation.CreateCycle, GraphCreation.CreateTieredGraph, GraphCreation.CreateDoubleTieredGraph, GraphCreation.CreateUniformGraph}
	names := []string{"complete_", "cycle_", "tiered_", "doubleTiered_", "uniform_"}
	for vertices := 1000; vertices <= 10000; vertices += 1000 {
		fmt.Println("vertex: ", vertices)
		max := (vertices / 2) * (vertices / 10)
		if max > 2000000 {
			max = 2000000
		}
		bldrs := []strings.Builder{
			strings.Builder{},
			strings.Builder{},
			strings.Builder{},
			strings.Builder{},
			strings.Builder{},
		}
		for edges := vertices * 5; edges <= 2000000 && edges < max; edges += max / 20 {
			fmt.Println("\tedges: ", edges)

			for i, create := range graphCreationList {
				list := Graph.New(vertices, false)
				//fmt.Println("\t", names[i])
				createTimeStart := time.Unix(0, time.Now().UnixNano())
				create(list, edges)
				createTimeStamp := time.Since(createTimeStart)
				//fmt.Println("\tafter: ", names[i])
				name := names[i] + strconv.Itoa(vertices) + "_" + strconv.Itoa(edges)
				GraphCreation.OutputEdgeCount(list, name)
				bldrs[i].WriteString(strconv.Itoa(edges) + "," + strconv.Itoa(int(createTimeStamp.Nanoseconds())) + "\n")
			}

		}
		for i := 0; i < 5; i++ {
			fileName := "Output/Timing/GraphCreation/" + names[i] + strconv.Itoa(vertices) + "_" + "times.csv"
			file, fileErr := os.Create(fileName)
			if fileErr != nil {
				fmt.Println("file creation error for times")
			}
			_, writeErr := file.WriteString(bldrs[i].String())
			if writeErr != nil {
				fmt.Println("write to file errored")
			}
		}
	}
}

func DegreeDeleted() {
	graphCreationList := []func(*Graph.AdjacencyList, int){GraphCreation.CreateComplete, GraphCreation.CreateCycle, GraphCreation.CreateTieredGraph, GraphCreation.CreateDoubleTieredGraph, GraphCreation.CreateUniformGraph}
	names := []string{"complete_", "cycle_", "tiered_", "doubleTiered_", "uniform_"}
	orderingAlgos := []func(*Graph.AdjacencyList) *LinkedList.LinkedList[*Util.Pair[int, int]]{SmallestLastVertex.Run, SmallestOriginalDegreeLast.Run, RandomDegree.Run}
	orderingAlgoNames := []string{"smallest_last_vertex_", "smallest_original_last_", "random_degree_"}
	for vertices := 1000; vertices <= 10000; vertices += 1000 {
		bldrs := make([][]*strings.Builder, 20)
		for i := 0; i < 20; i++ {
			bldrs[i] = make([]*strings.Builder, 5)
			for j := 0; j < 5; j++ {
				bldrs[i][j] = &strings.Builder{}
			}
		}
		graphOutput := make([]*strings.Builder, 5)
		fmt.Println("vertex: ", vertices)
		max := (vertices / 2) * (vertices / 10)
		if max > 2000000 {
			max = 2000000
		}
		x := -1
		for edges := vertices * 5; edges <= 2000000 && edges < max; edges += max / 20 {
			x++
			fmt.Println("\tedges: ", edges)
			for i := range graphCreationList {
				bldrs[x][i].WriteString(strconv.Itoa(edges) + "\n")
			}
			for i, create := range graphCreationList {
				list := Graph.New(vertices, false)
				create(list, edges)
				name := names[i] + strconv.Itoa(vertices) + "_" + strconv.Itoa(edges)
				GraphCreation.OutputEdgeCount(list, name)
				for j, ordering := range orderingAlgos {
					if graphOutput[i] == nil {
						graphOutput[i] = &strings.Builder{}
						graphOutput[i].WriteString(list.ToString())
					}
					payload := Color.Color(ordering, list)
					var maxDeleted int
					var maxColor int
					for _, stat := range *payload.Stats {
						if stat == nil {
							continue
						}
						maxDeleted = int(math.Max(float64(stat.DegreeWhenDeleted), float64(maxDeleted)))
						maxColor = int(math.Max(float64(stat.Color), float64(maxColor)))
					}
					maxDeletedStr := strconv.Itoa(maxDeleted)
					maxColorStr := strconv.Itoa(maxColor)
					bldrs[x][i].WriteString(orderingAlgoNames[j] + "," + maxDeletedStr + "," + maxColorStr + "\n")
				}
			}
		}
		for i := 0; i < x; i++ {
			for j := range graphCreationList {
				fileName := "Output/DegreeDeleted/" + names[j] + strconv.Itoa(vertices) + "_" + "times.csv"
				file, fileErr := os.Create(fileName)
				if fileErr != nil {
					fmt.Println("file creation error for times")
				}
				_, writeErr := file.WriteString(bldrs[i][j].String())
				if writeErr != nil {
					fmt.Println("write to file errored")
				}
			}
		}
	}
}
