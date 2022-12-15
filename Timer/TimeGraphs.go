package Timer

import (
	"Program2/Graph"
	"os"
	"strconv"
	"strings"
	"time"
)

func TimeCompleteGraphs() *strings.Builder {
	buf := strings.Builder{}
	for i := 0; i < 1000; i += 10 {
		buf.WriteString(strconv.Itoa(i) + "," + strconv.Itoa(int(TimeCompleteGraph(i))) + "\n")
	}
	return &buf
}

func TimeCycleGraphs() *strings.Builder {
	buf := strings.Builder{}
	for i := 0; i < 1000; i += 10 {
		buf.WriteString(strconv.Itoa(i) + "," + strconv.Itoa(int(TimeCycleGraph(i))) + "\n")
	}
	return &buf
}

func OutputTimesToFile(buf *strings.Builder) {
	csvFile, osCreateErr := os.Create("./timedGraphs.csv")
	if osCreateErr != nil {
		return
	}
	csvFile.WriteString(buf.String())
}

func TimeCompleteGraph(vertexCount int) int64 {
	list := Graph.New(vertexCount, false)
	t := time.Unix(0, time.Now().UnixNano())
	list.InitializeComplete()
	elapsed := time.Since(t)
	return elapsed.Nanoseconds()
}

func TimeCycleGraph(vertexCount int) int64 {
	list := Graph.New(vertexCount, false)
	t := time.Unix(0, time.Now().UnixNano())
	list.InitializeCycle()
	elapsed := time.Since(t)
	return elapsed.Nanoseconds()
}
