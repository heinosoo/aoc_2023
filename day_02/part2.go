package main

import (
	"log"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

func Part2(input <-chan string, output chan<- string) {
	sum := 0
	for line := range input {
		sum += minimumCubePower(line)
	}

	log.Println(sum)
	output <- strconv.Itoa(sum)
}

type CubeSet map[string]int

func (cubeSet CubeSet) power() int {
	power := 1
	for _, cubes := range maps.Values(cubeSet) {
		power *= cubes
	}
	return power
}

func minimumCubePower(line string) int {
	line_splitAtColon := strings.Split(line, ": ")
	reveals := strings.Split(line_splitAtColon[1], "; ")
	cubeSet := make(CubeSet)

	for _, reveal := range reveals {
		cubeSet.addReveal(reveal)
	}

	return cubeSet.power()
}

func (cubeSet CubeSet) addReveal(reveal string) {
	for _, revealCube := range strings.Split(reveal, ", ") {
		revealCubeSplitAtSpace := strings.Split(revealCube, " ")
		cubesRevealed, _ := strconv.Atoi(revealCubeSplitAtSpace[0])
		cubeColor := revealCubeSplitAtSpace[1]
		cubeSet[cubeColor] = max(cubeSet[cubeColor], cubesRevealed)
	}
}
