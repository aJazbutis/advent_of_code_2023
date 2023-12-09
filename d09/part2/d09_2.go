package main

import (
	"advent/utils"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func allNils(s []int) bool {
	for i := range s {
		if s[i] != 0 {
			return false
		}
	}
	return true
}

func stringsToInts(s []string) []int {
	ret := []int{}
	for i := range s {
		val, _ := strconv.Atoi(s[i])
		ret = append(ret, val)
	}
	return ret
}

func diferenz(s []int) []int {
	ret := []int{}
	for i := 0; i < len(s)-1; i++ {
		ret = append(ret, s[i+1]-s[i])
	}
	return ret
}

func extrapolate(s string) int {
	data := [][]int{stringsToInts(strings.Fields(s))}
	for !allNils(data[len(data)-1]) {
		data = append(data, diferenz(data[len(data)-1]))
	}
	slices.Reverse(data)
	ret := 0
	for i := 0; i < len(data)-1; i++ {
		ret = data[i+1][0] - ret
	}
	return ret
}

func main() {
	defer utils.Panicked()
	utils.CheckArgs(2, os.Args)
	lines := utils.GetLinesFromFile(os.Args[1])
	sum := 0
	for l := range lines {
		sum += extrapolate(lines[l])
	}
	fmt.Println(sum)
}
