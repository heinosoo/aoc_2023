package main

import (
	"log"
	"strconv"
	"strings"
)

func Part1(input <-chan string, output chan string) {
	sum := 0
	for line := range input {
		sum += gameNumberIfPossible(line)
	}

	log.Println(sum)
	output <- strconv.Itoa(sum)
	close(output)
}

func gameNumberIfPossible(line string) int {
	line_splitAtColon := strings.Split(line, ": ")
	gameNumber, _ := strconv.Atoi(strings.Split(line_splitAtColon[0], " ")[1])
	reveals := strings.Split(line_splitAtColon[1], "; ")

	for _, reveal := range reveals {
		if !revealPossible(reveal) {
			return 0
		}
	}
	return gameNumber
}

func revealPossible(reveal string) bool {
	for _, revealCube := range strings.Split(reveal, ", ") {
		revealCubeSplitAtSpace := strings.Split(revealCube, " ")
		cubesRvealed, _ := strconv.Atoi(revealCubeSplitAtSpace[0])
		maxCubes := MAX_CUBES[revealCubeSplitAtSpace[1]]
		if cubesRvealed > maxCubes {
			return false
		}
	}
	return true
}

var MAX_CUBES = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}
