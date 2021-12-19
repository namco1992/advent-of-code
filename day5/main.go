package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

type line struct {
	start, end point
}

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

func scan(reader io.Reader) ([]line, int, int) {
	scanner := bufio.NewScanner(reader)
	var (
		lines      []line
		maxX, maxY int
	)

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " -> ")
		start := strings.Split(s[0], ",")
		end := strings.Split(s[1], ",")
		sp := point{x: mustInt(start[0]), y: mustInt(start[1])}
		ep := point{x: mustInt(end[0]), y: mustInt(end[1])}

		maxX = max(maxX, max(sp.x, ep.x))
		maxY = max(maxY, max(sp.y, ep.y))

		l := line{start: sp, end: ep}
		lines = append(lines, l)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines, maxX + 1, maxY + 1
}

func part1(lines []line, x, y int) int {
	matrix := make([][]int, y)
	for i := 0; i < y; i++ {
		matrix[i] = make([]int, x)
	}

	for _, l := range lines {
		if l.start.x == l.end.x {
			var start point
			if l.start.y >= l.end.y {
				start = l.end
			} else {
				start = l.start
			}
			d := abs(l.start.y - l.end.y)
			for i := 0; i <= d; i++ {
				matrix[start.y+i][l.start.x] += 1
			}
		} else if l.start.y == l.end.y {
			var start point
			if l.start.x >= l.end.x {
				start = l.end
			} else {
				start = l.start
			}
			d := abs(l.start.x - l.end.x)
			for i := 0; i <= d; i++ {
				matrix[l.start.y][start.x+i] += 1
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

func part2(lines []line, x, y int) int {
	matrix := make([][]int, y)
	for i := 0; i < y; i++ {
		matrix[i] = make([]int, x)
	}

	for _, l := range lines {
		switch {
		case l.start.x == l.end.x:
			var start point
			if l.start.y >= l.end.y {
				start = l.end
			} else {
				start = l.start
			}
			d := abs(l.start.y - l.end.y)
			for i := 0; i <= d; i++ {
				matrix[start.y+i][l.start.x] += 1
			}
		case l.start.y == l.end.y:
			var start point
			if l.start.x >= l.end.x {
				start = l.end
			} else {
				start = l.start
			}
			d := abs(l.start.x - l.end.x)
			for i := 0; i <= d; i++ {
				matrix[l.start.y][start.x+i] += 1
			}
		case l.start.x < l.end.x && l.start.y < l.end.y:
			l.start, l.end = l.end, l.start
			fallthrough
		case l.start.x > l.end.x && l.start.y > l.end.y:
			d := l.start.x - l.end.x
			for ; d >= 0; d-- {
				matrix[l.start.y-d][l.start.x-d] += 1
			}
		case l.start.x < l.end.x && l.start.y > l.end.y:
			l.start, l.end = l.end, l.start
			fallthrough
		case l.start.x > l.end.x && l.start.y < l.end.y:
			d := l.start.x - l.end.x
			for ; d >= 0; d-- {
				matrix[l.start.y+d][l.start.x-d] += 1
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
