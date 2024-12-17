package main

import (
	"fmt"
	"math"
	"os"

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
	prev   *Node
	next   lib.Location
	loc    lib.Location
	facing lib.Location
}

func (n Node) GetF() int {
	return n.gVal + n.hVal
}

func (n Node) GetID(columns int) int {
	return n.loc.X*columns + n.loc.Y
}

func (n Node) String() string {
	return fmt.Sprintf("Node{Prev: %v, Next: %v, loc: %v, facing:%v}", n.prev.loc, n.next, n.loc, n.facing)
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

func FindBestPath(course Course) (int, Node) {
	startHVal := GetHVal(course.start, course.end)
	startNode := Node{gVal: 0, hVal: startHVal, loc: course.start, facing: East}
	searchNodes := []Node{startNode}
	finishedNodes := []Node{}
	var endNode Node
	for {
		if len(searchNodes) == 0 {
			// hit the end
			return -1, Node{}
		}
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
			northNode.prev = &current
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
			southNode.prev = &current
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
			eastNode.prev = &current
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
			westNode.prev = &current
		}

		finishedNodes = append(finishedNodes, current)
	}

	return ScoreNodePath(endNode), endNode
}

func FindAllPaths(course Course, content string, maxVal int, endNode Node, previousBlocks []lib.Location) []lib.Location {
	blockNodes := []Node{}
	successNodes := []lib.Location{course.start, endNode.loc}
	node := *endNode.prev
	node.next = endNode.loc
	for !node.loc.Equal(course.start) {
		node.prev.next = node.loc
		blockNodes = append(blockNodes, node)
		successNodes = lib.AppendDistinct(successNodes, node.loc)
		node = *node.prev
	}

	if len(previousBlocks) == 0 {
		fmt.Fprintf(os.Stderr, "Initial Solution: %v\n", successNodes)
	}

	for len(blockNodes) != 0 {
		blockNode := blockNodes[0]
		blockNodes = blockNodes[1:]

		if numBranches(course, blockNode.next) > 2 {
			// build a new maze
			newMaze := lib.ParseRuneGrid(content)
			newMaze.PutLocation('#', blockNode.loc)
			for _, prevBlock := range previousBlocks {
				newMaze.PutLocation('#', prevBlock)
			}
			newCourse := Course{maze: newMaze, start: course.start, end: course.end}
			score, newEnd := FindBestPath(newCourse)
			if score == maxVal {
				// found another success
				node = *newEnd.prev
				newSuccess := []lib.Location{}
				for !node.loc.Equal(course.start) {
					successNodes = lib.AppendDistinct(successNodes, node.loc)
					newSuccess = append(newSuccess, node.loc)
					node = *node.prev
				}
				fmt.Fprintf(os.Stderr, "Found New Solution [%d]: %v\n", len(previousBlocks), newSuccess)
				fmt.Fprintf(os.Stderr, "Blocks: %v\n", previousBlocks)
				newBlocks := append(previousBlocks, blockNode.loc)

				recursiveSuccess := FindAllPaths(course, content, maxVal, newEnd, newBlocks)
				for _, success := range recursiveSuccess {
					successNodes = lib.AppendDistinct(successNodes, success)
				}
			}
		}
	}

	return successNodes
}

func numBranches(course Course, loc lib.Location) int {
	numBranches := 0
	northLoc := lib.Location{X: loc.X + North.X, Y: loc.Y + North.Y}
	southLoc := lib.Location{X: loc.X + South.X, Y: loc.Y + South.Y}
	eastLoc := lib.Location{X: loc.X + East.X, Y: loc.Y + East.Y}
	westLoc := lib.Location{X: loc.X + West.X, Y: loc.Y + West.Y}
	if *course.maze.GetLocation(northLoc).Get() != '#' {
		numBranches++
	}
	if *course.maze.GetLocation(southLoc).Get() != '#' {
		numBranches++
	}
	if *course.maze.GetLocation(eastLoc).Get() != '#' {
		numBranches++
	}
	if *course.maze.GetLocation(westLoc).Get() != '#' {
		numBranches++
	}

	return numBranches
}

type BFSNode struct {
	path    []lib.Location
	current lib.Location
	face    lib.Location
	score   int
}

func FindAllPathsBFS(course Course, maxVal int) []lib.Location {
	processingQueue := lib.NewQueue[BFSNode]()
	processingQueue.Enqueue(BFSNode{path: []lib.Location{}, current: course.start, face: East, score: 0})
	successLocations := []lib.Location{course.end}
	for !processingQueue.IsEmpty() {
		node := processingQueue.Dequeue().Get()
		if node.current.Equal(course.end) {
			fmt.Fprintf(os.Stderr, "Found path [%d  - %v]\n", node.score, node.path)
			for _, loc := range node.path {
				successLocations = lib.AppendDistinct(successLocations, loc)
			}
			continue
		}

		// process north
		locNorth := lib.Location{X: node.current.X + North.X, Y: node.current.Y + North.Y}
		scoreNorth := GetFacingScore(node.face, North) + node.score
		if scoreNorth <= maxVal && !IsInPrevious(node.path, locNorth) && *course.maze.GetLocation(locNorth).Get() != '#' {
			newPath := make([]lib.Location, len(node.path))
			copy(newPath, node.path)
			newPath = append(newPath, node.current)
			processingQueue.Enqueue(BFSNode{path: newPath, current: locNorth, face: North, score: scoreNorth})
		}
		// process south
		locSouth := lib.Location{X: node.current.X + South.X, Y: node.current.Y + South.Y}
		scoreSouth := GetFacingScore(node.face, South) + node.score
		if scoreSouth <= maxVal && !IsInPrevious(node.path, locSouth) && *course.maze.GetLocation(locSouth).Get() != '#' {
			newPath := make([]lib.Location, len(node.path))
			copy(newPath, node.path)
			newPath = append(newPath, node.current)
			processingQueue.Enqueue(BFSNode{path: newPath, current: locSouth, face: South, score: scoreSouth})
		}
		// process east
		locEast := lib.Location{X: node.current.X + East.X, Y: node.current.Y + East.Y}
		scoreEast := GetFacingScore(node.face, East) + node.score
		if scoreEast <= maxVal && !IsInPrevious(node.path, locEast) && *course.maze.GetLocation(locEast).Get() != '#' {
			newPath := make([]lib.Location, len(node.path))
			copy(newPath, node.path)
			newPath = append(newPath, node.current)
			processingQueue.Enqueue(BFSNode{path: newPath, current: locEast, face: East, score: scoreEast})
		}
		// process west
		locWest := lib.Location{X: node.current.X + West.X, Y: node.current.Y + West.Y}
		scoreWest := GetFacingScore(node.face, West) + node.score
		if scoreWest <= maxVal && !IsInPrevious(node.path, locWest) && *course.maze.GetLocation(locWest).Get() != '#' {
			newPath := make([]lib.Location, len(node.path))
			copy(newPath, node.path)
			newPath = append(newPath, node.current)
			processingQueue.Enqueue(BFSNode{path: newPath, current: locWest, face: West, score: scoreWest})
		}
	}

	return successLocations
}

type DJKNode struct {
	score   int
	setFace lib.Location
	parents []lib.Location
	loc     lib.Location
}

type DJKProcess struct {
	node *DJKNode
	face lib.Location
}

func (n *DJKNode) GetScoreOut(face lib.Location) int {
	// if n.setFace.Equal(face) {
	// 	return n.score - 1000
	// }
	return n.score
}

func (n *DJKNode) GetScoreIn(face lib.Location) int {
	if n.setFace.Equal(face) {
		return n.score
	}
	return n.score - 1000
}

func GetFace(face lib.Location) string {
	if face.Equal(North) {
		return "N"
	}
	if face.Equal(South) {
		return "S"
	}
	if face.Equal(West) {
		return "W"
	}
	if face.Equal(East) {
		return "E"
	}
	return "U"
}

func (n *DJKNode) UpdateScore(newScore int, face lib.Location) {
	n.score = newScore
	n.setFace = face
}

func lookAhead(course Course, djkGrid lib.Grid[DJKNode], process DJKProcess, locskip lib.Location, dir lib.Location) bool {
	nextLoc := lib.Location{X: locskip.X + dir.X, Y: locskip.Y + dir.Y}
	nextItem := *course.maze.GetLocation(nextLoc).Get()
	if nextItem != '#' {
		score := djkGrid.GetLocation(nextLoc).Get().GetScoreIn(process.face)
		scoreOut := 2 + process.node.GetScoreOut(process.face)
		locSkipScore := djkGrid.GetLocation(locskip).Get().GetScoreOut(process.face)
		return score >= scoreOut && lib.Abs(locSkipScore-scoreOut) <= 1001
	}
	return false
}

func FindAllPathsDJK(course Course, maxVal int) []lib.Location {
	djkGrid := lib.GetGrid[DJKNode](course.maze.Rows(), course.maze.Columns())
	for i := range djkGrid.Rows() {
		for j := range djkGrid.Columns() {
			item := djkGrid.Get(i, j).Get()
			item.score = math.MaxInt32
			item.loc = lib.Location{X: i, Y: j}
			item.parents = []lib.Location{}
		}
	}
	node := djkGrid.GetLocation(course.start).Get()
	node.UpdateScore(0, East)
	queue := lib.NewQueue[DJKProcess]()
	queue.Enqueue(DJKProcess{node: node, face: East})
	for !queue.IsEmpty() {
		process := queue.Dequeue().Get()
		// walk all nodes and set scores + add parents

		if !process.face.Equal(South) {
			// check north
			locNorth := lib.Location{X: process.node.loc.X + North.X, Y: process.node.loc.Y + North.Y}
			scoreNorth := GetFacingScore(process.face, North) + process.node.GetScoreOut(process.face)
			if *course.maze.GetLocation(locNorth).Get() != '#' {
				nodeNorth := djkGrid.GetLocation(locNorth).Get()
				if (scoreNorth == nodeNorth.GetScoreIn(process.face) || lookAhead(course, djkGrid, process, locNorth, North)) && !IsInPrevious(nodeNorth.parents, process.node.loc) {
					nodeNorth.parents = append(nodeNorth.parents, process.node.loc)
				}
				if scoreNorth < nodeNorth.GetScoreIn(process.face) {
					nodeNorth.UpdateScore(scoreNorth, North)
					nodeNorth.parents = []lib.Location{process.node.loc}
					queue.Enqueue(DJKProcess{node: nodeNorth, face: North})
				}
			}
		}

		if !process.face.Equal(North) {
			// check south
			locSouth := lib.Location{X: process.node.loc.X + South.X, Y: process.node.loc.Y + South.Y}
			scoreSouth := GetFacingScore(process.face, South) + process.node.GetScoreOut(process.face)
			if *course.maze.GetLocation(locSouth).Get() != '#' {
				nodeSouth := djkGrid.GetLocation(locSouth).Get()
				if (scoreSouth == nodeSouth.GetScoreIn(process.face) || lookAhead(course, djkGrid, process, locSouth, South)) && !IsInPrevious(nodeSouth.parents, process.node.loc) {
					nodeSouth.parents = append(nodeSouth.parents, process.node.loc)
				}
				if scoreSouth < nodeSouth.GetScoreIn(process.face) {
					nodeSouth.UpdateScore(scoreSouth, South)
					nodeSouth.parents = []lib.Location{process.node.loc}
					queue.Enqueue(DJKProcess{node: nodeSouth, face: South})
				}
			}
		}

		if !process.face.Equal(West) {
			// check east
			locEast := lib.Location{X: process.node.loc.X + East.X, Y: process.node.loc.Y + East.Y}
			scoreEast := GetFacingScore(process.face, East) + process.node.GetScoreOut(process.face)
			if *course.maze.GetLocation(locEast).Get() != '#' {
				nodeEast := djkGrid.GetLocation(locEast).Get()
				if (scoreEast == nodeEast.GetScoreIn(process.face) || lookAhead(course, djkGrid, process, locEast, East)) && !IsInPrevious(nodeEast.parents, process.node.loc) {
					nodeEast.parents = append(nodeEast.parents, process.node.loc)
				}
				if scoreEast < nodeEast.GetScoreIn(process.face) {
					nodeEast.UpdateScore(scoreEast, East)
					nodeEast.parents = []lib.Location{process.node.loc}
					queue.Enqueue(DJKProcess{node: nodeEast, face: East})
				}
			}
		}

		if !process.face.Equal(East) {
			// check west
			locWest := lib.Location{X: process.node.loc.X + West.X, Y: process.node.loc.Y + West.Y}
			scoreWest := GetFacingScore(process.face, West) + process.node.GetScoreOut(process.face)
			if *course.maze.GetLocation(locWest).Get() != '#' {
				nodeWest := djkGrid.GetLocation(locWest).Get()
				if (scoreWest == nodeWest.GetScoreIn(process.face) || lookAhead(course, djkGrid, process, locWest, West)) && !IsInPrevious(nodeWest.parents, process.node.loc) {
					nodeWest.parents = append(nodeWest.parents, process.node.loc)
				}
				if scoreWest < nodeWest.GetScoreIn(process.face) {
					nodeWest.UpdateScore(scoreWest, West)
					nodeWest.parents = []lib.Location{process.node.loc}
					queue.Enqueue(DJKProcess{node: nodeWest, face: West})
				}
			}
		}
	}

	// processing finished
	successNodes := []lib.Location{}
	walkQueue := lib.NewQueue[lib.Location]()
	walkQueue.Enqueue(course.end)
	for !walkQueue.IsEmpty() {
		current := walkQueue.Dequeue().Get()
		successNodes = lib.AppendDistinct(successNodes, current)
		djkNode := djkGrid.GetLocation(current).Get()
		for _, parent := range djkNode.parents {
			walkQueue.Enqueue(parent)
		}
	}

	return successNodes
}

// func PrintScore(score map[lib.Location]int) string {
// 	return fmt.Sprintf("(N: %d, E: %d, S: %d, W: %d)", score[North], score[East], score[South], score[West])
// }

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
