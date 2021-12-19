package main

import (
	"fmt"

	"github.com/namco1992/aoc/util"
)

func part1(m util.Matrix) int {
	var cnt int

	for k := 0; k < 100; k++ {
		flashed := util.NewMatrix(m.X, m.Y)

		for j := 0; j < m.Y; j++ {
			for i := 0; i < m.X; i++ {
				calcEnergy(m, flashed, i, j, &cnt)
			}
		}
	}
	return cnt
}

func part2(m util.Matrix) int {
	var steps int

	for {
		flashed := util.NewMatrix(m.X, m.Y)
		var cnt int
		for j := 0; j < m.Y; j++ {
			for i := 0; i < m.X; i++ {
				calcEnergy(m, flashed, i, j, &cnt)
			}
		}
		steps++
		if cnt == m.X*m.Y {
			break
		}
	}
	return steps
}

func calcEnergy(input, flashed util.Matrix, i, j int, cnt *int) {
	if flashed.M[j][i] == 1 {
		return
	}
	if input.M[j][i] < 9 {
		input.M[j][i] += 1
	} else {
		input.M[j][i] = 0
		flashed.M[j][i] = 1
		*cnt += 1
		getNeighbours(input, flashed, i, j, cnt)
	}
}

func getNeighbours(input, flashed util.Matrix, i, j int, cnt *int) {
	x, y := input.X, input.Y
	if i > 0 {
		calcEnergy(input, flashed, i-1, j, cnt)
		if j > 0 {
			calcEnergy(input, flashed, i-1, j-1, cnt)
		}
		if j < y-1 {
			calcEnergy(input, flashed, i-1, j+1, cnt)
		}
	}
	if j > 0 {
		calcEnergy(input, flashed, i, j-1, cnt)
	}
	if i < x-1 {
		calcEnergy(input, flashed, i+1, j, cnt)
		if j > 0 {
			calcEnergy(input, flashed, i+1, j-1, cnt)
		}
		if j < y-1 {
			calcEnergy(input, flashed, i+1, j+1, cnt)
		}
	}
	if j < y-1 {
		calcEnergy(input, flashed, i, j+1, cnt)
	}
}

func main() {
	m := util.ConvertToMatrix(util.ScanFileToStrings("day11/input"))
	fmt.Println(part2(m))
}
