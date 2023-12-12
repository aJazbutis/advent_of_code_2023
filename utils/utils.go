package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func GetLinesFromFile(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
		// return lines, err;
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	file.Close()
	return lines
}

func Abc() string {
	abc := ""

	for c := 'a'; c <= 'z'; c++ {
		abc += string(c)
	}
	return abc
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ExtractInt(line string, idx int) int {
	b := idx
	e := idx
	for b > 0 && unicode.IsDigit(rune(line[b-1])) {
		b--
	}
	for e < len(line) && unicode.IsDigit(rune(line[e])) {
		e++
	}
	num, err := strconv.Atoi(string((line[b:e])))
	if err != nil {
		panic(err)
	}
	fmt.Println(num)
	return num
}

func ErrorExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func CheckArgs(i int, args []string) {
	if len(args) != i {
		panic("Bad args")
	}
}

func Panicked() {
	if err := recover(); err != nil {
		log.Fatal(err)
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM(a, b) * GCD(a, b) = a * b.

func LCM(a int, b int) int {
	lcm := (a * b) / gcd(a, b)
	return lcm
}

func StringIsertAtIdx(s string, idx int, insert string) string {
	if idx < 0 || idx > len(s)-1 {
		panic("Idx out of range")
	}
	return s[:idx] + insert + s[idx:]
}
