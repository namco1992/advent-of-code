package main

import (
	"fmt"

	"github.com/namco1992/aoc/util"
)

func part1(input []int) int {
	var cnt int
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			cnt++
		}
	}
	return cnt
}

func part2(input []int) int {
	var cnt int
	prevSum := input[2] + input[1] + input[0]

	for i := 3; i < len(input); i++ {
		currSum := input[i] + input[i-1] + input[i-2]
		if currSum > prevSum {
			cnt++
		}
		prevSum = currSum
	}
	return cnt
}

func main() {
	input := util.ScanFileToInts("day1/input")

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
