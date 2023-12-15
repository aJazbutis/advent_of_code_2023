package main

import (
	"advent/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hashAlgorithm(s string) int {
	currentValue := 0
	for i := range s {
		currentValue += int(s[i])
		currentValue *= 17
		currentValue %= 256
	}
	return currentValue
}

func step(s string) (int, []string) {
	ss := []string{}
	if strings.Index(s, "=") == -1 {
		ss = append(ss, s[:len(s)-1])
	} else {
		ss = strings.Split(s, "=")
	}
	return hashAlgorithm(ss[0]), ss
}

func removeLense(ss [][2]string, lense string) [][2]string {
	for i := range ss {
		if ss[i][0] == lense {
			s1 := ss[:i]
			s2 := ss[i+1:]
			s1 = append(s1, s2...)
			return s1
		}
	}
	return ss
}

func main() {
	defer utils.Panicked()
	utils.CheckArgs(2, os.Args)
	boxes := make(map[int][][2]string)
	for i := 0; i < 256; i++ {
		boxes[i] = [][2]string{}
	}
	s := utils.GetLinesFromFile(os.Args[1])
	if len(s) != 1 {
		panic("Supossed to be one string!")
	}
	ss := strings.Split(s[0], ",")
	sum := 0
	for i := range ss {
		box, step := step(ss[i])
		if len(step) == 1 {
			boxes[box] = removeLense(boxes[box], step[0])
		} else {
			lenseInBox := false
			for j := range boxes[box] {
				if boxes[box][j][0] == step[0] {
					boxes[box][j][1] = step[1]
					lenseInBox = true
					break
				}
			}
			if !lenseInBox {
				boxes[box] = append(boxes[box], [2]string{step[0], step[1]})
			}
		}
	}
	for i := range boxes {
		for j := range boxes[i] {
			focalLength, _ := strconv.Atoi(boxes[i][j][1])
			sum += (i + 1) * (j + 1) * focalLength
		}
	}
	fmt.Println(sum)
}
