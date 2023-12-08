// Copyright © 2023 Luka Ivanović
// This code is licensed under the 2-clause BSD licence (see LICENCE for details)

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Pair struct {
	hand string
	bid  int
}

//go:generate stringer -type=HandType

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func count[t comparable](ss []t, s t) int {
	c := 0
	for _, e := range ss {
		if e == s {
			c += 1
		}
	}
	return c
}

func stringToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func cardToInt(b byte) int {
	switch b {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case 'T':
		return 10
	default:
		return int(b - '0')
	}
}

func cardToIntUpdated(b byte) int {
	if b == 'J' {
		return 1
	}
	return cardToInt(b)
}

func handToType(hand string) HandType {
	m := make(map[rune]int)
	max := 0
	d := 0
	for _, r := range hand {
		if _, ok := m[r]; !ok {
			d += 1
		}
		m[r] += 1
		if m[r] > max {
			max = m[r]
		}
	}
	switch d {
	case 1:
		return FiveOfAKind
	case 2:
		if max == 4 {
			return FourOfAKind
		} else { // max == 3 // to satisfy the compiler..
			return FullHouse
		}
	case 3:
		if max == 3 {
			return ThreeOfAKind
		} else {
			return TwoPairs
		}
	case 4:
		return OnePair
	default:
		return HighCard
	}
}

func handToTypeUpdated(hand string) HandType {
	jokers := count([]byte(hand), 'J')
	t := handToType(hand)
	switch t {
	case FiveOfAKind:
		return FiveOfAKind
	case FourOfAKind:
		if jokers == 1 || jokers == 4 {
			return FiveOfAKind
		} else {
			return FourOfAKind
		}
	case FullHouse:
		if jokers == 2 || jokers == 3 {
			return FiveOfAKind
		} else {
			return FullHouse
		}
	case ThreeOfAKind:
		if jokers == 1 || jokers == 3 {
			return FourOfAKind
		} else {
			return ThreeOfAKind
		}
	case TwoPairs:
		if jokers == 1 {
			return FullHouse
		} else if jokers == 2 {
			return FourOfAKind
		} else {
			return TwoPairs
		}
	case OnePair:
		if jokers == 1 || jokers == 2 {
			return ThreeOfAKind
		} else {
			return OnePair
		}
	default:
		if jokers == 1 {
			return OnePair
		} else {
			return HighCard
		}
	}
}

func compareTwoHands(a, b string) bool {
	aHandType := handToType(a)
	bHandType := handToType(b)
	if aHandType == bHandType {
		for i := 0; i < len(a); i += 1 {
			if a[i] != b[i] {
				return cardToInt(a[i]) < cardToInt(b[i])
			}
		}
	}
	return aHandType < bHandType
}

func compareTwoHandsUpdated(a, b string) bool {
	aHandType := handToTypeUpdated(a)
	bHandType := handToTypeUpdated(b)
	if aHandType == bHandType {
		for i := 0; i < len(a); i += 1 {
			if a[i] != b[i] {
				return cardToIntUpdated(a[i]) < cardToIntUpdated(b[i])
			}
		}
	}
	return aHandType < bHandType
}

func part1() {
	lines, err := readLines("input.txt")
	if err != nil {
		panic(err)
	}
	_ = lines
	pairs := make([]Pair, 0)
	for _, line := range lines {
		hand, bid, _ := strings.Cut(line, " ")
		pairs = append(pairs, Pair{hand, stringToInt(bid)})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return compareTwoHands(pairs[i].hand, pairs[j].hand)
	})
	acc := 0
	for i, pair := range pairs {
		acc += (i + 1) * pair.bid
	}
	fmt.Println(acc)
}

func part2() {
	lines, err := readLines("input.txt")
	if err != nil {
		panic(err)
	}
	_ = lines
	pairs := make([]Pair, 0)
	for _, line := range lines {
		hand, bid, _ := strings.Cut(line, " ")
		pairs = append(pairs, Pair{hand, stringToInt(bid)})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return compareTwoHandsUpdated(pairs[i].hand, pairs[j].hand)
	})
	acc := 0
	for i, pair := range pairs {
		acc += (i + 1) * pair.bid
	}
	fmt.Println(acc)
}
