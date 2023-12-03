package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	aoc_utils "github.com/heinosoo/aoc_2023/aoc_utils"
)

func Part1(input <-chan string, output chan string) {
	schematic := aoc_utils.CreateMatrixFromInputChannel(input).Pad(".")
	partNumbers := findPartNumbers(schematic)

	sum := 0
	for _, partNumber := range partNumbers {
		sum += partNumber.value
	}

	log.Println("\n", sum)
	output <- strconv.Itoa(sum)
	close(output)
}

type PartNumber struct {
	x     int
	y     int
	len   int
	value int
}

func findPartNumbers(schematic aoc_utils.Matrix[string]) (numberlocations []PartNumber) {
	numberRegex := regexp.MustCompile(`\d+`)
	for rowNumber, schematicRow := range schematic {
		numberColumns := numberRegex.FindAllStringIndex(strings.Join(schematicRow, ""), -1)
		numbers := numberRegex.FindAllString(strings.Join(schematicRow, ""), -1)

		for i, number := range numbers {
			partNumber := PartNumber{
				x:     numberColumns[i][0],
				y:     rowNumber,
				len:   numberColumns[i][1] - numberColumns[i][0],
				value: aoc_utils.StringToInt(number),
			}
			if partNumber.isValidPartNumber(schematic) {
				numberlocations = append(numberlocations, partNumber)
			}
		}
	}
	return
}

func (numberLocation PartNumber) isValidPartNumber(schematic aoc_utils.Matrix[string]) bool {
	elements := []string{}
	elements = append(elements, schematic[numberLocation.y-1][numberLocation.x-1:numberLocation.x+1+numberLocation.len]...)
	elements = append(elements, schematic[numberLocation.y][numberLocation.x-1:numberLocation.x+1+numberLocation.len]...)
	elements = append(elements, schematic[numberLocation.y+1][numberLocation.x-1:numberLocation.x+1+numberLocation.len]...)

	for _, element := range elements {
		if !strings.Contains(NON_SYMBOLS, element) {
			return true
		}
	}
	return false
}

const NON_SYMBOLS = ".0123456789"
