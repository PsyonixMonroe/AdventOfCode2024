package main

import (
	"testing"

	"github.com/PsyonixMonroe/AOCLib/lib"
	"github.com/go-playground/assert/v2"
)

func TestSimpleTriePart1(t *testing.T) {
	content := lib.ReadInput("test.txt")
	assert.NotEqual(t, content, "")
	patterns, words := ParseInput(content)
	assert.NotEqual(t, len(patterns), 0)
	assert.NotEqual(t, len(words), 0)

	root := BuildTrie(patterns, words)
	sum := CountUniqueWordsTrie(root)
	assert.Equal(t, sum, 6)
}

func TestSimpleIterPart1(t *testing.T) {
	content := lib.ReadInput("test.txt")
	assert.NotEqual(t, content, "")
	patterns, words := ParseInput(content)
	assert.NotEqual(t, len(patterns), 0)
	assert.NotEqual(t, len(words), 0)

	sum := CountUniqueWordsIter(patterns, words)
	assert.Equal(t, sum, 6)
}

func TestFullIterPart1(t *testing.T) {
	content := lib.ReadInput("input.txt")
	assert.NotEqual(t, content, "")
	patterns, words := ParseInput(content)
	assert.NotEqual(t, len(patterns), 0)
	assert.NotEqual(t, len(words), 0)

	sum := CountUniqueWordsIter(patterns, words)
	assert.Equal(t, sum, 240)
}

func TestSimpleTriePart2(t *testing.T) {
	content := lib.ReadInput("test.txt")
	assert.NotEqual(t, content, "")
	patterns, words := ParseInput(content)
	assert.NotEqual(t, len(patterns), 0)
	assert.NotEqual(t, len(words), 0)

	root := BuildTrie(patterns, words)
	sum := CountAllWordsTrie(root)
	assert.Equal(t, sum, 16)
}

func TestSimpleIterPart2(t *testing.T) {
	content := lib.ReadInput("test.txt")
	assert.NotEqual(t, content, "")
	patterns, words := ParseInput(content)
	assert.NotEqual(t, len(patterns), 0)
	assert.NotEqual(t, len(words), 0)

	sum := CountAllWordsIter(patterns, words)
	assert.Equal(t, sum, 16)
}

func TestFullIterPart2(t *testing.T) {
	content := lib.ReadInput("input.txt")
	assert.NotEqual(t, content, "")
	patterns, words := ParseInput(content)
	assert.NotEqual(t, len(patterns), 0)
	assert.NotEqual(t, len(words), 0)

	sum := CountAllWordsIter(patterns, words)
	assert.Equal(t, sum, 848076019766013)
}

/// Will run OOM
// func TestFullPart1(t *testing.T) {
// 	content := lib.ReadInput("input.txt")
// 	assert.NotEqual(t, content, "")
// 	patterns, words := ParseInput(content)
// 	assert.NotEqual(t, len(patterns), 0)
// 	assert.NotEqual(t, len(words), 0)
//
// 	root := BuildTrie(patterns, words)
// 	fmt.Fprintf(os.Stderr, "Finished Building Trie\n")
// 	sum := CountUniqueWordsTrie(root)
// 	assert.Equal(t, sum, 6)
// }
