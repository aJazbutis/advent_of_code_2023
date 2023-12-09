package main

import (
	"advent/utils"
	"fmt"
	"os"
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

func diferenz(s []int, res *[]int) []int {
	ret := []int{}
	for i := 0; i < len(s)-1; i++ {
		ret = append(ret, s[i+1]-s[i])
	}
	*res = append(*res, ret[len(ret)-1])
	return ret
}

func extrapolate(s string) int {
	data := stringsToInts(strings.Fields(s))
	prelasts := []int{data[len(data)-1]}
	for !allNils(data) {
		data = diferenz(data, &prelasts)
	}
	sum := 0
	for _, val := range prelasts {
		sum += val
	}
	return sum
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
