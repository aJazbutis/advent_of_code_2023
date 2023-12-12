package main

import (
	"advent/utils"
	"fmt"
	"os"
	"strings"
)

func expandColumns(lines []string) []int {
	idxs := []int{}
	for i := range lines[0] {
		empty := true
		for j := range lines {
			if len(lines[j]) == 0 {
			}
			if lines[j][i] == '#' {
				empty = false
				break
			}
		}
		if empty {
			idxs = append(idxs, i)
		}
	}
	return idxs
}

func expandLines(lines []string) []int {
	idxs := []int{}
	for l := 0; l < len(lines); l++ {
		if strings.Index(lines[l], "#") == -1 {
			idxs = append(idxs, l)
		}
	}
	return idxs
}

func getGalaxies(lines []string) [][2]int {
	coords := [][2]int{}
	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == '#' {
				coords = append(coords, [2]int{i, j})
			}
		}
	}
	return coords
}

func distanz(a [2]int, b [2]int) int {
	d := 0
	switch {
	case a[1] >= b[1]:
		for true {
			if d%2 == 0 {
				if a[1] != b[1] {
					if a[1]-b[1] > 1000000 {
						a[1] -= 1000000
						d += 1000000
					} else {
						a[1]--
						d++
					}
					continue
				} else if a[0] != b[0] {
					if b[0]-a[0] > 1000000 {
						a[0] += 1000000
						d += 1000000
					} else {
						a[0]++
						d++
					}
					continue
				}
			}
			if d%2 == 1 {
				if a[0] != b[0] {
					a[0]++
					d++
					continue
				} else if a[1] != b[1] {
					if a[1]-b[1] > 1000000 {
						a[1] -= 1000000
						d += 1000000
					} else {
						a[1]--
						d++
					}
					continue
				}
			}
			if a[0] == b[0] && a[1] == b[1] {
				break
			}
		}
	case a[1] <= b[1]:
		for true {
			if d%2 == 0 {
				if a[1] != b[1] {
					if b[1]-a[1] > 1000000 {
						a[1] += 1000000
						d += 1000000
					} else {
						a[1]++
						d++
					}
					continue
				} else if a[0] != b[0] {
					if b[0]-a[0] > 1000000 {
						a[0] += 1000000
						d += 1000000
					} else {
						a[0]++
						d++
					}
					continue
				}
			}
			if d%2 == 1 {
				if a[0] != b[0] {
					if b[0]-a[0] > 1000000 {
						a[0] += 1000000
						d += 1000000
					} else {
						a[0]++
						d++
					}
					continue
				} else if a[1] != b[1] {
					if b[1]-a[1] > 1000000 {
						a[1] += 1000000
						d += 1000000
					} else {
						a[1]++
						d++
					}
					continue
				}
			}
			if a[0] == b[0] && a[1] == b[1] {
				break
			}
		}
	}
	return d
}

func expandGalaxies(galaxies [][2]int, cols []int, rows []int) [][2]int {
	for g := range galaxies {
		c := 0
		for j := range cols {
			if galaxies[g][1] > cols[j] {
				c++
			}
		}
		galaxies[g][1] = galaxies[g][1] + c*1000000 - c
	}
	for g := range galaxies {
		c := 0
		for j := range rows {
			if galaxies[g][0] > rows[j] {
				c++
			}
		}
		galaxies[g][0] = galaxies[g][0] + c*1000000 - c
	}
	return galaxies
}

func sumShortestDist(coords [][2]int) int {
	sum := 0
	for g := 0; g < len(coords)-1; g++ {
		for other := g + 1; other < len(coords); other++ {
			sum += distanz(coords[g], coords[other])
		}
	}
	return sum
}

func main() {
	defer utils.Panicked()
	utils.CheckArgs(2, os.Args)
	lines := utils.GetLinesFromFile(os.Args[1])
	g := getGalaxies(lines)
	c := expandColumns(lines)
	r := expandLines(lines)
	fmt.Println(sumShortestDist(expandGalaxies(g, c, r)))
}
