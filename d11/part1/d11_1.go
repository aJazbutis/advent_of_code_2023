package main

import (
	"advent/utils"
	"fmt"
	"os"
	"strings"
)

func expandColumns(lines []string) []string {
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
	for len(idxs) > 0 {
		for i := range lines {
			lines[i] = utils.StringIsertAtIdx(lines[i], idxs[0], ".")
		}
		if len(idxs) > 0 {
			idxs = idxs[1:]
		}
		for i := range idxs {
			idxs[i]++
		}
	}
	return lines
}

func expandLines(lines []string) []string {
	for l := 0; l < len(lines); l++ {
		if strings.Index(lines[l], "#") == -1 {
			lines = append(lines[:l+1], lines[l:]...)
			l++
		}
	}
	return lines
}

func expandGalaxies(lines []string) [][2]int {
	coords := [][2]int{}
	galaxies := expandLines(expandColumns(lines))
	for i := range galaxies {
		for j := range galaxies[i] {
			if galaxies[i][j] == '#' {
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
					a[1]--
					d++
					continue
				} else if a[0] != b[0] {
					a[0]++
					d++
					continue
				}
			}
			if d%2 == 1 {
				if a[0] != b[0] {
					a[0]++
					d++
					continue
				} else if a[1] != b[1] {
					a[1]--
					d++
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
					a[1]++
					d++
					continue
				} else if a[0] != b[0] {
					a[0]++
					d++
					continue
				}
			}
			if d%2 == 1 {
				if a[0] != b[0] {
					a[0]++
					d++
					continue
				} else if a[1] != b[1] {
					a[1]++
					d++
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
	fmt.Println(sumShortestDist(expandGalaxies(utils.GetLinesFromFile(os.Args[1]))))
}
