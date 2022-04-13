package util

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Matrix struct {
	M    [][]int
	X, Y int
}

func NewMatrix(x, y int) Matrix {
	m := make([][]int, y)
	for i := 0; i < y; i++ {
		m[i] = make([]int, x)
	}
	return Matrix{M: m, X: x, Y: y}
}

func ConvertToMatrix(input []string) Matrix {
	y := len(input)
	x := len(input[0])
	ret := make([][]int, y)
	for i := 0; i < y; i++ {
		ret[i] = make([]int, x)
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			n, err := strconv.Atoi(input[j][i : i+1])
			if err != nil {
				log.Fatal(err)
			}
			ret[j][i] = n
		}
	}
	return Matrix{
		M: ret,
		X: x,
		Y: y,
	}
}

func (m Matrix) String() string {
	sb := strings.Builder{}
	for i := 0; i < m.Y; i++ {
		for j := 0; j < m.X; j++ {
			sb.WriteString(strconv.Itoa(m.M[i][j]))
			sb.WriteString(" ")
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

type Line struct {
	Start, End Point
}
