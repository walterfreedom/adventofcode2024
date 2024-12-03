package main

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	part2()
}

func part1() {
	file, err := os.Open("input")
	if err != nil {
		print("oh no")
	}

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		text := scanner.Text()
		pattern := `mul\(\d+,\d+\)`

		re := regexp.MustCompile(pattern)

		matches := re.FindAllString(text, -1)

		for _, match := range matches {
			sum += mulxecute(match)
		}
	}
	print(sum)
}

func part2() {
	file, err := os.Open("input")

	if err != nil {
		print("oh no")
	}

	sum := 0

	var goodchunks []string

	textbyte, err := io.ReadAll(file)
	textstring := string(textbyte)
	dontremoved := strings.Split(textstring, "don't()")
	goodchunks = append(goodchunks, dontremoved[0])
	for _, chunk := range dontremoved {
		println(chunk)
		subchunk := strings.SplitN(chunk, "do()", 2)
		if len(subchunk) > 1 {
			goodchunks = append(goodchunks, subchunk[1])
		}
	}

	for _, goodchunk := range goodchunks {
		println(goodchunk)
		pattern := `mul\(\d+,\d+\)`

		re := regexp.MustCompile(pattern)

		matches := re.FindAllString(goodchunk, -1)

		for _, match := range matches {
			sum += mulxecute(match)
		}
	}

	print(sum)
}

func mulxecute(match string) int {
	//we get nummers
	trimmed := strings.TrimPrefix(match, "mul(")
	trimmed = strings.TrimSuffix(trimmed, ")")
	numbers := strings.Split(trimmed, ",")

	a, err1 := strconv.Atoi(numbers[0])
	b, err2 := strconv.Atoi(numbers[1])
	if err1 != nil || err2 != nil {
		print("I like ferrets")
		return 0
	}

	//complicated math
	return a * b
}
