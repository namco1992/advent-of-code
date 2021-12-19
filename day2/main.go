package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/namco1992/aoc/util"
)

const (
	forward = "forward"
	down    = "down"
	up      = "up"
)

func part1(input []string) int {
	var x, y int

	for _, s := range input {
		split := strings.Split(s, " ")
		if len(split) != 2 {
			log.Fatalf("invalid input: %s", s)
		}

		cmd, move := split[0], split[1]
		m, err := strconv.Atoi(move)
		if err != nil {
			log.Fatalf("invalid input: %s", move)
		}

		switch cmd {
		case forward:
			x += m
		case up:
			y -= m
		case down:
			y += m
		}
	}

	return x * y
}

func part2(input []string) int {
	var x, y, aim int

	for _, s := range input {
		split := strings.Split(s, " ")
		if len(split) != 2 {
			log.Fatalf("invalid input: %s", s)
		}
		cmd, move := split[0], split[1]
		m, err := strconv.Atoi(move)
		if err != nil {
			log.Fatalf("invalid input: %s", move)
		}

		switch cmd {
		case forward:
			x += m
			y += aim * m
		case up:
			aim -= m
		case down:
			aim += m
		}
	}

	return x * y
}

func main() {
	input := util.ScanFileToStrings("day2/input")

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
