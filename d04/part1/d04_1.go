package main

import (
	"advent/utils"
	"fmt"
	"os"
	"slices"
	"strings"
)

func getPoints(card string) int {
	points := 0
	a := strings.Split(card, ":")
	b := strings.Split(a[1], "|")
	lucky := strings.Fields(b[0])
	numbers := strings.Fields(b[1])
	for _, number := range numbers {
		if slices.Contains(lucky, number) {
			if points == 0 {
				points++
			} else {
				points *= 2
			}
		}
	}
	return points
}

func main() {
	defer utils.Panicked()
	utils.CheckArgs(2, os.Args)
	cards := utils.GetLinesFromFile(os.Args[1])
	sum := 0
	for _, card := range cards {
		sum += getPoints(card)
	}
	fmt.Println(sum)
}
