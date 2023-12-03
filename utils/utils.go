package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func GetLinesFromFile(fileName string) ([]string, error)	{
	var lines []string
	file, err := os.Open(fileName)
	if (err != nil)	{
		// panic(err)
		return lines, err;
	}
	fileScanner := bufio.NewScanner(file);
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan()	{
		lines = append(lines, fileScanner.Text())
	}
	file.Close()
	return lines, nil
}

func Abc()(string)	{
	abc := "";

	for c := 'a'; c <= 'z'; c++	{
		abc += string(c);
	}
	return abc
}

func ReverseString(s string) (string) {
	runes := []rune(s);
	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1	{
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes);
}

func ExtractInt(line string, idx int ) (int)	{
	b := idx;
	e := idx;
	for b > 0 && unicode.IsDigit(rune(line[b - 1]))	{
		b--;
	}	
	for e < len(line) && unicode.IsDigit(rune(line[e]))	{
		e++;
	}
	num, err := strconv.Atoi(string((line[b:e])))
	if err != nil	{
		panic(err)
	}
	fmt.Println(num)
	return num
}

func ErrorExit(err error)	{
	fmt.Println(err);
	os.Exit(1)
}