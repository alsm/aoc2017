package main

import (
	"fmt"

	"github.com/alsm/aoc2017/aoc"
)

// Lattice calculates the number of items in the nth ring
// of an expanding grid
func Lattice(n int64) int64 {
	return aoc.IPow(0, n) + 8*n
}

func main() {
	fmt.Println(do1(368078))
}

func do1(loc int64) int64{
	var layer int64
	for ; ;layer++ {
		count := Lattice(layer)
		if loc - count < 0 {
			break
		}
		loc -= count
	}
	side := layer * 2
	for loc - side > 0 {
		loc -= side
	}
	return (layer + aoc.Abs(loc - layer))
}

func do2() {
	// https://oeis.org/A141481/b141481.txt
}