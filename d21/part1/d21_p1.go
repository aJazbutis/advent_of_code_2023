package main

import (
	"fmt"
	"os"
	"strings"
	"utils"
)

func findS(lines []string) (int, int) {
	for y := range lines {
		x := strings.Index(lines[y], "S")
		if x >= 0 {
			return y, x
		}
	}
	panic("No S")
}

type point struct {
	Y int
	X int
}

func bfs(grid [][]byte, start point, steps int, dir [][]int) {
	q := []point{start}
	for len(q) > 0 && steps > 0 {
		choices := len(q)
		for choices > 0 {
			step := q[0]
			q = q[1:]
			grid[step.Y][step.X] = 'O'
			for _, d := range dir {
				x := step.X + d[1]
				y := step.Y + d[0]
				if x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid) && (grid[y][x] == '.' || grid[y][x] == 0) {
					q = append(q, point{X: x, Y: y})
					grid[y][x] = 'O'
				}
			}
			grid[step.Y][step.X] = '.'
			choices--
		}
		steps--
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	utils.CheckArgs(2, os.Args)
	lines := utils.GetLinesFromFile(os.Args[1])
	grid := make([][]byte, len(lines))
	for i := range grid {
		grid[i] = []byte(lines[i])
	}
	directions := [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
	y, x := findS(lines)
	steps := 64
	bfs(grid, point{Y: y, X: x}, steps, directions)
	c := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == byte('O') {
				c++
			}
		}
	}
	fmt.Println(c)
}
