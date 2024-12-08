package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	before int
	after  int
}

func (r Rule) Apply(update Update) bool {
	idxBefore := update.Index(r.before)
	idxAfter := update.Index(r.after)
	if idxBefore == -1 || idxAfter == -1 {
		return true
	}
	return idxBefore < idxAfter
}

func (r Rule) String() string {
	return fmt.Sprintf("Rule{before:%d,after:%d}", r.before, r.after)
}

type Update struct {
	pages []int
}

func (u Update) Index(page int) int {
	for i, p := range u.pages {
		if p == page {
			return i
		}
	}

	return -1
}

func (u Update) GetMiddle() int {
	return u.pages[len(u.pages)/2]
}

func (u Update) String() string {
	return fmt.Sprintf("Update{pages:%v | middle: %d}", u.pages, u.GetMiddle())
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
			middleSum += update.GetMiddle()
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

	b, err := strconv.Atoi(split[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse before from %s\n", split[0])
	}

	a, err := strconv.Atoi(split[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse after from %s\n", split[1])
	}

	return Rule{before: b, after: a}
}

func ParseUpdate(line string) Update {
	splits := strings.Split(line, ",")
	pages := []int{}
	if len(splits)%2 != 1 {
		fmt.Fprintf(os.Stderr, "Even number of splits for %s\n", line)
		return Update{pages: pages}
	}

	for _, page := range splits {
		p, err := strconv.Atoi(page)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to parse page %s\n", page)
		}
		pages = append(pages, p)
	}

	return Update{pages: pages}
}
