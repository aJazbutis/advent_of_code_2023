package main

import (
	"advent/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func linesToData(lines []string) (int, int) {
	tmp := []string{}
	for i := range lines {
		tmp = append(tmp, strings.Split(lines[i], ":")[1])
	}
	tmp2 := [][]string{}
	for i := range tmp {
		tmp2 = append(tmp2, strings.Fields(tmp[i]))
	}
	t, d := "", ""
	for i := range tmp2[0] {
		t += tmp2[0][i]
		d += tmp2[1][i]
	}
	time, _ := strconv.Atoi(t)
	dist, _ := strconv.Atoi(d)
	return time, dist
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
	fmt.Println(ways(linesToData(utils.GetLinesFromFile(os.Args[1]))))
}
