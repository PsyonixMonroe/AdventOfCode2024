package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	before string
	after  string
}

func (r Rule) Apply(update Update) bool {
	idxBefore := strings.Index(update.pages, r.before)
	idxAfter := strings.Index(update.pages, r.after)
	if idxBefore == -1 || idxAfter == -1 {
		return true
	}
	return idxBefore < idxAfter
}

func (r Rule) String() string {
	return fmt.Sprintf("Rule{before:%s,after:%s}", r.before, r.after)
}

type Update struct {
	pages  string
	middle int
}

func (u Update) String() string {
	return fmt.Sprintf("Update{pages:%s | middle: %d}", u.pages, u.middle)
}

func ReadInput(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read file %s: %v\n", filename, err)
		return ""
	}

	return string(content)
}

func ParseInput(content string) ([]Rule, []Update) {
	rules := []Rule{}
	updates := []Update{}

	for _, raw := range strings.Split(content, "\n") {
		line := strings.Trim(raw, "\n")
		if line == "" {
			continue
		}

		if strings.ContainsRune(line, '|') {
			rule := ParseRule(line)
			rules = append(rules, rule)
		} else {
			update := ParseUpdate(line)
			updates = append(updates, update)
		}
	}

	return rules, updates
}

func CountGoodUpdates(rules []Rule, updates []Update) int {
	middleSum := 0
	for _, update := range updates {
		pass := true
		for _, rule := range rules {
			pass = pass && rule.Apply(update)
			if !pass {
				break
			}
		}

		if pass {
			middleSum += update.middle
		}
	}

	return middleSum
}

func ParseRule(line string) Rule {
	split := strings.Split(line, "|")
	if len(split) != 2 {
		fmt.Fprintf(os.Stderr, "Unknown Format for Rule: %s\n", line)
		return Rule{}
	}

	return Rule{before: split[0], after: split[1]}
}

func ParseUpdate(line string) Update {
	splits := strings.Split(line, ",")
	if len(splits)%2 != 1 {
		fmt.Fprintf(os.Stderr, "Even number of splits for %s\n", line)
		return Update{pages: line, middle: 0}
	}
	middleStr := splits[len(splits)/2]
	middle, err := strconv.Atoi(middleStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse middle %s\n", middleStr)
		return Update{pages: line, middle: 0}
	}

	return Update{pages: line, middle: middle}
}
