package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/PsyonixMonroe/AOCLib/lib"
)

func ParseInput(content string) ([]string, []string) {
	lines := strings.Split(strings.Trim(content, " \n\t"), "\n")
	if len(lines) < 3 {
		fmt.Fprintf(os.Stderr, "Unable to parse input content: %s\n", content)
		return []string{}, []string{}
	}

	words := []string{}
	patterns := []string{}

	for _, pattern := range strings.Split(strings.Trim(lines[0], "\n \t"), ",") {
		patterns = append(patterns, strings.Trim(pattern, ", \n\t"))
	}

	// lines[0] => patterns
	// lines[1] => blank
	// lines[2] => first word
	// lines[3] => second word...
	for _, line := range lines[2:] {
		word := strings.Trim(line, "\n \t")
		if word == "" {
			fmt.Fprintf(os.Stderr, "Not adding empty string to words")
			continue
		}
		words = append(words, word)
	}

	return patterns, words
}

func CountUniqueWordsIter(patterns []string, words []string) int {
	sum := 0
	for _, word := range words {
		if FindWordFromPrefix(patterns, "", word) {
			sum++
		}
	}

	return sum
}

func FindWordFromPrefix(patterns []string, prefix string, fullWord string) bool {
	if prefix == fullWord {
		return true
	}

	for _, pattern := range patterns {
		newPrefix := prefix + pattern
		if strings.HasPrefix(fullWord, newPrefix) {
			found := FindWordFromPrefix(patterns, newPrefix, fullWord)
			if found {
				// found a match, that is all we need
				return true
			}
		}
	}

	return false
}

type TrieNode struct {
	children map[string]TrieNode
	pattern  string
	prefix   string
}

func (t TrieNode) IsWordMatch() bool {
	return len(t.children) == 0 && t.pattern == ""
}

func (t TrieNode) IsRoot() bool {
	return len(t.children) > 0 && t.pattern == ""
}

func (t TrieNode) GetMatchingChildren(current string, fullWord string) []TrieNode {
	matchNodes := []TrieNode{}
	for child, node := range t.children {
		childString := current + child
		if strings.HasPrefix(fullWord, childString) {
			matchNodes = append(matchNodes, node)
		}
	}

	return matchNodes
}

func BuildTrie(patterns []string, words []string) TrieNode {
	root := TrieNode{children: make(map[string]TrieNode), pattern: "", prefix: ""}
	for _, word := range words {
		GetTrieChildren(patterns, "", word, &root)
	}

	return root
}

func GetTrieChildren(patterns []string, prefix string, fullWord string, current *TrieNode) {
	if prefix == fullWord {
		// found a leaf
		current.children[""] = TrieNode{children: make(map[string]TrieNode), pattern: "", prefix: prefix}
		return
	}

	for _, pattern := range patterns {
		childPrefix := prefix + pattern
		if strings.HasPrefix(fullWord, childPrefix) {
			childNode, found := current.children[pattern]
			if !found {
				childNode = TrieNode{children: make(map[string]TrieNode), pattern: pattern, prefix: childPrefix}
			}
			GetTrieChildren(patterns, childPrefix, fullWord, &childNode)
			current.children[pattern] = childNode
		}
	}
}

func CountAllWordsTrie(root TrieNode) int {
	work := lib.NewQueue[TrieNode]()
	work.Enqueue(root)
	count := 0
	for !work.IsEmpty() {
		current := work.Dequeue().Get()
		if current.IsWordMatch() {
			count++
			continue
		}
		for _, child := range current.children {
			work.Enqueue(child)
		}
	}

	return count
}

func CountUniqueWordsTrie(root TrieNode) int {
	work := lib.NewQueue[TrieNode]()
	work.Enqueue(root)
	words := []string{}
	for !work.IsEmpty() {
		current := work.Dequeue().Get()
		if current.IsWordMatch() {
			words = lib.AppendDistinct(words, current.prefix)
			continue
		}
		for _, child := range current.children {
			work.Enqueue(child)
		}
	}

	return len(words)
}
