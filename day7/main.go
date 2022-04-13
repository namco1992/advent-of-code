package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func scanToIntArray(reader io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()

	ss := strings.Split(scanner.Text(), ",")
	input := make([]int, len(ss))
	for i, s := range ss {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		input[i] = n
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return input, nil
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

func part1(input []int) int {
	sort.Ints(input)
	median := input[len(input)/2]
	var cost int
	for _, n := range input {
		cost += abs(n - median)
	}
	return cost
}

func part2(input []int) float64 {
	var avg float64
	for _, n := range input {
		avg += float64(n)
	}
	avg = avg / float64(len(input))
	ceiling := math.Ceil(avg)
	// floor := math.go.Floor(avg)
	fmt.Println(ceiling)
	var cost int
	for _, n := range input {
		m := abs(int(ceiling) - n)
		cost += (m * (m + 1)) / 2
	}
	fmt.Println(cost)
	return avg
}

// 95167367
// 95167302

func main() {
	file, err := os.Open("day7/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input, err := scanToIntArray(file)
	fmt.Println(part2(input))
}
