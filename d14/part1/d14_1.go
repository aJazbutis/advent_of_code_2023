package main

import (
	"advent/utils"
	"fmt"
	"os"
	"slices"
)

func calculateLoad(s string) int {
	load := 1
	total := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 'O' {
			total += load
		}
		load++
	}
	return total
}

func rocksToNorth(s string) string {
	ret := ""
	for i := 0; i < len(s); i++ {
		for i < len(s) && s[i] == '#' {
			ret += "#"
			i++
		}
		gap := ""
		for i < len(s) && s[i] != '#' {
			gap += string(s[i])
			i++
		}
		i--
		runes := []rune(gap)
		slices.SortFunc(runes, func(a, b rune) int { return int(b) - int(a) })
		ret += string(runes)
	}
	return ret
}

func main() {
	defer utils.Panicked()
	utils.CheckArgs(2, os.Args)
	lines := utils.ColsToRows(utils.GetLinesFromFile(os.Args[1]))
	ret := 0
	for i := range lines {
		ret += calculateLoad(rocksToNorth(lines[i]))
	}
	fmt.Println(ret)
}
