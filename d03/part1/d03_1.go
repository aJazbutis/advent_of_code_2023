package main

import (
	"advent/utils"
	"os"
	"fmt"
	"strings"
	"unicode"
)

func getMarks(lines []string) string {
	m := ""
	for _, line := range lines {
		for _, c := range line {
			if strings.IndexRune(m, c) == -1 && c != '.' && !unicode.IsDigit(c) {
				m += string(c)
			}
		}
	}
	return m
}

func main() {
	if len(os.Args) != 2 {
		panic("No arg")
	}
	lines, err := utils.GetLinesFromFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	marks := getMarks(lines)
	sum := 0
	for i, line := range lines {

		for idx := 0; idx < len(line); idx++ {
			if strings.IndexByte(marks, line[idx]) != -1 {
				if idx != 0 {
					if unicode.IsDigit(rune(line[idx-1])) {
						sum += utils.ExtractInt(line, idx-1)
					}
				}
				if i != 0 {
					if unicode.IsDigit(rune(lines[i-1][idx])) {
						sum += utils.ExtractInt(lines[i-1], idx)
					} else {
						if idx != 0 {
							if unicode.IsDigit(rune(lines[i-1][idx-1])) {
								sum += utils.ExtractInt(lines[i-1], idx-1)
							}
						}
						if idx != len(line)-1 {
							if unicode.IsDigit(rune(lines[i-1][idx+1])) {
								sum += utils.ExtractInt(lines[i-1], idx+1)
							}
						}
					}
				}
				if idx != len(line)-1 {
					if unicode.IsDigit(rune(line[idx+1])) {
						sum += utils.ExtractInt(line, idx+1)
					}
				}
				if i != len(lines)-1 {
					if unicode.IsDigit(rune(lines[i+1][idx])) {
						sum += utils.ExtractInt(lines[i+1], idx)
					} else {
						if idx != 0 {
							if unicode.IsDigit(rune(lines[i+1][idx-1])) {
								sum += utils.ExtractInt(lines[i+1], idx-1)
							}
						}
						if idx != len(line)-1 {
							if unicode.IsDigit(rune(lines[i+1][idx+1])) {
								sum += utils.ExtractInt(lines[i+1], idx+1)
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(sum)
}
