package main

import (
	"advent/utils"
	"fmt"
	"os"
	"strings"
)

func hashAlgorithm(s string) int {
	currentValue := 0
	for i := range s {
		currentValue += int(s[i])
		currentValue *= 17
		currentValue %= 256
	}
	return currentValue
}
func main() {
	defer utils.Panicked()
	utils.CheckArgs(2, os.Args)
	s := utils.GetLinesFromFile(os.Args[1])
	if len(s) != 1 {
		panic("Supossed to be one string!")
	}
	ss := strings.Split(s[0], ",")
	sum := 0
	for i := range ss {
		sum += hashAlgorithm(ss[i])
	}
	fmt.Println(sum)
}
