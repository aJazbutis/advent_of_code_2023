package main

import (
	"advent/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	seedToSoil = iota
	soilToFert
	fertToWate
	wateToLigh
	lighToTemp
	tempToHumi
	humiToLoca
)

func putSeeds(data *[][2]int, seeds string) {
	s := strings.Fields(seeds)
	ints := []int{}
	for i := range s {
		n, err := strconv.Atoi(s[i])
		if err != nil {
			panic(err)
		}
		ints = append(ints, n)
	}
	for i := 0; i < len(ints)-1; i += 2 {
		seed := [2]int{ints[i], ints[i+1]}
		*data = append(*data, seed)
	}
}

func mapToMap(data *map[int][][3]int, lines *[]string, i int, id int) int {
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
	for i := range linesToInt {
		(*data)[id] = append((*data)[id], [3]int{linesToInt[i][1], linesToInt[i][2], linesToInt[i][0]})
	}
	return i
}

func getMapping(val int, mp [][3]int) int {
	for i := range mp {
		if val >= mp[i][0] && val < mp[i][0]+mp[i][1] {
			diff := val - mp[i][0]
			return mp[i][2] + diff
		}
	}
	return val
}

type Job struct {
	s    [2]int
	maps *map[int][][3]int
}

func minDist(jobs <-chan Job, res chan<- int) {
	for job := range jobs {
		min := job.s[0]
		for seed := job.s[0]; seed < job.s[0]+job.s[1]; seed++ {
			s := seed
			for key := seedToSoil; key <= humiToLoca; key++ {
				s = getMapping(s, (*job.maps)[key])
			}
			if s < min {
				min = s
			}
		}
		res <- min
	}
}

func moreSeeds(seeds [][2]int) [][2]int {
	moreSeeds := [][2]int{}
	for i := range seeds {
		r := seeds[i][1] / 10000
		left := seeds[i][1] % 10000
		s := seeds[i][0]
		for j := 0; j < 10000; j++ {
			if j == 9999 {
				r += left
			}
			moreSeeds = append(moreSeeds, [2]int{s, r})
			s += r
		}
	}
	return moreSeeds
}

func main() {
	defer utils.Panicked()
	utils.CheckArgs(2, os.Args)
	lines := utils.GetLinesFromFile(os.Args[1])
	seeds := [][2]int{}
	maps := make(map[int][][3]int)
	maps[seedToSoil] = [][3]int{}
	maps[soilToFert] = [][3]int{}
	maps[fertToWate] = [][3]int{}
	maps[wateToLigh] = [][3]int{}
	maps[lighToTemp] = [][3]int{}
	maps[tempToHumi] = [][3]int{}
	maps[humiToLoca] = [][3]int{}
	id := seedToSoil
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
				i = mapToMap(&maps, &lines, i+1, id)
				id++
			}
		}
	}
	seeds = moreSeeds(seeds)
	jobs := make(chan Job, len(seeds))
	results := make(chan int, len(seeds))
	for w := 0; w < len(seeds); w++ {
		go minDist(jobs, results)
	}
	for i := 0; i < len(seeds); i++ {
		jobs <- Job{
			s:    seeds[i],
			maps: &maps,
		}
	}
	close(jobs)
	res := seeds[0][0]
	for i := 0; i < len(seeds); i++ {
		tmp := <-results
		if res > tmp {
			res = tmp
		}
	}
	close(results)
	fmt.Println(res)
}
