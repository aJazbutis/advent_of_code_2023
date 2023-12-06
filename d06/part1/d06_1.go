package main

import (
	"advent/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func linesToData(lines []string) map[int]int {
	m := make(map[int]int)
	tmp := []string{}
	for i := range lines {
		tmp = append(tmp, strings.Split(lines[i], ":")[1])
	}
	tmp2 := [][]string{}
	for i := range tmp {
		tmp2 = append(tmp2, strings.Fields(tmp[i]))
	}
	for i := range tmp2[0] {
		time, _ := strconv.Atoi(tmp2[0][i])
		dist, _ := strconv.Atoi(tmp2[1][i])
		m[time] = dist
	}
	return m
}

func ways(time int, record int) int {
	ret := 0
	speed := 0
	for speed < time {
		dist := speed * (time - speed)
		if dist > record {
			ret++
		}
		speed++
	}
	return ret
}
func main() {
	defer utils.Panicked()
	utils.CheckArgs(2, os.Args)
	data := linesToData(utils.GetLinesFromFile(os.Args[1]))
	ret := 1
	for time, record := range data {
		ret *= ways(time, record)
	}
	fmt.Println(ret)
}
