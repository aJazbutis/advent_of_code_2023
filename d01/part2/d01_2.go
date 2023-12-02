package main

import (
	"fmt"
	"advent/utils"
	"os"
	"strings"
	"errors"
	"strconv"
)

func getDigit(m map[string]string, line string) string	{
	idx := len(line);
	ret := "";
	for key, val := range m {
		posKey := strings.Index(line, key)
		posVal := strings.Index(line, val)
		if posKey == 0 || posVal == 0	{
			return key
		}
		if posKey > 0 && posKey < idx	{
			idx = posKey;
			ret = key;
		}
		if posVal > 0 && posVal < idx	{
			idx = posVal
			ret = key;
		}
	}
	return ret
}

func main()	{
	if (len(os.Args) != 2)	{
		utils.ErrorExit(errors.New("No args"))
	}
	lines, err := utils.GetLinesFromFile(os.Args[1])
	if err != nil	{
		utils.ErrorExit(err)
	}
	words := []string{
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}
	numbers:= make(map[string]string);
	reverse:= make(map[string]string);
	for i, word := range words	{
		key:= strconv.Itoa(i + 1);
		numbers[key] = word;
		reverse[key] = utils.ReverseString(word);
	}
	sum := 0;
	for _, line := range lines	{
		rev := utils.ReverseString(line);
		num := getDigit(numbers, line) + getDigit(reverse, rev)
		n, err := strconv.Atoi(num)
		if err == nil {
			sum += n;
		} else {
			utils.ErrorExit(err);
		}
	}
	fmt.Println(sum)
}