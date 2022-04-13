package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/namco1992/aoc/util"
)

func part1(input []string) int {
	var (
		foldIndex, x, y int
	)
	pointMap := make(map[util.Point]struct{})
	for i, s := range input {
		if s == "" {
			foldIndex = i + 1
			break
		}
		ss := strings.Split(s, ",")
		p := util.Point{
			X: util.MustAtoi(ss[0]),
			Y: util.MustAtoi(ss[1]),
		}
		pointMap[p] = struct{}{}
		x = util.Max(x, p.X)
		y = util.Max(y, p.Y)
	}

	fmt.Printf("x=%v, y=%v\n", x, y)

	for i := foldIndex; i < len(input); i++ {
		fold := strings.Split(strings.TrimPrefix(input[i], "fold along "), "=")
		value := util.MustAtoi(fold[1])
		if fold[0] == "x" {
			for p := range pointMap {
				if p.X > value {
					foldedPoint := util.Point{X: value - (p.X - value), Y: p.Y}
					pointMap[foldedPoint] = struct{}{}
					delete(pointMap, p)
				}
			}
		} else {
			for p := range pointMap {
				if p.Y > value {
					foldedPoint := util.Point{X: p.X, Y: value - (p.Y - value)}
					pointMap[foldedPoint] = struct{}{}
					delete(pointMap, p)
				}
			}
		}
	}

	b := strings.Builder{}

	maxX, maxY := math.MinInt32, math.MinInt32
	for pos := range pointMap {
		maxX = util.Max(maxX, pos.X)
		maxY = util.Max(maxY, pos.Y)
	}
	fmt.Println(maxX, maxY)
	for i := 0; i <= maxY; i++ {
		for j := 0; j <= maxX; j++ {
			if _, ok := pointMap[util.Point{X: j, Y: i}]; ok {
				b.WriteRune('#')
			} else {
				b.WriteRune('.')
			}
		}
		b.WriteRune('\n')
	}
	fmt.Println(b.String())
	return 1
}

func main() {
	input := util.ScanFileToStrings("day13/input")
	fmt.Println(part1(input))
}
