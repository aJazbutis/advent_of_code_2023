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

func diferenz(s []int, firsts *[]int) []int {
	ret := []int{}
	for i := 0; i < len(s)-1; i++ {
		ret = append(ret, s[i+1]-s[i])
	}
	*firsts = append(*firsts, ret[0])
	return ret
}

func extrapolate(s string) int {
	data := stringsToInts(strings.Fields(s))
	firsts := []int{data[0]}
	for !allNils(data) {
		data = diferenz(data, &firsts)
	}
	slices.Reverse(firsts)
	ret := 0
	for i := 0; i < len(firsts)-1; i++ {
		ret = firsts[i+1] - ret
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
