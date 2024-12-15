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

func GetWideShop(shop Shop) Shop {
	newShopGrid := lib.GetGrid[rune](shop.shop.Rows(), shop.shop.Columns()*2)
	for i := range shop.shop.Rows() {
		for j := range shop.shop.Columns() {
			current := *shop.shop.Get(i, j).Get()
			if current == '#' {
				newShopGrid.Put('#', i, j*2)
				newShopGrid.Put('#', i, j*2+1)
			}
			if current == 'O' {
				newShopGrid.Put('[', i, j*2)
				newShopGrid.Put(']', i, j*2+1)
			}
			if current == '@' {
				newShopGrid.Put('@', i, j*2)
				newShopGrid.Put('.', i, j*2+1)
			}
			if current == '.' {
				newShopGrid.Put('.', i, j*2)
				newShopGrid.Put('.', i, j*2+1)
			}
		}
	}

	return Shop{shop: newShopGrid, cursor: lib.Location{X: 0, Y: 0}}
}

func ProcessMovesWide(shop *Shop, moves []Move) {
	for _, move := range moves {
		if tryMove(*shop, move, shop.cursor) {
			moveDirectionWide(shop, move, shop.cursor)
		}
	}
}

func tryMove(shop Shop, move Move, current lib.Location) bool {
	if move.X == 0 {
		return tryMoveHorizontal(shop, move, current)
	}
	return tryMoveVertical(shop, move, current)
}

func tryMoveHorizontal(shop Shop, move Move, current lib.Location) bool {
	nextLoc := lib.Location{X: current.X, Y: current.Y + move.Y}
	nextNextLoc := lib.Location{X: current.X, Y: current.Y + move.Y*2}
	if shop.shop.IsOOBLocation(nextLoc) {
		return false
	}
	nextItem := *shop.shop.GetLocation(nextLoc).Get()

	if nextItem == '#' {
		return false
	}

	if nextItem == '.' {
		return true
	}

	if nextItem == '[' || nextItem == ']' {
		return tryMoveHorizontal(shop, move, nextNextLoc)
	}

	fmt.Fprintf(os.Stderr, "[tryMoveHorizontal] Unkown item for Next Item: %s\n", string(nextItem))
	return false
}

func tryMoveVertical(shop Shop, move Move, current lib.Location) bool {
	adjLoc := lib.Location{X: current.X + move.X, Y: current.Y}

	if shop.shop.IsOOBLocation(adjLoc) {
		return false
	}

	adjItem := *shop.shop.GetLocation(adjLoc).Get()

	if adjItem == '.' {
		return true
	}

	if adjItem == '#' {
		return false
	}

	if adjItem == '[' {
		diagLoc := lib.Location{X: adjLoc.X, Y: adjLoc.Y + 1}
		return tryMoveVertical(shop, move, adjLoc) && tryMoveVertical(shop, move, diagLoc)
	}

	if adjItem == ']' {
		diagLoc := lib.Location{X: adjLoc.X, Y: adjLoc.Y - 1}
		return tryMoveVertical(shop, move, adjLoc) && tryMoveVertical(shop, move, diagLoc)
	}

	fmt.Fprintf(os.Stderr, "[tryMoveVertical] Unknown item at next location: %s\n", string(adjItem))
	return false
}

func moveDirectionWide(shop *Shop, move Move, current lib.Location) {
	if move.X == 0 {
		moveDirectionWideHorizontal(shop, move, current)
	}
	moveDirectionWideVertical(shop, move, current)
}

func moveDirectionWideHorizontal(shop *Shop, move Move, current lib.Location) {
	nextLoc := lib.Location{X: current.X, Y: current.Y + move.Y}
	nextNextLoc := lib.Location{X: current.X, Y: current.Y + move.Y*2}
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

	if nextItem == '[' || nextItem == ']' {
		nextNextItem := *shop.shop.GetLocation(nextNextLoc).Get()
		if nextNextItem != '.' {
			moveDirectionWideHorizontal(shop, move, nextNextLoc)
		}
		moveDirectionWideHorizontal(shop, move, nextLoc)
		moveDirectionWideHorizontal(shop, move, current)
		return
	}

	fmt.Fprintf(os.Stderr, "[moveDirectionWideHorizontal] Unkown item for Next Item: %s\n", string(nextItem))
}

func moveDirectionWideVertical(shop *Shop, move Move, current lib.Location) {
	adjLoc := lib.Location{X: current.X + move.X, Y: current.Y}
	adjItem := *shop.shop.GetLocation(adjLoc).Get()

	if adjItem == '.' {
		curItem := *shop.shop.GetLocation(current).Get()
		shop.shop.PutLocation(curItem, adjLoc)
		shop.shop.PutLocation('.', current)
		if curItem == '@' {
			shop.cursor = adjLoc
		}
		return
	}

	if adjItem == '[' {
		diagLoc := lib.Location{X: adjLoc.X, Y: adjLoc.Y + 1}
		moveDirectionWideVertical(shop, move, adjLoc)
		moveDirectionWideVertical(shop, move, diagLoc)
		moveDirectionWideVertical(shop, move, current)
		return
	}

	if adjItem == ']' {
		diagLoc := lib.Location{X: adjLoc.X, Y: adjLoc.Y - 1}
		moveDirectionWideVertical(shop, move, adjLoc)
		moveDirectionWideVertical(shop, move, diagLoc)
		moveDirectionWideVertical(shop, move, current)
		return
	}

	fmt.Fprintf(os.Stderr, "[moveDirectionWideVertical] Unknown item at next location: %s\n", string(adjItem))
}

func ScoreShopWide(shop Shop) int {
	score := 0
	for i := range shop.shop.Rows() {
		for j := range shop.shop.Columns() {
			slot := *shop.shop.Get(i, j).Get()
			if slot == '[' {
				score += 100*i + j
			}
		}
	}
	return score
}
