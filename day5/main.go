package main

import (
	"fmt"

	"github.com/alsm/aoc2017/aoc"
	//. "github.com/alsm/aoc2017/aoc/collections"
	"golang.org/x/exp/slices"
)

func main() {
	input := aoc.MustSliceFromFile("day5.txt", aoc.ReadInts)

	fmt.Println(do1(slices.Clone(input)))
	fmt.Println(do2(slices.Clone(input)))
}

func do1(s []int64) int64 {
	var ip int64
	var step int64

	for step =0; ip < int64(len(s)) && ip >= 0; step++ {
		nip := ip + s[ip]
		s[ip] += 1
		ip = nip
	}

	return step
}

func do2(s []int64) int64 {
	var ip int64
	var step int64

	for step =0; ip < int64(len(s)) && ip >= 0; step++ {
		nip := ip + s[ip]
		if s[ip] >= 3 {
			s[ip] -= 1
		} else {
			s[ip] += 1
		}
		ip = nip
	}

	return step
}