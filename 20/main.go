package main

import (
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
