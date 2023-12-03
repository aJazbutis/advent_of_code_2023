package main

import (
	"advent/utils"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) != 2 {
		panic("No arg")
	}
	lines, err := utils.GetLinesFromFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	sum := 0
	for i, line := range lines {
		for idx := strings.IndexRune(line, '*'); idx > 0; idx = strings.IndexRune(line, '*') {
			line = strings.Replace(line, "*", ".", 1)
			gear := 1
			adjacent := 0
			if idx != 0 {
				if unicode.IsDigit(rune(line[idx-1])) {
					adjacent++
					if adjacent > 2 {
						continue
					}
					gear *= utils.ExtractInt(line, idx-1)
				}
			}
			if i != 0 {
				if unicode.IsDigit(rune(lines[i-1][idx])) {
					adjacent++
					if adjacent > 2 {
						continue
					}
					gear *= utils.ExtractInt(lines[i-1], idx)
				} else {
					if idx != 0 {
						if unicode.IsDigit(rune(lines[i-1][idx-1])) {
							adjacent++
							if adjacent > 2 {
								continue
							}
							gear *= utils.ExtractInt(lines[i-1], idx-1)
						}
					}
					if idx != len(line)-1 {

						if unicode.IsDigit(rune(lines[i-1][idx+1])) {
							adjacent++
							if adjacent > 2 {
								continue
							}
							gear *= utils.ExtractInt(lines[i-1], idx+1)
						}
					}
				}
			}
			if idx != len(line)-1 {
				if unicode.IsDigit(rune(line[idx+1])) {
					adjacent++
					if adjacent > 2 {
						continue
					}
					gear *= utils.ExtractInt(line, idx+1)
				}
			}
			if i != len(lines)-1 {
				if unicode.IsDigit(rune(lines[i+1][idx])) {
					adjacent++
					if adjacent > 2 {
						continue
					}
					gear *= utils.ExtractInt(lines[i+1], idx)
				} else {
					if idx != 0 {
						if unicode.IsDigit(rune(lines[i+1][idx-1])) {
							adjacent++
							if adjacent > 2 {
								continue
							}
							gear *= utils.ExtractInt(lines[i+1], idx-1)
						}
					}
					if idx != len(line)-1 {
						if unicode.IsDigit(rune(lines[i+1][idx+1])) {
							adjacent++
							if adjacent > 2 {
								continue
							}
							gear *= utils.ExtractInt(lines[i+1], idx+1)
						}
					}
				}
			}
			if adjacent == 2 {
				sum += gear
			}
		}
	}
	fmt.Println(sum)
}
