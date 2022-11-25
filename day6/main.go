package main

import (
	"fmt"
	"strings"

	"github.com/alsm/aoc2017/aoc"
	. "github.com/alsm/aoc2017/aoc/collections"
	"golang.org/x/exp/slices"
)

func main() {
	input := aoc.InputToInts(strings.Fields(aoc.MustReadInput("day6.txt")[0]))

	fmt.Println(do1(slices.Clone(input)))
	fmt.Println(do2(slices.Clone(input)))
}

func do1(s []int) int {
	loop := len(s)
	seen := make(map[string]struct{})

	for i := 1;; i++ {
		idx, v := MaxWithIndex(s)
		s[idx], idx = 0, idx+1
		for x := 0;x < v; x++ {
			s[(idx+x)%loop] += 1
		}
		st := Join(s, ",")
		if _, ok := seen[st]; ok {
			return i
		} else {
			seen[st] = struct{}{}
		}
	}
}

func do2(s []int) int {
	loop := len(s)
	seen := make(map[string]int)

	for i := 1;; i++ {
		idx, v := MaxWithIndex(s)
		s[idx], idx = 0, idx+1
		for x := 0;x < v; x++ {
			s[(idx+x)%loop] += 1
		}
		st := Join(s, ",")
		if x, ok := seen[st]; ok {
			return i-x
		} else {
			seen[st] = i
		}
	}
}