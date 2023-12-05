package main

import (
	"advent/utils"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func putSeeds(data *[][]int, seeds string) {
	s := strings.Fields(seeds)
	ints := []int{}
	for i := range s {
		n, err := strconv.Atoi(s[i])
		if err != nil {
			panic(err)
		}
		ints = append(ints, n)
	}
	slices.Sort(ints)
	for i := range ints {
		seed := []int{ints[i]}
		*data = append(*data, seed)
	}
}
func mapToMap(data *[][]int, lines *[]string, i int) int {
	linesToInt := [][]int{}
	for i < len(*lines) && (*lines)[i] != "" {
		info := strings.Fields((*lines)[i])
		a := []int{}
		for i := range info {
			n, err := strconv.Atoi(info[i])
			utils.CheckError(err)
			a = append(a, n)
		}
		linesToInt = append(linesToInt, a)
		i++
	}
	for i := range *data {
		mapped := false
		x := (*data)[i][len((*data)[i])-1]
		for _, mapping := range linesToInt {
			if mapped {
				break
			}
			if x >= mapping[1] && x < mapping[1]+mapping[2] {
				diff := x - mapping[1]
				(*data)[i] = append((*data)[i], mapping[0]+diff)
				mapped = true
			}
		}
		if !mapped {
			(*data)[i] = append((*data)[i], x)
		}
	}
	return i
}
func main() {
	defer utils.Panicked()
	utils.CheckArgs(2, os.Args)
	lines := utils.GetLinesFromFile(os.Args[1])
	seeds := [][]int{}
	for i := range lines {
		if lines[i] == "" {
			continue
		}
		if strings.Index(lines[i], ":") != -1 {
			info := strings.Split(lines[i], ":")
			switch info[0] {
			case "seeds":
				putSeeds(&seeds, info[1])
			default:
				i = mapToMap(&seeds, &lines, i+1)
			}
		}
	}
	ret := seeds[0][len(seeds[0])-1]
	for i := range seeds {
		if ret > seeds[i][len(seeds[i])-1] {
			ret = seeds[i][len(seeds[i])-1]
		}
	}
	fmt.Println(ret)
}
