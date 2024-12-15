package main

import (
	"testing"

	"github.com/PsyonixMonroe/AOCLib/lib"
	"github.com/go-playground/assert/v2"
)

func TestSimplePart1(t *testing.T) {
	content := lib.ReadInput("test.txt")
	assert.NotEqual(t, content, "")

	shop, moves := ParseShop(content)
	assert.NotEqual(t, len(moves), 0)
	ProcessMoves(&shop, moves)
	sum := ScoreShop(shop)

	assert.Equal(t, sum, 2028)
}

func TestSimple2Part1(t *testing.T) {
	content := lib.ReadInput("test2.txt")
	assert.NotEqual(t, content, "")

	shop, moves := ParseShop(content)
	assert.NotEqual(t, len(moves), 0)
	ProcessMoves(&shop, moves)
	sum := ScoreShop(shop)

	assert.Equal(t, sum, 10092)
}

func TestScoreShop(t *testing.T) {
	content := "#######\n#...O..\n#......\n\n^"
	shop, moves := ParseShop(content)
	assert.Equal(t, *shop.shop.Get(1, 4).Get(), 'O')
	assert.Equal(t, len(moves), 1)

	score := ScoreShop(shop)
	assert.Equal(t, score, 104)
}

func TestFullPart1(t *testing.T) {
	content := lib.ReadInput("input.txt")
	assert.NotEqual(t, content, "")

	shop, moves := ParseShop(content)
	assert.NotEqual(t, len(moves), 0)
	ProcessMoves(&shop, moves)
	sum := ScoreShop(shop)

	assert.Equal(t, sum, 1492518)
}

func TestSimplePart2(t *testing.T) {
	content := lib.ReadInput("test.txt")
	assert.NotEqual(t, content, "")

	shop, moves := ParseShop(content)
	shop = GetWideShop(shop)
	shop.cursor = FindCursor(shop.shop)
	assert.NotEqual(t, len(moves), 0)
	ProcessMovesWide(&shop, moves)
	sum := ScoreShopWide(shop)

	assert.Equal(t, sum, 1751)
}

func TestSimple2Part2(t *testing.T) {
	content := lib.ReadInput("test2.txt")
	assert.NotEqual(t, content, "")

	shop, moves := ParseShop(content)
	shop = GetWideShop(shop)
	shop.cursor = FindCursor(shop.shop)
	assert.NotEqual(t, len(moves), 0)
	ProcessMovesWide(&shop, moves)
	sum := ScoreShopWide(shop)

	assert.Equal(t, sum, 9021)
}

func TestFullPart2(t *testing.T) {
	content := lib.ReadInput("input.txt")
	assert.NotEqual(t, content, "")

	shop, moves := ParseShop(content)
	shop = GetWideShop(shop)
	shop.cursor = FindCursor(shop.shop)
	// shop.shop.Print(lib.PrintRune)
	assert.NotEqual(t, len(moves), 0)
	ProcessMovesWide(&shop, moves)
	sum := ScoreShopWide(shop)

	assert.Equal(t, sum, 1512860)
}
