package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"utils"
)

const (
	J = iota
	two
	three
	four
	five
	six
	seven
	eight
	nine
	T
	Q
	K
	A
	highCard
	onePair
	twoPair
	threeOfKind
	fullHouse
	fourOfKind
	fiveOfKind
)

func runeToLabel(r rune) int {
	switch r {
	case 'A':
		return A
	case 'K':
		return K
	case 'Q':
		return Q
	case 'J':
		return J
	case 'T':
		return T
	case '9':
		return nine
	case '8':
		return eight
	case '7':
		return seven
	case '6':
		return six
	case '5':
		return five
	case '4':
		return four
	case '3':
		return three
	default:
		return two
	}
}

func linesToMap(lines []string) map[string]string {
	m := make(map[string]string)
	for i := range lines {
		kv := strings.Fields(lines[i])
		m[kv[0]] = kv[1]
	}
	return m
}

func handToType(hand string) int {
	count := []int{}
	jokers := strings.Count(hand, "J")
	if jokers > 0 {
		hand = strings.Replace(hand, "J", "", -1)
	}
	for len(hand) > 0 {
		count = append(count, strings.Count(hand, string(hand[0])))
		hand = strings.Replace(hand, string(hand[0]), "", -1)
	}
	if jokers == 5 {
		count = append(count, 5)
	} else {
		max := slices.Max(count)
		idx := slices.Index(count, max)
		count[idx] += jokers
	}
	switch {
	case slices.Contains(count, 5):
		return fiveOfKind
	case slices.Contains(count, 4):
		return fourOfKind
	case slices.Contains(count, 3):
		switch {
		case slices.Contains(count, 2):
			return fullHouse
		default:
			return threeOfKind
		}
	case slices.Contains(count, 2):
		idx := slices.Index(count, 2)
		count = slices.Delete(count, idx, idx+1)
		switch {
		case slices.Contains(count, 2):
			return twoPair
		default:
			return onePair
		}
	default:
		return highCard
	}
}

func handsToTypes(hands map[string]string) map[string]int {
	m := make(map[string]int)
	for hand, _ := range hands {
		m[hand] = handToType(hand)
	}
	return m
}

func handCompare(a string, b string) int {
	for i := range a {
		cmp := runeToLabel(rune(a[i])) - runeToLabel(rune(b[i]))
		switch cmp {
		case 0:
			continue
		default:
			return cmp
		}
	}
	return 0
}

func rankByType(hands map[string]int, rank int) map[string]int {
	forSort := []string{}
	for hand := range hands {
		forSort = append(forSort, hand)
	}
	slices.SortFunc(forSort, handCompare)
	for _, hand := range forSort {
		rank++
		hands[hand] = rank
	}
	return hands
}

func handsToRanks(hands map[string]int) map[string]int {
	m := make(map[string]int)
	mm := make(map[int]map[string]int)
	for i := highCard; i <= fiveOfKind; i++ {
		mm[i] = make(map[string]int)
	}
	for hand, comb := range hands {
		switch comb {
		case fiveOfKind:
			mm[fiveOfKind][hand] = 0
		case fourOfKind:
			mm[fourOfKind][hand] = 0
		case fullHouse:
			mm[fullHouse][hand] = 0
		case threeOfKind:
			mm[threeOfKind][hand] = 0
		case twoPair:
			mm[twoPair][hand] = 0
		case onePair:
			mm[onePair][hand] = 0
		default:
			mm[highCard][hand] = 0
		}
	}
	mm[highCard] = rankByType(mm[highCard], 0)
	mm[onePair] = rankByType(mm[onePair], len(mm[highCard]))
	mm[twoPair] = rankByType(mm[twoPair], len(mm[onePair])+len(mm[highCard]))
	mm[threeOfKind] = rankByType(mm[threeOfKind], len(mm[twoPair])+len(mm[onePair])+len(mm[highCard]))
	mm[fullHouse] = rankByType(mm[fullHouse], len(mm[threeOfKind])+len(mm[twoPair])+len(mm[onePair])+len(mm[highCard]))
	mm[fourOfKind] = rankByType(mm[fourOfKind], len(mm[fullHouse])+len(mm[threeOfKind])+len(mm[twoPair])+len(mm[onePair])+len(mm[highCard]))
	mm[fiveOfKind] = rankByType(mm[fiveOfKind], len(mm[fourOfKind])+len(mm[fullHouse])+len(mm[threeOfKind])+len(mm[twoPair])+len(mm[onePair])+len(mm[highCard]))
	for _, mp := range mm {
		for hand, rank := range mp {
			m[hand] = rank
		}

	}
	return m
}

func main() {
	defer utils.Panicked()
	utils.CheckArgs(2, os.Args)
	bids := linesToMap(utils.GetLinesFromFile(os.Args[1]))
	ranks := handsToRanks(handsToTypes(bids))
	winnings := 0
	for hand, rank := range ranks {
		bid, _ := strconv.Atoi(bids[hand])
		winnings += bid * rank
	}
	fmt.Println(winnings)
}
