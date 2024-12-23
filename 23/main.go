package main

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/PsyonixMonroe/AOCLib/lib"
)

type Tuple struct {
	a string
	b string
}

func ParseConnections(content string) []Tuple {
	lines := strings.Split(strings.Trim(content, "\n \t"), "\n")
	connections := []Tuple{}
	for _, line := range lines {
		conn := strings.Trim(line, "\n \t")
		splits := strings.Split(conn, "-")
		if len(splits) != 2 {
			fmt.Fprintf(os.Stderr, "Unable to parse connection: %s\n", conn)
			continue
		}
		connections = append(connections, Tuple{a: splits[0], b: splits[1]})
	}

	return connections
}

func IndexFromEdges(connections []Tuple) map[string][]string {
	index := make(map[string][]string)
	for _, conn := range connections {
		AddToMap(&index, conn.a, conn.b)
		AddToMap(&index, conn.b, conn.a)
	}

	return index
}

func AddToMap(connections *map[string][]string, source string, target string) {
	v, found := (*connections)[source]
	if !found {
		v = []string{}
	}
	v = append(v, target)
	(*connections)[source] = v
}

func CountNetworkedComputers(connections map[string][]string) int {
	sum := 0

	generatedWork := []string{}
	findTComputers := func() (string, bool) {
		for k := range connections {
			idx, _ := lib.FindInArray(k, generatedWork)
			if strings.HasPrefix(k, "t") && idx == -1 {
				generatedWork = append(generatedWork, k)
				return k, false
			}
		}
		return "", true
	}

	findNetworked := func(first string) []string {
		baseConnections := connections[first]
		sets := []string{}
		for _, second := range baseConnections {
			for _, third := range connections[second] {
				if third == first {
					continue
				}
				thirdConnections := connections[third]
				firstIdx, _ := lib.FindInArray(first, thirdConnections)
				if firstIdx == -1 {
					continue
				}
				items := []string{first, second, third}
				sort.Strings(items)
				set := fmt.Sprintf("%s,%s,%s", items[0], items[1], items[2])
				sets = append(sets, set)
			}
		}

		return sets
	}

	foundNetworked := []string{}
	countUniqueNetworked := func(input []string) {
		for _, set := range input {
			idx, _ := lib.FindInArray(set, foundNetworked)
			if idx == -1 {
				foundNetworked = append(foundNetworked, set)
				sum++
			}
		}
	}

	wp := lib.WorkerPool[string, []string]{NumWorkers: 12, GetWorkItem: findTComputers, WorkerFn: findNetworked, ReducerFn: countUniqueNetworked, Verbose: false}
	wp.Run(context.Background())

	return sum
}
