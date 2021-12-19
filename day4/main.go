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

var numPos = [5][2]int{{0, 2}, {3, 5}, {6, 8}, {9, 11}, {12, 14}}

func scan(reader io.Reader) (input []int, boards [][5][5]int) {
	scanner := bufio.NewScanner(reader)

	scanner.Scan()
	s := strings.Split(scanner.Text(), ",")
	input = make([]int, len(s))
	for i, v := range s {
		number, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		input[i] = number
	}

	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			board := [5][5]int{}
			for i := 0; i < 5; i++ {
				scanner.Scan()
				row := scanner.Text()
				for j := 0; j < 5; j++ {
					num := row[numPos[j][0]:numPos[j][1]]
					var err error
					board[i][j], err = strconv.Atoi(strings.TrimSpace(num))
					if err != nil {
						log.Fatal(err)
					}
				}
			}
			boards = append(boards, board)
		}
	}
	return input, boards
}

func calcScore(board, scoreBoard [5][5]int, multipler int) int {
	var sum int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if scoreBoard[i][j] == 0 {
				sum += board[i][j]
			}
		}
	}
	return sum * multipler
}

func isWinning(scoreBoard [5][5]int) bool {
	for i := 0; i < 5; i++ {
		if scoreBoard[i][0] == 1 && scoreBoard[i][1] == 1 && scoreBoard[i][2] == 1 && scoreBoard[i][3] == 1 && scoreBoard[i][4] == 1 {
			return true
		}
		if scoreBoard[0][i] == 1 && scoreBoard[1][i] == 1 && scoreBoard[2][i] == 1 && scoreBoard[3][i] == 1 && scoreBoard[4][i] == 1 {
			return true
		}
	}
	return false
}

func part1(input []int, boards [][5][5]int) int {
	scoreBoards := make([][5][5]int, len(boards))
	for i, board := range boards {
		for j := 0; j < 5; j++ {
			for m := 0; m < 5; m++ {
				for n := 0; n < 5; n++ {
					if input[j] == board[m][n] {
						scoreBoards[i][m][n] = 1
					}
				}
			}
		}
		if isWinning(scoreBoards[i]) {
			return calcScore(boards[i], scoreBoards[i], input[4]) // it has to be the 5th one let it win
		}
	}

	for i := 5; i < len(input); i++ {
		for j, board := range boards {
			var found bool
			for m := 0; m < 5; m++ {
				for n := 0; n < 5; n++ {
					if input[i] == board[m][n] {
						found = true
						scoreBoards[j][m][n] = 1
					}
				}
			}
			if found && isWinning(scoreBoards[j]) {
				fmt.Printf("winner: %d, %v\n%v\n", j, scoreBoards[j], boards[j])
				return calcScore(boards[j], scoreBoards[j], input[i])
			}
		}
	}

	return 0
}

func part2(input []int, boards [][5][5]int) int {
	scoreBoards := make([][5][5]int, len(boards))
	var lastWinnerScore int

	// just fill, no check
	for i, board := range boards {
		for j := 0; j < 5; j++ {
			for m := 0; m < 5; m++ {
				for n := 0; n < 5; n++ {
					if input[j] == board[m][n] {
						scoreBoards[i][m][n] = 1
					}
				}
			}
		}
	}

	for i := 5; i < len(input); i++ {
		for j, board := range boards {
			for m := 0; m < 5; m++ {
				for n := 0; n < 5; n++ {
					if input[i] == board[m][n] {
						scoreBoards[j][m][n] = 1
					}
				}
			}
		}

		delIndex := make(map[int]struct{})
		for j, board := range boards {
			if isWinning(scoreBoards[j]) {
				lastWinnerScore = calcScore(board, scoreBoards[j], input[i])
				if len(boards) == 1 {
					return lastWinnerScore
				}

				delIndex[j] = struct{}{}
			}
		}

		var b, s [][5][5]int
		for j, board := range boards {
			if _, ok := delIndex[j]; !ok {
				b = append(b, board)
			}
		}

		for j, scoreBoard := range scoreBoards {
			if _, ok := delIndex[j]; !ok {
				s = append(s, scoreBoard)
			}
		}

		scoreBoards = s
		boards = b
	}

	return lastWinnerScore
}

func main() {
	file, err := os.Open("day4/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	input, boards := scan(file)
	// fmt.Println(part1(input, boards))
	fmt.Println(part2(input, boards))
}
