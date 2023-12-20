package main

import (
	"advent/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Part struct {
	x int
	m int
	a int
	s int
}

type Rule struct {
	category  string
	condition string
	criteria  string
	result    string
}

type Workflow map[string][]Rule

func getPart(line string) Part {
	part := Part{}
	line = strings.Trim(line, "{}")
	lin := strings.Split(line, ",")
	for i := range lin {
		li := strings.Split(lin[i], "=")
		value, _ := strconv.Atoi(li[1])
		switch li[0] {
		case "x":
			part.x = value
		case "m":
			part.m = value
		case "a":
			part.a = value
		case "s":
			part.s = value
		}
	}
	return part
}

func getRules(line string) []Rule {
	rules := []Rule{}
	line = strings.Trim(line, "}")
	lin := strings.Split(line, ",")
	for i := range lin {
		li := strings.Split(lin[i], ":")
		if len(li) == 1 {
			rules = append(rules, Rule{result: li[0]})
		} else {
			rules = append(rules, Rule{
				category:  li[0][:1],
				condition: li[0][1:2],
				criteria:  li[0][2:],
				result:    li[1],
			})
		}
	}
	return rules
}

func getWorkflowAndParts(lines []string) (Workflow, []Part) {
	parts := []Part{}
	workflow := make(map[string][]Rule)
	for i := 0; i < len(lines); i++ {
		if lines[i] != "" {
			flow := strings.Split(lines[i], "{")
			workflow[flow[0]] = getRules(flow[1])
		} else {
			i++
			for ; i < len(lines); i++ {
				parts = append(parts, getPart(lines[i]))
			}
		}
	}
	return workflow, parts
}

func flow(part Part, workflow *Workflow, rules []Rule) bool {
	if rules[0].category == "" {
		rule, ok := (*workflow)[rules[0].result]
		if ok {
			return flow(part, workflow, rule)
		} else {
			return rules[0].result == "A"
		}
	}
	num, _ := strconv.Atoi(rules[0].criteria)
	n := 0
	switch rules[0].category {
	case "x":
		n = part.x
	case "m":
		n = part.m
	case "a":
		n = part.a
	case "s":
		n = part.s
	}
	passes := true
	if rules[0].condition == "<" {
		passes = n < num
	} else {
		passes = n > num
	}
	if passes {
		res := rules[0].result
		rule, ok := (*workflow)[res]
		if ok {
			return flow(part, workflow, rule)
		} else {
			return res == "A"
		}
	} else {
		return flow(part, workflow, rules[1:])
	}
}

func main() {
	defer utils.Panicked()
	utils.CheckArgs(2, os.Args)
	workflow, parts := getWorkflowAndParts(utils.GetLinesFromFile(os.Args[1]))
	sum := 0
	for _, part := range parts {
		passes := flow(part, &workflow, workflow["in"])
		if passes {
			sum += part.x + part.m + part.a + part.s
		}
	}
	fmt.Println(sum)
}
