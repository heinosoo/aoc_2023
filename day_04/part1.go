package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/heinosoo/aoc_2023/aoc_utils"
)

func Part1(input <-chan string, output chan<- string) {
	sum := 0
	for line := range input {
		winningNumbers, numbers := parseLine(line)
		points := calculatePointsOfCard(winningNumbers, numbers)
		sum += points
	}

	log.Println(sum)
	output <- strconv.Itoa(sum)
}

func parseLine(line string) (winningNumbers, numbers []int) {
	numberParts := strings.Split(strings.Split(line, ":")[1], "|")
	numberRegex := regexp.MustCompile(`\b\d+\b`)
	winningNumberStrings := numberRegex.FindAllString(numberParts[0], -1)
	numberStrings := numberRegex.FindAllString(numberParts[1], -1)

	for _, number := range winningNumberStrings {
		winningNumbers = append(winningNumbers, aoc_utils.StringToInt(number))
	}

	for _, number := range numberStrings {
		numbers = append(numbers, aoc_utils.StringToInt(number))
	}

	return
}

func calculatePointsOfCard(winningNumbers, numbers []int) (points int) {
	for _, number := range numbers {
		if aoc_utils.Contains(winningNumbers, number) {
			points = max(1, points*2)
		}
	}
	return
}
