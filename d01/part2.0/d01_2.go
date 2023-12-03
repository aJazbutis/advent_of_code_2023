package main

import (
	"advent/utils"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getNumber(m map[string]string, line string) string {
	idx := len(line)
	firstDigit := ""
	for key, val := range m {
		posKey := strings.Index(line, key)
		posVal := strings.Index(line, val)
		if posKey == 0 || posVal == 0 {
			firstDigit = key
			break
		}
		if posKey > 0 && posKey < idx {
			idx = posKey
			firstDigit = key
		}
		if posVal > 0 && posVal < idx {
			idx = posVal
			firstDigit = key
		}
	}
	lastDigit := ""
	idx = -1
	for key, val := range m {
		posKey := strings.LastIndex(line, key)
		posVal := strings.LastIndex(line, val)
		if posKey == len(line)-1 {
			lastDigit = key
			break
		}
		if posKey > idx {
			idx = posKey
			lastDigit = key
		}
		if posVal > idx {
			idx = posVal
			lastDigit = key
		}
	}
	return firstDigit + lastDigit
}

func main() {
	if len(os.Args) != 2 {
		utils.ErrorExit(errors.New("No args"))
	}
	lines, err := utils.GetLinesFromFile(os.Args[1])
	if err != nil {
		utils.ErrorExit(err)
	}
	words := []string{
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}
	numbers := make(map[string]string)
	for i, word := range words {
		key := strconv.Itoa(i + 1)
		numbers[key] = word
	}
	sum := 0
	for _, line := range lines {
		num := getNumber(numbers, line)
		if len(num) < 2 {

			fmt.Println(num)
			fmt.Println(line)
		}
		n, err := strconv.Atoi(num)
		if err == nil {
			sum += n
		} else {
			utils.ErrorExit(err)
		}
	}
	fmt.Println(sum)
}
