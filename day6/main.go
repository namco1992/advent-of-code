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

func part1(input []int) int {
	for i := 0; i < 80; i++ {
		var newFishCnt int
		for i, _ := range input {
			input[i] -= 1
			if input[i] < 0 {
				newFishCnt++
				input[i] = 6
			}
		}
		for j := 0; j < newFishCnt; j++ {
			input = append(input, 8)
		}
	}
	return len(input)
}

func part2(input []int) int {
	fishes := make([]int, 9)
	for _, n := range input {
		fishes[n]++
	}

	for i := 0; i < 256; i++ {
		// Keep count of fishes ready to reproduce
		last0 := fishes[0]

		// Decrease reproduce timer
		for k := 0; k <= 7; k++ {
			fishes[k] = fishes[k+1]
		}
		// Add new fishes
		fishes[8] = last0
		// Reset fishes that have reproduced
		fishes[6] += last0
	}

	totalFishes := 0
	for _, v := range fishes {
		totalFishes += v
	}

	return totalFishes
}

func main() {
	file, err := os.Open("day6/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input, err := scanToIntArray(file)
	fmt.Println(part2(input))
}
