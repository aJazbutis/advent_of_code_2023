package main

import (
	"advent/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("No arguments")
		os.Exit(1)
	}
	lines, err := utils.GetLinesFromFile(os.Args[1])
	if err == nil {
		abc := utils.Abc()
		sum := 0
		for _, line := range lines {
			line = strings.Trim(line, abc)
			if len(line) > 0 {
				num := string(line[0]) + string(line[len(line)-1])
				n, err := strconv.Atoi(num)
				if err == nil {
					sum += n
				} else {
					fmt.Println(err)
					os.Exit(1)
				}
			}
		}
		fmt.Println(sum)
	} else {
		fmt.Println(err)
		os.Exit(1)
	}
}
