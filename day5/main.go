package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/namco1992/aoc/util"
)

func scan(reader io.Reader) ([]util.Line, int, int) {
	scanner := bufio.NewScanner(reader)
	var (
		lines      []util.Line
		maxX, maxY int
	)

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " -> ")
		start := strings.Split(s[0], ",")
		end := strings.Split(s[1], ",")
		sp := util.Point{X: util.MustAtoi(start[0]), Y: util.MustAtoi(start[1])}
		ep := util.Point{X: util.MustAtoi(end[0]), Y: util.MustAtoi(end[1])}

		maxX = util.Max(maxX, util.Max(sp.X, ep.X))
		maxY = util.Max(maxY, util.Max(sp.Y, ep.Y))

		l := util.Line{Start: sp, End: ep}
		lines = append(lines, l)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines, maxX + 1, maxY + 1
}

func part1(lines []util.Line, x, y int) int {
	matrix := make([][]int, y)
	for i := 0; i < y; i++ {
		matrix[i] = make([]int, x)
	}

	for _, l := range lines {
		if l.Start.X == l.End.X {
			var start util.Point
			if l.Start.Y >= l.End.Y {
				start = l.End
			} else {
				start = l.Start
			}
			d := util.Abs(l.Start.Y - l.End.Y)
			for i := 0; i <= d; i++ {
				matrix[start.Y+i][l.Start.X] += 1
			}
		} else if l.Start.Y == l.End.Y {
			var start util.Point
			if l.Start.X >= l.End.X {
				start = l.End
			} else {
				start = l.Start
			}
			d := util.Abs(l.Start.X - l.End.X)
			for i := 0; i <= d; i++ {
				matrix[l.Start.Y][start.X+i] += 1
			}
		}
	}

	var cnt int
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if matrix[i][j] >= 2 {
				cnt++
			}
		}
	}
	fmt.Println(matrix)
	return cnt
}

func part2(lines []util.Line, x, y int) int {
	matrix := make([][]int, y)
	for i := 0; i < y; i++ {
		matrix[i] = make([]int, x)
	}

	for _, l := range lines {
		switch {
		case l.Start.X == l.End.X:
			var start util.Point
			if l.Start.Y >= l.End.Y {
				start = l.End
			} else {
				start = l.Start
			}
			d := util.Abs(l.Start.Y - l.End.Y)
			for i := 0; i <= d; i++ {
				matrix[start.Y+i][l.Start.X] += 1
			}
		case l.Start.Y == l.End.Y:
			var start util.Point
			if l.Start.X >= l.End.X {
				start = l.End
			} else {
				start = l.Start
			}
			d := util.Abs(l.Start.X - l.End.X)
			for i := 0; i <= d; i++ {
				matrix[l.Start.Y][start.X+i] += 1
			}
		case l.Start.X < l.End.X && l.Start.Y < l.End.Y:
			l.Start, l.End = l.End, l.Start
			fallthrough
		case l.Start.X > l.End.X && l.Start.Y > l.End.Y:
			d := l.Start.X - l.End.X
			for ; d >= 0; d-- {
				matrix[l.Start.Y-d][l.Start.X-d] += 1
			}
		case l.Start.X < l.End.X && l.Start.Y > l.End.Y:
			l.Start, l.End = l.End, l.Start
			fallthrough
		case l.Start.X > l.End.X && l.Start.Y < l.End.Y:
			d := l.Start.X - l.End.X
			for ; d >= 0; d-- {
				matrix[l.Start.Y+d][l.Start.X-d] += 1
			}
		}
	}

	var cnt int
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if matrix[i][j] >= 2 {
				cnt++
			}
		}
	}
	fmt.Println(matrix)
	return cnt
}

func main() {
	file, err := os.Open("day5/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines, x, y := scan(file)
	fmt.Println(part2(lines, x, y))
}
