package main

import (
	"fmt"

	"github.com/alsm/aoc2017/aoc"
	. "github.com/alsm/aoc2017/aoc/collections"
)

func main() {
	input := aoc.MustSliceFromFile("day1.txt", aoc.ReadInts)

	fmt.Println(do1(append(input, input[0])))
	fmt.Println(do2(input))
}

func do1(input []int64) int64 {
	return Sum(Map(Cons(input, 1), func(x [2]int64) int64 {
		if x[0] == x[1] {
			return x[0]
		}
		return 0
	}))
}

func do2(input []int64) int64 {
	half := len(input) / 2
	return Sum(Map(Cons(append(input, input[:half]...), half), func(x [2]int64) int64 {
		if x[0] == x[1] {
			return x[0]
		}
		return 0
	}))
}
