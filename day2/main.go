package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/alsm/aoc2017/aoc"
	. "github.com/alsm/aoc2017/aoc/collections"
)

func main() {
	input := Map(aoc.MustReadInput("day2.txt"), func(v string) []int {
		return Map(strings.Split(v, " "), func(x string) int {
			n, err := strconv.Atoi(x)
			if err != nil {
				log.Fatalf("could not parse '%s' as int", x)
			}
			return n
		})
	})

	fmt.Println(do1(input))
	fmt.Println(do2(input))
}

func do1(input [][]int) int {
	return Sum(Map(input, func(x []int) int {
		min, max := MinMax(x)
		return max - min
	}))
}

func do2(input [][]int) int {
	return Sum(Map(input, func(x []int) int {
		return Sum(Map(Combinations(x), func(y [2]int) int {
			if y[0]%y[1] == 0 {
				return y[0] / y[1]
			}
			if y[1]%y[0] == 0 {
				return y[1] / y[0]
			}
			return 0
		}))
	}))
}

// func do1(input []Command) int {
// 	final := Reduce(input, [2]int{0, 0}, func(initial [2]int, v Command) [2]int {
// 		switch v.Direction {
// 		case "forward":
// 			initial[0] += v.Value
// 		case "down":
// 			initial[1] += v.Value
// 		case "up":
// 			initial[1] -= v.Value
// 		}

// 		return initial
// 	})

// 	return final[0] * final[1]
// }

// func do2(input []Command) int {
// 	final := Reduce(input, [3]int{0, 0, 0}, func(initial [3]int, v Command) [3]int {
// 		switch v.Direction {
// 		case "forward":
// 			initial[0] += v.Value
// 			initial[1] += v.Value * initial[2]
// 		case "down":
// 			initial[2] += v.Value
// 		case "up":
// 			initial[2] -= v.Value
// 		}

// 		return initial
// 	})

// 	return final[0] * final[1]
// }
