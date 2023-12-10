package main

import (
	"log"
	"strconv"

	"github.com/heinosoo/aoc_2023/aoc_utils"
)

func Part2(input <-chan string, output chan<- string) {
	cards := Cards{}
	for line := range input {
		cards = append(cards, newCardFromLine(line))
	}

	more := true
	for more {
		more = cards.addCardCopies()
	}

	total := 0
	for _, card := range cards {
		total += card.total
	}

	log.Println(total)
	output <- strconv.Itoa(total)
}

type Cards []*Card

type Card struct {
	total    int
	unparsed int
	matches  int
}

func newCardFromLine(line string) *Card {
	winningNumbers, numbers := parseLine(line)
	matches := 0
	for _, number := range numbers {
		if aoc_utils.Contains(winningNumbers, number) {
			matches++
		}
	}
	return &Card{0, 1, matches}
}

func (cards Cards) addCardCopies() (more bool) {
	for i, card := range cards {
		if card.unparsed > 0 {
			for j := i + 1; j < i+card.matches+1; j++ {
				cards[j].unparsed += card.unparsed
			}
			card.total += card.unparsed
			card.unparsed = 0
			more = true
		}
	}
	return
}
