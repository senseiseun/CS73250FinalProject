package Analytics

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CompareOrdering() {
	dir := "/Users/seunsuberu/Desktop/AlgorithmEngineering/Program2/Output/Timing"
	items, _ := ioutil.ReadDir(dir)
	m := map[string]map[int]*strings.Builder{}
	for _, item := range items {
		if !item.IsDir() {
			if strings.Contains(item.Name(), "random_degree") {
				splitBySlash := strings.Split(item.Name(), "/")
				fileName := splitBySlash[len(splitBySlash)-1]
				splitByUnder := strings.Split(fileName, "_")
				vertexNumStr := splitByUnder[3]
				vertexNum, convErr := strconv.Atoi(vertexNumStr)
				if convErr != nil {
					fmt.Println("convert error")
					return
				}
				file, openErr := os.Open(dir + "/" + item.Name())
				if openErr != nil {
					fmt.Println("file open error")
				}
				fileScanner := bufio.NewScanner(file)
				fmt.Println(splitByUnder[0], " ", vertexNum)
				if m[splitByUnder[0]] == nil {
					m[splitByUnder[0]] = map[int]*strings.Builder{}
				}
				if m[splitByUnder[0]][vertexNum] == nil {
					m[splitByUnder[0]][vertexNum] = &strings.Builder{}
				}
				fileScanner.Split(bufio.ScanLines)
				count := 0
				sum := 0
				for fileScanner.Scan() {
					line := fileScanner.Text()
					lineSplit := strings.Split(line, ",")
					orderTime, convErr := strconv.Atoi(lineSplit[2])
					sum += orderTime
					count++
					if convErr != nil {
						fmt.Println("conversion error")
					}
				}
				m[splitByUnder[0]][vertexNum].WriteString(strconv.Itoa(int(float64(sum) / float64(count))))
			}
		}
	}

	for key := range m {
		fmt.Println(key)
		keys := make([]int, 0, len(m[key]))

		for k := range m[key] {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		for _, k := range keys {
			fmt.Println(m[key][k].String())

		}
	}
}
