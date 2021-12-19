package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/namco1992/aoc/util"
)

const length = 12

func scanToInts(input []string) []int64 {
	var ret []int64

	for _, s := range input {
		i, err := strconv.ParseInt(s, 2, 32)
		if err != nil {
			log.Fatal(err)
		}
		ret = append(ret, i)
	}

	return ret
}

func part1(input []int64) int {
	var gamma, epsilon int
	total := len(input)

	for i := 0; i < length; i++ {
		var count int
		for j := 0; j < total; j++ {
			if (input[j] & (1 << i)) > 0 {
				count++
			}
		}

		if count > total/2 {
			gamma += 1 << i
		} else {
			epsilon += 1 << i
		}
	}
	return gamma * epsilon
}

func part2(input []int64) int64 {
	candidates := input[:]
	for i := length - 1; i >= 0; i-- {
		var ones, zeros []int64
		for _, n := range candidates {
			if (n & (1 << i)) > 0 {
				ones = append(ones, n)
			} else {
				zeros = append(zeros, n)
			}
		}
		if len(ones) < len(zeros) { // change to >= for majority
			candidates = ones
		} else {
			candidates = zeros
		}
		if len(candidates) == 1 {
			return candidates[0]
		}
	}
	return -1
}

func main() {
	input := util.ScanFileToStrings("day3/input")

	fmt.Println(part2(scanToInts(input)))
}
