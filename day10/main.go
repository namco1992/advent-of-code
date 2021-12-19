package main

import (
	"fmt"
	"sort"

	"github.com/namco1992/aoc/util"
)

var (
	openMap           = map[int32]struct{}{'{': {}, '[': {}, '(': {}, '<': {}}
	closeOpenMap      = map[int32]int32{'}': '{', ']': '[', ')': '(', '>': '<'}
	openCloseMap      = map[int32]int32{'{': '}', '[': ']', '(': ')', '<': '>'}
	corruptedPointMap = map[int32]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	repairPointMap    = map[int32]int{')': 1, ']': 2, '}': 3, '>': 4}
)

func part1(input []string) int {
	var points int

	for _, s := range input {
		var stack []int32
		for _, c := range s {
			if _, ok := openMap[c]; ok {
				stack = append(stack, c)
			} else {
				peek := stack[len(stack)-1]
				if closeOpenMap[c] == peek {
					stack = stack[:len(stack)-1]
				} else {
					points += corruptedPointMap[c]
					break
				}
			}
		}
	}

	return points
}

func part2(input []string) int {
	var points []int

	for _, s := range input {
		var (
			stack       []int32
			isCorrupted bool
		)
		for _, c := range s {
			if _, ok := openMap[c]; ok {
				stack = append(stack, c)
			} else {
				peek := stack[len(stack)-1]
				if closeOpenMap[c] == peek {
					stack = stack[:len(stack)-1]
				} else {
					isCorrupted = true
					break
				}
			}
		}

		if isCorrupted {
			continue
		}
		var repairs []int32
		for i := len(stack) - 1; i >= 0; i-- {
			repairs = append(repairs, openCloseMap[stack[i]])
		}
		points = append(points, calcPoints(repairs))
	}
	sort.Ints(points)
	return points[len(points)/2]
}

func calcPoints(repairs []int32) int {
	var ret int
	for _, r := range repairs {
		ret = ret*5 + repairPointMap[r]
	}
	return ret
}

func main() {
	input := util.ScanFileToStrings("day10/input")
	fmt.Println(part2(input))
}
