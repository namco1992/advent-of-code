package util

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

func ScanFileToStrings(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return ScanToStrings(file)
}

func ScanFileToInts(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return ScanToInts(file)
}

func ScanToStrings(reader io.Reader) []string {
	scanner := bufio.NewScanner(reader)
	var ret []string

	for scanner.Scan() {
		s := scanner.Text()
		ret = append(ret, s)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return ret
}

func ScanToInts(reader io.Reader) []int {
	scanner := bufio.NewScanner(reader)
	var ret []int

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("invalid input: %s", scanner.Text())
		}
		ret = append(ret, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return ret
}

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
