package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Part2(input <-chan string, output chan<- string) {
	sum := 0
	for line := range input {
		code := parse2(line)
		// log.Println(line, code)
		sum += code
	}

	log.Println(sum)
	output <- strconv.Itoa(sum)
}

func parse2(line string) int {
	// Hack to deal with overlapping matches
	line = strings.ReplaceAll(line, "one", "oonee")
	line = strings.ReplaceAll(line, "eight", "eeightt")

	allDigitsOrDigitNames := regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`).FindAllString(line, -1)

	firstDigit := DIGIT_NAMES[allDigitsOrDigitNames[0]]
	lastDigit := DIGIT_NAMES[allDigitsOrDigitNames[len(allDigitsOrDigitNames)-1]]
	code, _ := strconv.Atoi(firstDigit + lastDigit)
	return code
}

var DIGIT_NAMES = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
}
