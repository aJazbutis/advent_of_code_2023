package main

import (
	"advent/utils"
	"fmt"
	"os"
	"slices"
	"strings"
)

func wins(card string) int {
	wins := 0
	a := strings.Split(card, ":")
	b := strings.Split(a[1], "|")
	lucky := strings.Fields(b[0])
	numbers := strings.Fields(b[1])
	for _, number := range numbers {
		if slices.Contains(lucky, number) {
			wins++
		}
	}
	return wins
}

func main() {
	defer utils.Panicked()
	utils.CheckArgs(2, os.Args)
	cards := utils.GetLinesFromFile(os.Args[1])
	won := make([]int, len(cards))
	for i := range won {
		won[i] = 1
	}
	for i := range cards {
		wins := wins(cards[i])
		have := won[i]
		for have > 0 {
			for w := wins; w > 0; w-- {
				won[i+w]++
			}
			have--
		}
	}
	sum := 0
	for i := range won {
		sum += won[i]
	}
	fmt.Println(sum)
}
