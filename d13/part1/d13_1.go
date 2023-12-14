package main

import (
	"advent/utils"
	"fmt"
	"os"
	"strings"
)

func checkReflection(idx, multiply int, lines []string) (int, bool) {
	c := 0
	up := idx
	down := idx + 1
	for up >= 0 && down < len(lines) {
		if strings.Compare(lines[up], lines[down]) != 0 {
			return 0, false
		}
		c++
		up--
		down++
	}
	for up >= 0 {
		up--
		c++
	}
	return c * multiply, true
}

func rowChecker(lines []string, isRow bool) int {
	for i := 0; i < len(lines)-1; i++ {
		if strings.Compare(lines[i], lines[i+1]) == 0 {
			if isRow {
				ret, ok := checkReflection(i, 100, lines)
				if ok {
					return ret
				}
			} else {
				ret, ok := checkReflection(i, 1, lines)
				if ok {
					return ret
				}
			}

		}
	}
	return 0
}

func colsToRows(lines []string) []string {
	ret := []string{}
	for i := range lines[0] {
		s := ""
		for j := range lines {
			s += string(lines[j][i])
		}
		ret = append(ret, s)
	}
	return ret
}

func analysePatern(lines []string) int {
	ret := rowChecker(lines, true)
	if ret == 0 {
		ret = rowChecker(colsToRows(lines), false)
	}
	return ret
}

func loopThroughLines(lines []string) int {
	ret := 0
	patern := []string{}
	for i := range lines {
		switch lines[i] {
		case "":
			ret += analysePatern(patern)
			patern = []string{}
		default:
			patern = append(patern, lines[i])
			if i == len(lines)-1 {
				ret += analysePatern(patern)
			}
		}
	}
	return ret
}

func main() {
	defer utils.Panicked()
	utils.CheckArgs(2, os.Args)
	fmt.Println(loopThroughLines(utils.GetLinesFromFile(os.Args[1])))
}
