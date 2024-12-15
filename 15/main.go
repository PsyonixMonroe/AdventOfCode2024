package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/PsyonixMonroe/AOCLib/lib"
)

type Shop struct {
	shop   lib.Grid[rune]
	cursor lib.Location
}

type Move = lib.Location

func MoveFromRune(r rune) Move {
	switch r {
	case '^':
		return lib.Location{X: -1, Y: 0}
	case '>':
		return lib.Location{X: 0, Y: 1}
	case 'v':
		return lib.Location{X: 1, Y: 0}
	case '<':
		return lib.Location{X: 0, Y: -1}
	}

	// should not get here
	fmt.Fprintf(os.Stderr, "Unkown Rune provided: %s\n", string(r))
	return lib.Location{X: 100, Y: 100}
}

func ParseShop(content string) (Shop, []Move) {
	splits := strings.Split(content, "\n\n")
	if len(splits) != 2 {
		fmt.Fprintf(os.Stderr, "incorrect format for input\n")
		return Shop{shop: lib.Grid[rune]{}, cursor: lib.Location{X: 0, Y: 0}}, []Move{}
	}
	grid := lib.ParseRuneGrid(splits[0])
	moveList := lib.ParseRuneLine(splits[1])
	moves := []Move{}
	for _, r := range moveList {
		m := MoveFromRune(r)
		if m.X > 1 || m.X < -1 {
			continue
		}
		moves = append(moves, m)
	}
	cursor := FindCursor(grid)
	return Shop{shop: grid, cursor: cursor}, moves
}

func FindCursor(grid lib.Grid[rune]) lib.Location {
	for i := range grid.Rows() {
		for j := range grid.Columns() {
			current := *grid.Get(i, j).Get()
			if current == '@' {
				return lib.Location{X: i, Y: j}
			}
		}
	}

	fmt.Fprintf(os.Stderr, "Unable to find cursor\n")
	return lib.Location{X: 0, Y: 0}
}

func ProcessMoves(shop *Shop, moves []Move) {
	for _, move := range moves {
		moveDirection(shop, move, shop.cursor)
	}
}

func moveDirection(shop *Shop, dir Move, current lib.Location) {
	nextLoc := lib.Location{X: current.X + dir.X, Y: current.Y + dir.Y}
	if shop.shop.IsOOBLocation(nextLoc) {
		return
	}
	nextItem := *shop.shop.GetLocation(nextLoc).Get()
	if nextItem == '.' {
		curItem := *shop.shop.GetLocation(current).Get()
		shop.shop.PutLocation(curItem, nextLoc)
		shop.shop.PutLocation('.', current)
		if curItem == '@' {
			shop.cursor = nextLoc
		}
		return
	}
	if nextItem == 'O' {
		moveDirection(shop, dir, nextLoc)
		newNextItem := *shop.shop.GetLocation(nextLoc).Get()
		if newNextItem == '.' {
			// items have moved, we can now call this for the actual move logic
			moveDirection(shop, dir, current)
		}
		return
	}
	if nextItem == '#' {
		return
	}
	fmt.Fprintf(os.Stderr, "Unknown Item: %s\n", string(nextItem))
}

func ScoreShop(shop Shop) int {
	score := 0
	for i := range shop.shop.Rows() {
		for j := range shop.shop.Columns() {
			slot := *shop.shop.Get(i, j).Get()
			if slot == 'O' {
				score += 100*i + j
			}
		}
	}
	return score
}
