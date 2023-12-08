package main

import (
	"advent/utils"
	"fmt"
	"os"
	"strings"
)

func nodeToData(node string) (string, [2]string) {
	s := strings.Split(node, "=")
	key := strings.Trim(s[0], " ")
	v := strings.Trim(s[1], " ()")
	va := strings.Split(v, ",")
	val := [2]string{strings.Trim(va[0], " "), strings.Trim(va[1], " ")}
	return key, val
}

func getData(lines []string, instructions *string) map[string][2]string {
	m := make(map[string][2]string)
	for i := range lines {
		switch {
		case i == 0:
			*instructions = lines[i]
		case lines[i] == "":
			continue
		default:
			key, val := nodeToData(lines[i])
			m[key] = val
		}
	}
	return m
}
func main() {
	defer utils.Panicked()
	utils.CheckArgs(2, os.Args)
	instructions := ""
	data := getData(utils.GetLinesFromFile(os.Args[1]), &instructions)
	aaa := []string{}
	for key := range data {
		if key[2] == 'A' {
			aaa = append(aaa, key)
		}
	}
	steps := []int{}
	for _, a := range aaa {
		c := 0
		for i := 0; i < len(instructions); i++ {
			c++
			switch instructions[i] {
			case 'L':
				a = data[a][0]
			case 'R':
				a = data[a][1]
			}
			if a[2] == 'Z' {
				steps = append(steps, c)
				break
			}
			if i == len(instructions)-1 {
				i = -1
			}
		}
	}
	ret := utils.LCM(steps[0], steps[1])
	for i := 2; i < len(steps); i++ {
		ret = utils.LCM(ret, steps[i])
	}
	fmt.Println(ret)
}
