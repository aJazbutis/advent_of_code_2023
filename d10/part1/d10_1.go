package main

import (
	"advent/utils"
	"fmt"
	"os"
	"strings"
)

func findS(pipes []string) (int, int) {
	for i := range pipes {
		j := strings.Index(pipes[i], "S")
		if j != -1 {
			return i, j
		}
	}
	panic("No S")
}

const (
	west = iota
	north
	east
	south
)

func doALoop(pipes []string) int {
	Si, Sj := findS(pipes)
	i, j := Si, Sj
	dist := 0
	direction := 0
	switch {
	case j != 0 && pipes[i][j-1] != '|' && pipes[i][j-1] != '.':
		j--
		direction = west
	case j != len(pipes[0])-1 && pipes[i][j+1] != '|' && pipes[i][j+1] != '.':
		j++
		direction = east
	case i != 0 && pipes[i-1][j] != '-' && pipes[i-1][j] != '.':
		i--
		direction = north
	case i != len(pipes)-1 && pipes[i+1][j] != '-' && pipes[i+1][j] != '.':
		i++
		direction = south
	default:
		return dist
	}
	dist++
	for pipes[i][j] != 'S' {
		switch pipes[i][j] {
		case '7':
			if direction == north {
				j--
				direction = west
			} else {
				direction = south
				i++
			}
		case 'J':
			if direction == east {
				direction = north
				i--
			} else {
				direction = west
				j--
			}
		case 'L':
			if direction == south {
				direction = east
				j++
			} else {
				direction = north
				i--
			}
		case 'F':
			if direction == north {
				direction = east
				j++
			} else {
				direction = south
				i++
			}
		case '-':
			if direction == west {
				j--
			} else {
				j++
			}
		case '|':
			if direction == north {
				i--
			} else {
				i++
			}
		}
		dist++
	}
	return dist
}

func main() {
	defer utils.Panicked()
	utils.CheckArgs(2, os.Args)
	pipes := utils.GetLinesFromFile(os.Args[1])
	dist := doALoop(pipes)
	fmt.Println((dist / 2) + dist%2)
}
