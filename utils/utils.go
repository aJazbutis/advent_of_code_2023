package utils

import	(
	"bufio"
	"os"
)

func GetLinesFromFile(fileName string) ([]string, error)	{
	var lines []string
	file, err := os.Open(fileName)
	if (err != nil)	{
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