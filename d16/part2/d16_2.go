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
	beam    Beam
	grid    *[][]rune
	history *[]Beam
}

func bouncing(start Contraption) [][2]int {
	beam := start.beam
	result := [][2]int{}
	history := start.history
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
					grid:    start.grid,
					history: history,
				}
				bounce := bouncing(split)
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
					grid:    start.grid,
					history: history,
				}
				bounce = bouncing(split)
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
					grid:    start.grid,
					history: history,
				}
				bounce := bouncing(split)
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
					grid:    start.grid,
					history: history,
				}
				bounce = bouncing(split)
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

func bounce(ins <-chan Contraption, outs chan<- [][2]int) {
	for i := range ins {
		res := bouncing(i)
		outs <- res
	}
}

func main() {
	defer utils.Panicked()
	utils.CheckArgs(2, os.Args)
	lines := utils.GetLinesFromFile(os.Args[1])
	grid := [][]rune{}
	for i := range lines {
		grid = append(grid, []rune(lines[i]))
	}
	buffer := 2*len(grid) + 2*len(grid[0])
	ins := make(chan Contraption, buffer)
	outs := make(chan [][2]int, buffer)
	for i := 0; i < buffer; i++ {
		go bounce(ins, outs)
	}
	for y := 0; y < len(grid); y++ {
		history := []Beam{}
		ins <- Contraption{
			beam: Beam{
				dir: east,
				x:   0,
				y:   y,
			},
			grid:    &grid,
			history: &history,
		}
	}
	for y := 0; y < len(grid); y++ {
		history := []Beam{}
		ins <- Contraption{
			beam: Beam{
				dir: west,
				x:   len(grid[0]) - 1,
				y:   y,
			},
			grid:    &grid,
			history: &history,
		}
	}
	for x := 0; x < len(grid[0]); x++ {
		history := []Beam{}
		ins <- Contraption{
			beam: Beam{
				dir: south,
				x:   x,
				y:   0,
			},
			grid:    &grid,
			history: &history,
		}
	}
	for x := 0; x < len(grid[0]); x++ {
		history := []Beam{}
		ins <- Contraption{
			beam: Beam{
				dir: north,
				x:   x,
				y:   len(grid) - 1,
			},
			grid:    &grid,
			history: &history,
		}
	}
	close(ins)
	max := 0
	for i := 0; i < buffer; i++ {
		energized := <-outs
		if max < len(energized) {
			max = len(energized)
		}
	}
	fmt.Println(max)
}
