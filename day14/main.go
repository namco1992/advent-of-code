package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/namco1992/aoc/util"
)

func part1(input []string) int {
	pairInsertionMap := make(map[string]string)
	template := input[0]
	lastChar := string(template[len(template)-1])
	for i := 2; i < len(input); i++ {
		s := strings.Split(input[i], " -> ")
		if len(s) != 2 {
			log.Fatal("failed to parse the pair insertion rule")
		}
		pairInsertionMap[s[0]] = s[1]
	}

	pairFreq := make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		pairFreq[template[i:i+2]]++
	}

	for j := 0; j < 40; j++ {
		output := make(map[string]int)
		for pair, freq := range pairFreq {
			element, ok := pairInsertionMap[pair]
			if ok {
				output[pair[0:1]+element] += freq
				output[element+pair[1:2]] += freq
			}
		}
		pairFreq = output
	}
	fmt.Println(pairFreq)

	elementFreqs := make(map[string]int)
	for pair, freq := range pairFreq {
		elementFreqs[pair[0:1]] += freq
	}
	elementFreqs[lastChar]++
	counts := make([]int, len(elementFreqs))
	var i int
	for _, f := range elementFreqs {
		counts[i] = f
		i++
	}
	sort.Ints(counts)
	fmt.Println(counts)
	return 0
}

func main() {
	input := util.ScanFileToStrings("day14/input")
	fmt.Println(part1(input))
}
