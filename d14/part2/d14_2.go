package main

import (
	"advent/utils"
	"fmt"
	"os"
	"slices"
)

func calculateLoad(s []string) int {
	total := 0
	ss := utils.ColsToRows(s)
	for j := range ss {
		load := 1
		for i := len(s) - 1; i >= 0; i-- {
			if ss[j][i] == 'O' {
				total += load
			}
			load++
		}
	}
	return total
}

func rocksToLeft(s string) string {
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

func north(s []string) []string {
	s = utils.ColsToRows(s)
	for i := range s {
		s[i] = rocksToLeft(s[i])
	}
	return utils.RowsToCols(s)
}

func west(s []string) []string {
	for i := range s {
		s[i] = rocksToLeft(s[i])
	}
	return s
}

func rocksToRight(s string) string {
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
		slices.SortFunc(runes, func(a, b rune) int { return int(a) - int(b) })
		ret += string(runes)
	}
	return ret
}

func east(s []string) []string {
	for i := range s {
		s[i] = rocksToRight(s[i])
	}
	return (s)
}

func south(s []string) []string {
	s = utils.ColsToRows(s)
	for i := range s {
		s[i] = rocksToRight(s[i])
	}
	return utils.RowsToCols(s)
}

const cycle = 34

func main() {
	defer utils.Panicked()
	utils.CheckArgs(2, os.Args)
	platform := utils.GetLinesFromFile(os.Args[1])
	for i := 0; i < 20*cycle+1000000000%cycle; i++ {
		platform = east(south(west(north(platform))))
	}
	fmt.Println(calculateLoad(platform))
}
