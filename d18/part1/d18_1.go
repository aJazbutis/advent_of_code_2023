package main

import (
	"advent/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func plan(lines []string) [][2]string {
	ret := [][2]string{}
	for i := range lines {
		step := strings.Fields(lines[i])
		ret = append(ret, [2]string{step[0], step[1]})
	}
	return ret
}

func initSite(step [2]string, site *[][]rune) [2]int {
	count, _ := strconv.Atoi(step[1])
	switch step[0] {
	case "U", "D":
		for count > 0 {
			*site = append(*site, []rune("#"))
			count--
		}
	case "R", "L":
		s := ""
		for count > 0 {
			s += "#"
			count--
		}
		*site = append(*site, []rune(s))
	}
	switch step[0] {
	case "R":
		return [2]int{0, len((*site)[0]) - 1}
	case "D":
		return [2]int{len(*site) - 1, 0}
	default:
		return [2]int{0, 0}
	}
}

func verticaliai(y int, site *[][]rune, pos *[2]int) {
	diff := (*pos)[0] + y
	if diff < 0 || diff > len(*site)-1 {
		s := ""
		for i := len((*site)[0]); i > 0; i-- {
			s += "."
		}
		if diff < 0 {
			tmp := [][]rune{}
			for i := 0; i > diff; i-- {
				tmp = append(tmp, []rune(s))
			}
			tmp = append(tmp, (*site)...)
			*site = tmp
			(*pos)[0] += -diff
		} else {
			for i := (*pos)[0]; i < diff; i++ {
				*site = append(*site, []rune(s))
			}
		}
	}
	if y > 0 {
		for y > 0 {
			(*pos)[0]++
			y--
			(*site)[(*pos)[0]][(*pos)[1]] = '#'
		}
	} else {
		for y < 0 {
			(*pos)[0]--
			y++
			(*site)[(*pos)[0]][(*pos)[1]] = '#'
		}
	}
}

func horizontaliai(x int, site *[][]rune, pos *[2]int) {
	diff := (*pos)[1] + x
	if diff < 0 || diff >= len((*site)[0])-1 {
		s := ""
		if diff < 0 {
			for i := 0; i > diff; i-- {
				s += "."
			}
			for i := range *site {
				(*site)[i] = append([]rune(s), (*site)[i]...)
			}
			(*pos)[1] += -diff
		} else {
			for i := 0; i < diff-(len((*site)[0])-1); i++ {
				s += "."
			}
			for i := range *site {
				(*site)[i] = append((*site)[i], []rune(s)...)
			}
		}
	}
	if x > 0 {
		for x > 0 {
			(*pos)[1]++
			x--
			(*site)[(*pos)[0]][(*pos)[1]] = '#'
		}
	} else {
		for x < 0 {
			(*pos)[1]--
			x++
			(*site)[(*pos)[0]][(*pos)[1]] = '#'
		}
	}
}

func count(site *[][]rune) int {
	c := 0
	for _, line := range *site {
		for i := range line {
			if line[i] == '#' || line[i] == '.' {
				c++
			}
		}
	}
	return c
}

func clear(site *[][]rune) {
	for y := range *site {
		for x := 0; x < len((*site)[y]); x++ {
			if (*site)[y][x] == '.' {
				(*site)[y][x] = ' '
			}
			if (*site)[y][x] == '#' {
				break
			}
		}
		for x := len((*site)[y]) - 1; x >= 0; x-- {
			if (*site)[y][x] == '.' {
				(*site)[y][x] = ' '
			}
			if (*site)[y][x] == '#' {
				break
			}
		}
	}
	for x := range (*site)[0] {
		for y := 0; y < len(*site); y++ {
			if (*site)[y][x] == '.' {
				(*site)[y][x] = ' '
			}
			if (*site)[y][x] == '#' {
				break
			}
		}
		for y := len(*site) - 1; y >= 0; y-- {
			if (*site)[y][x] == '.' {
				(*site)[y][x] = ' '
			}
			if (*site)[y][x] == '#' {
				break
			}
		}
	}
	for i := 2; i > 0; i-- {
		for y := range *site {
			for x := 0; x < len((*site)[y]); x++ {
				switch {
				case (*site)[y][x] == ' ':
					if y > 0 && (*site)[y-1][x] == '.' {
						y1 := y - 1
						for y1 >= 0 && (*site)[y1][x] == '.' {
							(*site)[y1][x] = ' '
							y1--
						}
					}
					if y < len(*site)-1 && (*site)[y+1][x] == '.' {
						y1 := y + 1
						for y1 < len(*site) && (*site)[y1][x] == '.' {
							(*site)[y1][x] = ' '
							y1++
						}
					}
					if x+1 < len((*site)[y])-1 && (*site)[y][x+1] == '.' {
						(*site)[y][x+1] = ' '
					}
					if x-1 > 0 && (*site)[y][x-1] == '.' {
						x1 := x - 1
						for x1 >= 0 && (*site)[y][x1] == '.' {
							(*site)[y][x1] = ' '
							x1--
						}
					}
				}
			}
			for x := len((*site)[y]) - 1; x >= 0; x-- {
				switch {
				case (*site)[y][x] == ' ':
					if y > 0 && (*site)[y-1][x] == '.' {
						y1 := y - 1
						for y1 >= 0 && (*site)[y1][x] == '.' {
							(*site)[y1][x] = ' '
							y1--
						}
					}
					if y < len(*site)-1 && (*site)[y+1][x] == '.' {
						y1 := y + 1
						for y1 < len(*site) && (*site)[y1][x] == '.' {
							(*site)[y1][x] = ' '
							y1++
						}
					}
					if x+1 < len((*site)[y])-1 && (*site)[y][x+1] == '.' {
						x1 := x + 1
						for x1 <= len((*site)[y])-1 && (*site)[y][x1] == '.' {
							(*site)[y][x1] = ' '
							x1++
						}
					}
					if x-1 > 0 && (*site)[y][x-1] == '.' {
						(*site)[y][x-1] = ' '
					}
				}
			}
		}
	}
}

func dig(plan [][2]string) int {
	site := [][]rune{}
	pos := initSite(plan[0], &site)
	for i := 1; i < len(plan); i++ {
		dist, _ := strconv.Atoi(plan[i][1])
		switch plan[i][0] {
		case "U":
			verticaliai(-dist, &site, &pos)
		case "D":
			verticaliai(dist, &site, &pos)
		case "R":
			horizontaliai(dist, &site, &pos)
		case "L":
			horizontaliai(-dist, &site, &pos)
		}
	}
	clear(&site)
	return count(&site)
}

func main() {
	defer utils.Panicked()
	utils.CheckArgs(2, os.Args)
	plan := plan(utils.GetLinesFromFile(os.Args[1]))
	fmt.Println(dig(plan))
}
