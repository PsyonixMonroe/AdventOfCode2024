package main

import (
	"context"
	"sync"

	"github.com/PsyonixMonroe/AOCLib/lib"
)

type RaceTrack struct {
	track lib.Grid[rune]
	start lib.Location
	end   lib.Location
}

func ParseInput(content string) RaceTrack {
	grid := lib.ParseRuneGrid(content)
	start, end := FindStartEnd(grid)
	return RaceTrack{track: grid, start: start, end: end}
}

func FindStartEnd(track lib.Grid[rune]) (lib.Location, lib.Location) {
	var start lib.Location
	var end lib.Location
	startFound := false
	endFound := false
	for i := range track.Rows() {
		for j := range track.Columns() {
			value := *track.Get(i, j).Get()
			if value == 'S' {
				start = lib.Location{X: i, Y: j}
				startFound = true
			}
			if value == 'E' {
				end = lib.Location{X: i, Y: j}
				endFound = true
			}
			if startFound && endFound {
				break
			}
		}
		if startFound && endFound {
			break
		}
	}

	return start, end
}

func CountGoodPaths(track RaceTrack, content string, minSavings int) int {
	basePath := lib.FindPath(track.track, track.start, track.end, isWall, lib.SimpleGridDistanceHVal, lib.SimpleNeighborGDiff)
	base := len(basePath)
	sum := 0
	for i := range track.track.Rows() {
		for j := range track.track.Columns() {
			locVal := *track.track.Get(i, j).Get()
			if locVal == '#' && openNeighbors(track.track, i, j, isWall) >= 2 {
				newGrid := lib.ParseRuneGrid(content)
				newGrid.Put('.', i, j)
				modPath := lib.FindPath(newGrid, track.start, track.end, isWall, lib.SimpleGridDistanceHVal, lib.SimpleNeighborGDiff)
				modPathLen := len(modPath)
				if base-modPathLen >= minSavings {
					sum++
				}
			}
		}
	}

	return sum
}

func openNeighbors(g lib.Grid[rune], i int, j int, isWall func(g lib.Grid[rune], loc lib.Location) bool) int {
	north := lib.Location{X: i - 1, Y: j}
	south := lib.Location{X: i + 1, Y: j}
	east := lib.Location{X: i, Y: j + 1}
	west := lib.Location{X: i, Y: j - 1}
	count := 0
	if !g.IsOOBLocation(north) && !isWall(g, north) {
		count++
	}
	if !g.IsOOBLocation(south) && !isWall(g, south) {
		count++
	}
	if !g.IsOOBLocation(east) && !isWall(g, east) {
		count++
	}
	if !g.IsOOBLocation(west) && !isWall(g, west) {
		count++
	}

	return count
}

func isWall(g lib.Grid[rune], loc lib.Location) bool {
	return *g.GetLocation(loc).Get() == '#'
}

type CheatProcReq struct {
	start lib.Location
	dist  int
}

func CountCheats(track RaceTrack, minSavings int) int {
	basePath := lib.FindPath(track.track, track.start, track.end, isWall, lib.SimpleGridDistanceHVal, lib.SimpleNeighborGDiff)
	base := len(basePath) - 1
	sum := 0
	startPositions := make(chan CheatProcReq, 100)
	endPositions := make(chan lib.Location, 100)
	procCtx, procCancel := context.WithCancel(context.Background())
	procwg := sync.WaitGroup{}
	procwg.Add(12)
	endPosCache := GetEmptyDistGrid(track.track)
	for range 12 {
		go func() {
			defer procwg.Done()
			done := false
			for {
				select {
				case <-procCtx.Done():
					done = true
				case s := <-startPositions:
					endPos := GetEndPositions(track, &endPosCache, s, minSavings, base)
					for _, end := range endPos {
						endPositions <- end
					}
				}
				if done {
					break
				}
			}
		}()
	}

	endCtx, endCancel := context.WithCancel(context.Background())
	endwg := sync.WaitGroup{}
	endwg.Add(1)
	go func() {
		defer endwg.Done()
		done := false
		for {
			select {
			case <-endCtx.Done():
				done = true
			case <-endPositions:
				sum++
			}
			if done {
				break
			}
		}
	}()

	// generate Start positions
	for i := range track.track.Rows() {
		for j := range track.track.Columns() {
			startCheatLoc := lib.Location{X: i, Y: j}
			if !isWall(track.track, startCheatLoc) {
				// we always start on a track at index -1
				startPath := lib.FindPath(track.track, track.start, startCheatLoc, isWall, lib.SimpleGridDistanceHVal, lib.SimpleNeighborGDiff)
				startPathLen := len(startPath) - 1
				if startPathLen > base-minSavings {
					// no improvement will get us where we need to go
					continue
				}
				startPositions <- CheatProcReq{start: startCheatLoc, dist: startPathLen}
			}
		}
	}

	// Finish all the work
	for len(startPositions) > 0 {
	}
	procCancel()
	procwg.Wait()

	for len(endPositions) > 0 {
	}
	endCancel()
	endwg.Wait()

	return sum
}

func GetEmptyDistGrid(grid lib.Grid[rune]) lib.Grid[int] {
	distGrid := lib.GetGrid[int](grid.Rows(), grid.Columns())
	for i := range distGrid.Rows() {
		for j := range distGrid.Columns() {
			distGrid.Put(-1, i, j)
		}
	}

	return distGrid
}

func GetEndPositions(track RaceTrack, distGrid *lib.Grid[int], req CheatProcReq, minSavings int, base int) []lib.Location {
	const cheatLength = 20
	directions := []lib.Location{{X: -1, Y: 1}, {X: 1, Y: 1}, {X: -1, Y: -1}, {X: 1, Y: -1}}
	endPositions := []lib.Location{}
	for i := range cheatLength + 1 {
		for j := range cheatLength + 1 {
			if (i + j) > cheatLength {
				continue
			}
			for _, dir := range directions {
				testLoc := lib.Location{X: req.start.X + i*dir.X, Y: req.start.Y + j*dir.Y}
				if !track.track.IsOOBLocation(testLoc) && !isWall(track.track, testLoc) && !LocInArray(testLoc, endPositions) {
					distGridVal := *distGrid.GetLocation(testLoc).Get()
					if distGridVal == -1 {
						// compute and save dist
						endPath := lib.FindPath(track.track, testLoc, track.end, isWall, lib.SimpleGridDistanceHVal, lib.SimpleNeighborGDiff)
						distGridVal = len(endPath) - 1
						distGrid.PutLocation(distGridVal, testLoc)
					}

					if req.dist+i+j+distGridVal <= (base - minSavings) {
						// fmt.Fprintf(os.Stderr, "Found Cheat: %d saves %d - Start:%v End:%v\nBaseDist: %d, reqDist: %d, i:%d, j:%d, endPathDist: %d\n", 1+i+j, base-(req.dist+i+j+distGridVal), req.start, testLoc, base, req.dist, i, j, distGridVal)
						endPositions = append(endPositions, testLoc)
						// } else {
						// if req.dist+i+j+distGridVal-2 <= (base - minSavings) {
						// 	fmt.Fprintf(os.Stderr, "Near Miss: %d saves %d - Start:%v End:%v\nBaseDist: %d, reqDist: %d, i:%d, j:%d, endPathDist: %d\n", 1+i+j, base-(req.dist+i+j+distGridVal), req.start, testLoc, base, req.dist, i, j, distGridVal)
						// }
					}
				}
			}
		}
	}

	return endPositions
}

func LocInArray(needle lib.Location, haystack []lib.Location) bool {
	for _, straw := range haystack {
		if straw.Equal(needle) {
			return true
		}
	}
	return false
}
