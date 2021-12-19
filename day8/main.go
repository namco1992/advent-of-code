package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/namco1992/aoc/util"
)

var lenNumMap = map[int]int{2: 1, 3: 7, 4: 4, 7: 8}

func part1(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)
	var cnt int

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " | ")
		output := strings.Split(s[1], " ")
		for _, o := range output {
			l := len(o)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				cnt++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return -1
	}

	return cnt
}

func part2(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)
	var cnt int

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " | ")
		input := strings.Split(s[0], " ")
		output := strings.Split(s[1], " ")
		baseMap := make(map[int]util.CharSet)
		for _, i := range input {
			switch l := len(i); l {
			case 2, 3, 4, 7:
				baseMap[lenNumMap[l]] = util.NewSet(i)
			}
		}

		for _, o := range output {
			switch l := len(o); l {
			case 2, 3, 4, 7:
				baseMap[lenNumMap[l]] = util.NewSet(o)
			}
		}

		var ret string

		for _, o := range output {
			switch l := len(o); l {
			case 2, 3, 4, 7:
				ret += strconv.Itoa(lenNumMap[l])
			case 5:
				set := util.Intersection(baseMap[1], util.NewSet(o))
				if len(set) == 2 {
					ret += "3"
				} else {
					set := util.Intersection(baseMap[4], util.NewSet(o))
					if len(set) == 2 {
						ret += "2"
					} else {
						ret += "5"
					}
				}
			case 6:
				set := util.Intersection(baseMap[1], util.NewSet(o))
				if len(set) == 1 {
					ret += "6"
				} else {
					set := util.Intersection(baseMap[4], util.NewSet(o))
					if len(set) == 3 {
						ret += "0"
					} else {
						ret += "9"
					}
				}
			}
		}

		n, err := strconv.Atoi(ret)
		if err != nil {
			log.Fatal(err)
		}
		cnt += n
	}
	return cnt
}

func main() {
	file, err := os.Open("day8/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println(part2(file))
}
