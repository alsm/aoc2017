package main

import (
	"fmt"
	"strings"

	"github.com/alsm/aoc2017/aoc"
	. "github.com/alsm/aoc2017/aoc/collections"
	"golang.org/x/exp/slices"
)

func main() {
	input := aoc.MustSliceFromFile("day4.txt", func (l string) []string {
		return strings.Fields(l)
	})

	fmt.Println(do1(input))
	fmt.Println(do2(input))
}

func do1(i [][]string) int{
	return Count(i, func(p []string) bool {
		slices.Sort(p)
		return len(p) == len(slices.Compact(p))
	})
}

func do2(i [][]string) int{
	return Count(i, func(p []string) bool {
		n := Map(p, func(s string) string{
			x := []rune(s)
			slices.Sort(x)
			return string(x)
		})
		slices.Sort(n)
		return len(n) == len(slices.Compact(n))
	})
}