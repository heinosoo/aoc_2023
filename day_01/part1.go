package main

import (
	"strconv"
)

func Part1(input <-chan string, output chan<- string) {
	sum := 0
	for line := range input {
		code := parse(line)
		// log.Println(code)
		sum += code
	}

	output <- strconv.Itoa(sum)
}

func parse(line string) int {
	code, _ := strconv.Atoi(firstDigit(line) + lastDigit(line))
	return code
}

func firstDigit(line string) string {
	for i := 0; i < len(line); i++ {
		if line[i] >= '0' && line[i] <= '9' {
			return string(line[i])
		}
	}
	// panic("no first digit found for line: " + line)
	return "0"
}
func lastDigit(line string) string {
	for i := len(line) - 1; i >= 0; i-- {
		if line[i] >= '0' && line[i] <= '9' {
			return string(line[i])
		}
	}
	// panic("no last digit found for line: " + line)
	return "0"
}
