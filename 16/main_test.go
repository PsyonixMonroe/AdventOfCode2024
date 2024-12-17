package main

import (
	"testing"

	"github.com/PsyonixMonroe/AOCLib/lib"
	"github.com/go-playground/assert/v2"
)

func TestSimple3Part1(t *testing.T) {
	content := lib.ReadInput("test3.txt")
	assert.NotEqual(t, content, "")

	course := ParseCourse(content)

	score, _ := FindBestPath(course)
	assert.Equal(t, score, 2008)
}

func TestSimplePart1(t *testing.T) {
	content := lib.ReadInput("test.txt")
	assert.NotEqual(t, content, "")

	course := ParseCourse(content)

	score, _ := FindBestPath(course)
	assert.Equal(t, score, 7036)
}

func TestSimple2Part1(t *testing.T) {
	content := lib.ReadInput("test2.txt")
	assert.NotEqual(t, content, "")

	course := ParseCourse(content)

	score, _ := FindBestPath(course)
	assert.Equal(t, score, 11048)
}

func TestFullPart1(t *testing.T) {
	content := lib.ReadInput("input.txt")
	assert.NotEqual(t, content, "")

	course := ParseCourse(content)

	score, _ := FindBestPath(course)
	assert.Equal(t, score, 98416)
}

// func TestSimple3Part2(t *testing.T) {
// 	content := lib.ReadInput("test3.txt")
// 	assert.NotEqual(t, content, "")
//
// 	course := ParseCourse(content)
//
// 	score, endNode := FindBestPath(course)
// 	allLocs := FindAllPaths(course, content, score, endNode, []lib.Location{})
// 	assert.Equal(t, score, 2008)
// 	assert.Equal(t, len(allLocs), 9)
// }
//
// func TestSimplePart2(t *testing.T) {
// 	content := lib.ReadInput("test.txt")
// 	assert.NotEqual(t, content, "")
//
// 	course := ParseCourse(content)
//
// 	score, endNode := FindBestPath(course)
// 	allLocs := FindAllPaths(course, content, score, endNode, []lib.Location{})
// 	assert.Equal(t, score, 7036)
// 	assert.Equal(t, len(allLocs), 45)
// }
//
// func TestSimple2Part2(t *testing.T) {
// 	content := lib.ReadInput("test2.txt")
// 	assert.NotEqual(t, content, "")
//
// 	course := ParseCourse(content)
//
// 	score, endNode := FindBestPath(course)
// 	allLocs := FindAllPaths(course, content, score, endNode, []lib.Location{})
// 	assert.Equal(t, score, 11048)
// 	assert.Equal(t, len(allLocs), 64)
// }

// func TestFullPart2(t *testing.T) {
// 	content := lib.ReadInput("input.txt")
// 	assert.NotEqual(t, content, "")
//
// 	course := ParseCourse(content)
//
// 	score, endNode := FindBestPath(course)
// 	allLocs := FindAllPaths(course, content, score, endNode, []lib.Location{})
// 	assert.Equal(t, score, 98416)
// 	assert.Equal(t, len(allLocs), 9)
// }

// func TestSimple3Part2BFS(t *testing.T) {
// 	content := lib.ReadInput("test3.txt")
// 	assert.NotEqual(t, content, "")
//
// 	course := ParseCourse(content)
//
// 	score, _ := FindBestPath(course)
// 	allLocs := FindAllPathsBFS(course, score)
// 	assert.Equal(t, score, 2008)
// 	assert.Equal(t, len(allLocs), 9)
// }
//
// func TestSimplePart2BFS(t *testing.T) {
// 	content := lib.ReadInput("test.txt")
// 	assert.NotEqual(t, content, "")
//
// 	course := ParseCourse(content)
//
// 	score, _ := FindBestPath(course)
// 	allLocs := FindAllPathsBFS(course, score)
// 	assert.Equal(t, score, 7036)
// 	assert.Equal(t, len(allLocs), 45)
// }
//
// func TestSimple2Part2BFS(t *testing.T) {
// 	content := lib.ReadInput("test2.txt")
// 	assert.NotEqual(t, content, "")
//
// 	course := ParseCourse(content)
//
// 	score, _ := FindBestPath(course)
// 	allLocs := FindAllPathsBFS(course, score)
// 	assert.Equal(t, score, 11048)
// 	assert.Equal(t, len(allLocs), 64)
// }

// func TestFullPart2BFS(t *testing.T) {
// 	content := lib.ReadInput("input.txt")
// 	assert.NotEqual(t, content, "")
//
// 	course := ParseCourse(content)
//
// 	score, _ := FindBestPath(course)
// 	allLocs := FindAllPathsBFS(course, score)
// 	assert.Equal(t, score, 98416)
// 	assert.Equal(t, len(allLocs), 9)
// }

func TestSimple3Part2DJK(t *testing.T) {
	content := lib.ReadInput("test3.txt")
	assert.NotEqual(t, content, "")

	course := ParseCourse(content)

	score, _ := FindBestPath(course)
	allLocs := FindAllPathsDJK(course, score)
	assert.Equal(t, score, 2008)
	assert.Equal(t, len(allLocs), 9)
}

func TestSimplePart2DJK(t *testing.T) {
	content := lib.ReadInput("test.txt")
	assert.NotEqual(t, content, "")

	course := ParseCourse(content)

	score, _ := FindBestPath(course)
	allLocs := FindAllPathsDJK(course, score)
	assert.Equal(t, score, 7036)
	assert.Equal(t, len(allLocs), 45)
}

func TestSimple2Part2DJK(t *testing.T) {
	content := lib.ReadInput("test2.txt")
	assert.NotEqual(t, content, "")

	course := ParseCourse(content)

	score, _ := FindBestPath(course)
	allLocs := FindAllPathsDJK(course, score)
	assert.Equal(t, score, 11048)
	assert.Equal(t, len(allLocs), 64)
}

func TestFullPart2DJK(t *testing.T) {
	content := lib.ReadInput("input.txt")
	assert.NotEqual(t, content, "")

	course := ParseCourse(content)

	score, _ := FindBestPath(course)
	allLocs := FindAllPathsDJK(course, score)
	assert.Equal(t, len(allLocs), 471)
}
