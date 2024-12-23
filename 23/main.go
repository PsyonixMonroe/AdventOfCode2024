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

func FindLanPassword(connections map[string][]string) string {
	work := []string{}
	finalSet := ""

	// seed work queue
	for k, v := range connections {
		for _, child := range v {
			set := []string{k, child}
			sort.Strings(set)
			work = lib.AppendDistinct(work, fmt.Sprintf("%s,%s", set[0], set[1]))
		}
	}

	var idx int
	addSets := func() (string, bool) {
		if idx == len(work) {
			return "", true
		}
		item := work[idx]
		idx += 1
		return item, false
	}

	processItem := func(set string) []string {
		newSets := []string{}

		// parse set into []string for all items in set
		currentSet := strings.Split(set, ",")

		// get all children of items in set that are not included in set
		candidates := []string{}
		for _, included := range currentSet {
			children := connections[included]
			for _, child := range children {
				if !IsInArray(child, currentSet) && !IsInArray(child, candidates) {
					candidates = append(candidates, child)
				}
			}
		}

		// check if child could be included
		for _, candidate := range candidates {
			candConns := connections[candidate]
			include := true
			for _, s := range currentSet {
				if !IsInArray(s, candConns) {
					// candidate doesn't have s as a connection, drop candidate
					include = false
					break
				}
			}
			if include {
				newSet := append(currentSet, candidate)
				sort.Strings(newSet)
				newItem := strings.Join(newSet, ",")
				newSets = lib.AppendDistinct(newSets, newItem)
			}
		}

		return newSets
	}

	newWork := []string{}
	collectNewWork := func(input []string) {
		for _, i := range input {
			newWork = lib.AppendDistinct(newWork, i)
		}
	}

	for len(work) > 0 {
		idx = 0
		wp := lib.WorkerPool[string, []string]{NumWorkers: 12, Verbose: false, GetWorkItem: addSets, WorkerFn: processItem, ReducerFn: collectNewWork}
		wp.Run(context.Background())
		if len(work) == 1 {
			finalSet = work[0]
			break
		}
		work = newWork
		newWork = []string{}
	}

	return finalSet
}

func IsInArray(cand string, items []string) bool {
	idx, _ := lib.FindInArray(cand, items)
	return idx != -1
}
