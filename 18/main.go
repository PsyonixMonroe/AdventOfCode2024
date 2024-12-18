package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/PsyonixMonroe/AOCLib/lib"
)

func ParseBytes(content string, rows int, columns int) (lib.Grid[rune], []lib.Location) {
	grid := lib.GetGrid[rune](rows, columns)
	locations := []lib.Location{}
	for _, line := range strings.Split(strings.Trim(content, "\n"), "\n") {
		coords := strings.Trim(line, "\n \t")
		splits := strings.Split(coords, ",")
		if len(splits) != 2 {
			fmt.Fprintf(os.Stderr, "Unable to parse splits from coords: %s\n", coords)
			continue
		}
		y, err := strconv.Atoi(splits[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to parse y coord: %s\n", splits[0])
			continue
		}

		x, err := strconv.Atoi(splits[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to parse x coord: %s\n", splits[0])
			continue
		}

		locations = append(locations, lib.Location{X: x, Y: y})
	}

	for i := range rows {
		for j := range columns {
			grid.Put('.', i, j)
		}
	}

	return grid, locations
}

func FindPath(grid lib.Grid[rune], start lib.Location, end lib.Location, bytes []lib.Location, numBytes int) []lib.Location {
	if len(bytes) <= numBytes {
		fmt.Fprintf(os.Stderr, "Unable to fill in bytes, not enough bytes provided\n")
		return []lib.Location{}
	}
	for i := range numBytes {
		byte := bytes[i]
		grid.Put('#', byte.X, byte.Y)
	}

	path := lib.FindPath(grid, start, end, isWall, lib.SimpleGridDistanceHVal, lib.SimpleNeighborGDiff)
	path, _ = lib.RemoveItem(path, start)
	return path
}

func FindBlockingByte(grid lib.Grid[rune], start lib.Location, end lib.Location, bytes []lib.Location, numBytes int) lib.Location {
	if len(bytes) <= numBytes {
		fmt.Fprintf(os.Stderr, "Unable to fill in bytes, not enough bytes provided\n")
		return lib.Location{X: -1, Y: -1}
	}
	for i := range numBytes {
		byte := bytes[i]
		grid.Put('#', byte.X, byte.Y)
	}

	path := lib.FindPath(grid, start, end, isWall, lib.SimpleGridDistanceHVal, lib.SimpleNeighborGDiff)
	for i, byte := range bytes {
		if i < numBytes {
			// already added these
			continue
		}
		grid.Put('#', byte.X, byte.Y)
		blocksCurrentBest, _ := lib.FindInArray(byte, path)
		if blocksCurrentBest != -1 {
			path = lib.FindPath(grid, start, end, isWall, lib.SimpleGridDistanceHVal, lib.SimpleNeighborGDiff)
			if len(path) == 0 {
				// found the last byte
				return byte
			}
		}
	}
	return lib.Location{X: -1, Y: -1}
}

func isWall(g lib.Grid[rune], loc lib.Location) bool {
	return g.IsOOBLocation(loc) || *g.GetLocation(loc).Get() == '#'
}
