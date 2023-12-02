package main

import (
	"advent/utils"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inventory = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func commaToMap(line string) map[string]int {
	colors := strings.Split(line, ",")
	m := make(map[string]int)
	for _, color := range colors {
		color = strings.Trim(color, " ")
		details := strings.Split(color, " ")
		val, err := strconv.Atoi(details[0])
		if err != nil {
			panic(err)
		}
		m[details[1]] = val
	}
	return m
}

func isPossible(sets []map[string]int) bool {
	for _, set := range sets {
		for color, amount := range set {
			if inventory[color] < amount {
				return false
			}
		}
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		panic(errors.New("No args"))
	}
	lines, err := utils.GetLinesFromFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	data := make(map[int][]map[string]int)
	for _, line := range lines {
		forData := strings.Split(line, ":")
		forKey, found := strings.CutPrefix(forData[0], "Game ")
		if !found {
			panic(errors.New("Game <-not found"))
		}
		forVAl := strings.Split(forData[1], ";")
		val := []map[string]int{}
		for _, line := range forVAl {
			val = append(val, commaToMap(line))
		}
		key, err := strconv.Atoi(forKey)
		if err != nil {
			panic(err)
		}
		data[key] = val
	}
	sum := 0
	for idx, game := range data {
		if isPossible(game) {
			sum += idx
		}
	}
	fmt.Println(sum)
}
