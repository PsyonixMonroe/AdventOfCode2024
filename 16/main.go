package main

import (
	"github.com/PsyonixMonroe/AOCLib/lib"
)

var (
	// these are vars but don't change them
	North = lib.Location{X: -1, Y: 0}
	East  = lib.Location{X: 0, Y: 1}
	South = lib.Location{X: 1, Y: 0}
	West  = lib.Location{X: 0, Y: -1}
)

type Node struct {
	gVal   int
	hVal   int
	loc    lib.Location
	facing lib.Location
}

func (n Node) GetF() int {
	return n.gVal + n.hVal
}

func (n Node) GetID(columns int) int {
	return n.loc.X*columns + n.loc.Y
}

func GetHVal(current lib.Location, end lib.Location) int {
	val := lib.Abs(end.Y-current.Y) + lib.Abs(end.X-current.X)
	if end.Y-current.Y == 0 || end.X-current.X == 0 {
		return val
	}
	return 1000 + val
}

type Course struct {
	maze  lib.Grid[rune]
	start lib.Location
	end   lib.Location
}

func ParseCourse(content string) Course {
	maze := lib.ParseRuneGrid(content)
	start, end := FindStartAndEnd(maze)
	return Course{maze, start, end}
}

func FindStartAndEnd(maze lib.Grid[rune]) (lib.Location, lib.Location) {
	var start lib.Location
	var end lib.Location
	foundStart := false
	foundEnd := false
	for i := range maze.Rows() {
		for j := range maze.Columns() {
			item := *maze.Get(i, j).Get()
			if item == 'E' {
				end = lib.Location{X: i, Y: j}
				foundEnd = true
			}
			if item == 'S' {
				start = lib.Location{X: i, Y: j}
				foundStart = true
			}
			if foundStart && foundEnd {
				break
			}
		}
		if foundStart && foundEnd {
			break
		}
	}

	return start, end
}

func FindBestPath(course Course) int {
	startHVal := GetHVal(course.start, course.end)
	startNode := Node{gVal: 0, hVal: startHVal, loc: course.start, facing: East}
	searchNodes := []Node{startNode}
	finishedNodes := []Node{}
	var endNode Node
	for {
		// Get New Node to process
		current := searchNodes[0]
		for _, node := range searchNodes {
			if node.loc.Equal(current.loc) {
				continue
			}
			if node.GetF() < current.GetF() {
				current = node
			}
			if node.GetF() == current.GetF() {
				if node.hVal < current.hVal {
					current = node
				}
			}
		}
		facing := current.facing
		searchNodes, _ = lib.RemoveItem(searchNodes, current)

		// Check if we are done
		if current.loc.Equal(course.end) {
			endNode = current
			break
		}

		// Add All the children
		northLoc := lib.Location{X: current.loc.X + North.X, Y: current.loc.Y + North.Y}
		southLoc := lib.Location{X: current.loc.X + South.X, Y: current.loc.Y + South.Y}
		eastLoc := lib.Location{X: current.loc.X + East.X, Y: current.loc.Y + East.Y}
		westLoc := lib.Location{X: current.loc.X + West.X, Y: current.loc.Y + West.Y}
		if *course.maze.GetLocation(northLoc).Get() != '#' && !IsLocInNodes(finishedNodes, northLoc) {
			// need to process north
			// check if in search
			idx := FindLocInArray(searchNodes, northLoc)
			var northNode *Node
			if idx == -1 {
				// not in search
				northNode = &Node{}
				northNode.loc = northLoc
				northNode.hVal = GetHVal(northLoc, course.end)

				idx = len(searchNodes)
				searchNodes = append(searchNodes, *northNode)
			}
			northNode = &searchNodes[idx]
			cost := GetNodeCost(North, facing)
			northNode.gVal = current.gVal + cost
			northNode.facing = North
		}

		// south
		if *course.maze.GetLocation(southLoc).Get() != '#' && !IsLocInNodes(finishedNodes, southLoc) {
			// need to process south
			// check if in search
			idx := FindLocInArray(searchNodes, southLoc)
			var southNode *Node
			if idx == -1 {
				// not in search
				southNode = &Node{}
				southNode.loc = southLoc
				southNode.hVal = GetHVal(southLoc, course.end)

				idx = len(searchNodes)
				searchNodes = append(searchNodes, *southNode)
			}
			southNode = &searchNodes[idx]
			cost := GetNodeCost(South, facing)
			southNode.gVal = current.gVal + cost
			southNode.facing = South
		}

		// east
		if *course.maze.GetLocation(eastLoc).Get() != '#' && !IsLocInNodes(finishedNodes, eastLoc) {
			// need to process east
			// check if in search
			idx := FindLocInArray(searchNodes, eastLoc)
			var eastNode *Node
			if idx == -1 {
				// not in search
				eastNode = &Node{}
				eastNode.loc = eastLoc
				eastNode.hVal = GetHVal(eastLoc, course.end)

				idx = len(searchNodes)
				searchNodes = append(searchNodes, *eastNode)
			}
			eastNode = &searchNodes[idx]
			cost := GetNodeCost(East, facing)
			eastNode.gVal = current.gVal + cost
			eastNode.facing = East
		}

		// west
		if *course.maze.GetLocation(westLoc).Get() != '#' && !IsLocInNodes(finishedNodes, westLoc) {
			// need to process west
			// check if in search
			idx := FindLocInArray(searchNodes, westLoc)
			var westNode *Node
			if idx == -1 {
				// not in search
				westNode = &Node{}
				westNode.loc = westLoc
				westNode.hVal = GetHVal(westLoc, course.end)

				idx = len(searchNodes)
				searchNodes = append(searchNodes, *westNode)
			}
			westNode = &searchNodes[idx]
			cost := GetNodeCost(West, facing)
			westNode.gVal = current.gVal + cost
			westNode.facing = West
		}

		finishedNodes = append(finishedNodes, current)
	}

	return ScoreNodePath(endNode)
}

func GetNodeCost(dir lib.Location, facing lib.Location) int {
	newGVal := 1
	if !dir.Equal(facing) {
		newGVal += 1000
	}

	return newGVal
}

func FindLocInArray(nodes []Node, loc lib.Location) int {
	for i, node := range nodes {
		if loc.Equal(node.loc) {
			return i
		}
	}

	return -1
}

func IsLocInNodes(nodes []Node, loc lib.Location) bool {
	for _, node := range nodes {
		if loc.Equal(node.loc) {
			return true
		}
	}

	return false
}

func ScoreNodePath(end Node) int {
	return end.GetF()
}

func FindBestFScore(nodes []Node) Node {
	bestF := nodes[0]
	for _, node := range nodes {
		if node.GetF() < bestF.GetF() {
			bestF = node
		}
	}

	return bestF
}

func GetFacingScore(facing lib.Location, comp lib.Location) int {
	if facing.X == comp.X && facing.Y == comp.Y {
		return 1
	}
	return 1001
}

func IsInPrevious(previous []lib.Location, test lib.Location) bool {
	for _, prev := range previous {
		if test.X == prev.X && test.Y == prev.Y {
			return true
		}
	}
	return false
}
