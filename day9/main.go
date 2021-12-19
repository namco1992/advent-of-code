package main

import (
	"fmt"
	"sort"

	"github.com/namco1992/aoc/util"
)

func part1(input []string) int {
	matrix, x, y := util.ConvertToMatrix(input)
	var cnt int
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			curr := matrix[j][i]
			// 0 must be the low point
			if curr == 0 {
				cnt += 1
				continue
			}

			// 9 must not be the low point
			if curr == 9 {
				continue
			}

			neighbours := getNeighbours(matrix, i, j, x, y)
			isLow := true
			for _, n := range neighbours {
				if curr >= n {
					isLow = false
					break
				}
			}
			if isLow {
				cnt += curr + 1
			}
		}
	}

	return cnt
}

func part2(input []string) int {
	matrix, x, y := util.ConvertToMatrix(input)
	var basins []int
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			curr := matrix[j][i]
			// 0 must be the low point
			if curr == 0 {
				size := 0
				visited := make([][]int, y)
				for i := 0; i < y; i++ {
					visited[i] = make([]int, x)
				}
				getBasinSize(matrix, visited, i, j, x, y, &size)
				basins = append(basins, size)
				continue
			}

			// 9 must not be the low point
			if curr == 9 {
				continue
			}

			neighbours := getNeighbours(matrix, i, j, x, y)
			isLow := true
			for _, n := range neighbours {
				if curr >= n {
					isLow = false
					break
				}
			}
			if isLow {
				size := 0
				visited := make([][]int, y)
				for i := 0; i < y; i++ {
					visited[i] = make([]int, x)
				}
				getBasinSize(matrix, visited, i, j, x, y, &size)
				basins = append(basins, size)
			}
		}
	}

	sort.Ints(basins)
	l := len(basins)
	return basins[l-1] * basins[l-2] * basins[l-3]
}

func getBasinSize(matrix, visited [][]int, i, j, x, y int, prev *int) {
	if i < 0 || j < 0 || i >= x || j >= y {
		return
	}
	if visited[j][i] == 1 {
		return
	}
	curr := matrix[j][i]
	if curr == 9 {
		return
	}
	*prev += 1
	visited[j][i] = 1

	getBasinSize(matrix, visited, i-1, j, x, y, prev)
	getBasinSize(matrix, visited, i, j-1, x, y, prev)
	getBasinSize(matrix, visited, i+1, j, x, y, prev)
	getBasinSize(matrix, visited, i, j+1, x, y, prev)

	return
}

func getNeighbours(matrix [][]int, i, j, x, y int) []int {
	var ret []int
	if i > 0 {
		ret = append(ret, matrix[j][i-1])
	}
	if j > 0 {
		ret = append(ret, matrix[j-1][i])
	}
	if i < x-1 {
		ret = append(ret, matrix[j][i+1])
	}
	if j < y-1 {
		ret = append(ret, matrix[j+1][i])
	}
	return ret
}

func main() {
	input := util.ScanFileToStrings("day9/input")
	fmt.Println(part2(input))
}
