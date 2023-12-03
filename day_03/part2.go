package main

import (
	"log"
	"strconv"

	aoc_utils "github.com/heinosoo/aoc_2023/aoc_utils"
)

func Part2(input <-chan string, output chan<- string) {
	schematic := aoc_utils.CreateMatrixFromInputChannel(input).Pad(".")
	partNumbers := findPartNumbers(schematic)
	potentialGears := findPotentialGears(schematic, partNumbers)

	sum := 0
	for _, partNumbers := range potentialGears {
		if len(partNumbers) == 2 {
			sum += partNumbers[0].value * partNumbers[1].value
		}
	}

	log.Println(sum)
	output <- strconv.Itoa(sum)
}

type Gears map[[2]int][]PartNumber

func findPotentialGears(schematic aoc_utils.Matrix[string], partNumbers []PartNumber) (gears Gears) {
	gears = make(Gears)
	for _, partNumber := range partNumbers {
		gearLocs := findGearLocationsAroundPartNumber(schematic, partNumber)
		for _, gearLoc := range gearLocs {
			gears[gearLoc] = append(gears[gearLoc], partNumber)
		}
	}
	return
}

func findGearLocationsAroundPartNumber(schematic aoc_utils.Matrix[string], partNumber PartNumber) (gearLocs [][2]int) {
	gearSymbol := "*"

	for x := partNumber.x - 1; x <= partNumber.x+partNumber.len; x++ {
		for y := partNumber.y - 1; y <= partNumber.y+1; y++ {
			if schematic[y][x] == gearSymbol {
				gearLocs = append(gearLocs, [2]int{x, y})
			}
		}
	}

	return
}
