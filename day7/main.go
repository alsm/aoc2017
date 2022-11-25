package main

import (
	"fmt"
	"strings"

	"github.com/alsm/aoc2017/aoc"
	. "github.com/alsm/aoc2017/aoc/collections"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type Prog struct {
	Weight int
	Carries []string
	CarriedBy string
}

func (p *Prog) TotalWeight(t map[string]*Prog) int {
	total := p.Weight
	for _, n := range p.Carries {
		total += t[n].TotalWeight(t)
	}

	return total
}

func main() {
	input := make(map[string]*Prog)

	for _, s := range aoc.MustReadInput("day7.txt") {
		var name string
		var weight int
		var carries []string

		fmt.Sscanf(s, "%s (%d)", &name, &weight)
		if strings.Contains(s, " -> ") {
			carries = strings.Split(strings.Split(s, " -> ")[1], ", ")
		}
		if _, ok := input[name]; ok {
			input[name].Weight = weight
			input[name].Carries = carries
		} else {
			input[name] = &Prog{
				Weight: weight,
				Carries: carries,
			}
		}
		for _, c := range carries {
			if _, ok := input[c]; ok {
				input[c].CarriedBy = name
			} else {
				input[c] = &Prog {
					CarriedBy: name,
				}
			}
		}
	}

	fmt.Println(do1(input))
	fmt.Println(do2(do1(input), input))
}

func do1(i map[string]*Prog) string {
	for n, p := range i {
		if p.CarriedBy == "" {
			return n
		}
	}

	return "error"
}

func isBalanced(n string, i map[string]*Prog) (bool, string, int) {
	count := MapM(i[n].Carries, func(x string) (string, int) {
		return x, i[x].TotalWeight(i)
	})

	if len(slices.Compact(maps.Values(count))) == 1 {
		return true, "", 0
	}

	tally := Tally(maps.Values(count))
	odd := KeyWithValue(count, KeyWithValue(tally, Min(maps.Values(tally))))

	return false, odd, KeyWithMaxValue(tally)
}

func do2(base string, i map[string]*Prog) int {	
	balanced, odd := isBalanced(base, i)
	
	if !balanced {
		return do2(odd, i)
	}

	fmt.Println(target)
	return i[base].Weight - (i[base].TotalWeight(i) - target)
}