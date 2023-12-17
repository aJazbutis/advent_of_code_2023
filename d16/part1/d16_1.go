package main

import (
	"advent/utils"
	"fmt"
	"os"
	"slices"
)

const (
	north = iota
	east
	south
	west
)

type Beam struct {
	dir int
	x   int
	y   int
}

type Contraption struct {
	beam Beam
	grid *[][]rune
}

func bouncing(start Contraption, history *[]Beam) [][2]int {
	beam := start.beam
	result := [][2]int{}
	if beam.x < 0 || beam.y < 0 || beam.x > len((*start.grid)[0])-1 || beam.y > len((*start.grid))-1 || slices.Contains(*history, beam) {
		return result
	}
	for true {
		if beam.x < 0 || beam.y < 0 || beam.x > len((*start.grid)[0])-1 || beam.y > len((*start.grid))-1 || slices.Contains(*history, beam) {
			break
		}
		result = append(result, [2]int{beam.y, beam.x})
		*history = append(*history, beam)
		switch (*start.grid)[beam.y][beam.x] {
		case '.':
			switch beam.dir {
			case north:
				beam.y--
			case south:
				beam.y++
			case east:
				beam.x++
			case west:
				beam.x--
			}
		case '\\':
			switch beam.dir {
			case north:
				beam.dir = west
				beam.x--
			case south:
				beam.dir = east
				beam.x++
			case east:
				beam.dir = south
				beam.y++
			case west:
				beam.dir = north
				beam.y--
			}
		case '/':
			switch beam.dir {
			case north:
				beam.dir = east
				beam.x++
			case south:
				beam.dir = west
				beam.x--
			case east:
				beam.dir = north
				beam.y--
			case west:
				beam.dir = south
				beam.y++
			}
		case '|':
			switch beam.dir {
			case north:
				beam.y--
			case south:
				beam.y++
			default:
				split := Contraption{
					beam: Beam{
						dir: north,
						x:   beam.x,
						y:   beam.y - 1,
					},
					grid: start.grid,
				}
				bounce := bouncing(split, history)
				for i := range bounce {
					if !slices.Contains(result, bounce[i]) {
						result = append(result, bounce[i])
					}
				}
				split = Contraption{
					beam: Beam{
						dir: south,
						x:   beam.x,
						y:   beam.y + 1,
					},
					grid: start.grid,
				}
				bounce = bouncing(split, history)
				for i := range bounce {
					if !slices.Contains(result, bounce[i]) {
						result = append(result, bounce[i])
					}
				}
				return result
			}
		case '-':
			switch beam.dir {
			case east:
				beam.x++
			case west:
				beam.x--
			default:
				split := Contraption{
					beam: Beam{
						dir: east,
						x:   beam.x + 1,
						y:   beam.y,
					},
					grid: start.grid,
				}
				bounce := bouncing(split, history)
				for i := range bounce {
					if !slices.Contains(result, bounce[i]) {
						result = append(result, bounce[i])
					}
				}
				split = Contraption{
					beam: Beam{
						dir: west,
						x:   beam.x - 1,
						y:   beam.y,
					},
					grid: start.grid,
				}
				bounce = bouncing(split, history)
				for i := range bounce {
					if !slices.Contains(result, bounce[i]) {
						result = append(result, bounce[i])
					}
				}
				return result
			}
		}
	}
	return result
}

func main() {
	defer utils.Panicked()
	utils.CheckArgs(2, os.Args)
	lines := utils.GetLinesFromFile(os.Args[1])
	grid := [][]rune{}
	for i := range lines {
		grid = append(grid, []rune(lines[i]))
	}
	start := Contraption{
		beam: Beam{
			dir: east,
			x:   0,
			y:   0,
		},
		grid: &grid,
	}
	history := []Beam{}
	energized := bouncing(start, &history)
	fmt.Println(len(energized))
}
