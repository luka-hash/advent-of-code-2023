// Copyright © 2023 Luka Ivanović
// This code is licensed under the 2-clause BSD licence (see LICENCE for details)

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

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

func sum(ss []int) int {
	acc := 0
	for _, n := range ss {
		acc += n
	}
	return acc
}

func reverse[t any](ss []t) []t {
	for i := len(ss)/2 - 1; i >= 0; i-- {
		opp := len(ss) - 1 - i
		ss[i], ss[opp] = ss[opp], ss[i]
	}
	return ss
}

func prepend[t any](ss []t, s ...t) []t {
	return append(append([]t{}, s...), ss...)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func stringToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func part1() {
	lines, err := readLines("input.txt")
	if err != nil {
		panic(err) }
	totalPoints := 0
	for _, line := range lines {
		_, cards, _ := strings.Cut(line, ":")
		points := 0
		winningNumbers, myNumbersString, _ := strings.Cut(cards, "|")
		myNumbers := strings.Fields(myNumbersString)
		for _, number := range myNumbers {
			if strings.Contains(winningNumbers, " "+number+" ") {
				if points == 0 {
					points = 1
				} else {
					points <<= 1
				}
			}
		}
		totalPoints += points
	}
	fmt.Println(totalPoints)
}

func part2() {
	lines, err := readLines("input.txt")
	if err != nil {
		panic(err)
	}
	m := make(map[int]int)
	p := make(map[int]int)
	for i := 1; i <= len(lines); i += 1 {
		line := lines[i-1]
		_, cards, _ := strings.Cut(line, ":")
		matches := 0
		winningNumbers, myNumbersString, _ := strings.Cut(cards, "|")
		myNumbers := strings.Fields(myNumbersString)
		for _, number := range myNumbers {
			if strings.Contains(winningNumbers, " "+number+" ") {
				matches += 1
			}
		}
		m[i] = matches
		p[i] = 1
	}
	total := 0
	for i := 1; i <= len(lines); i += 1 {
		// fmt.Printf("from card %d (with %d matches, and %d occurances):\n", i, m[i], p[i])
		// There is no need to check if there are any matches, since the following loop will not work when there are none.
		l := i + 1
		h := min(i+m[i], len(lines))
		// fmt.Printf("\tUpdating cards %d..%d\n", l, h)
		for j := l; j <= h; j += 1 {
			p[j] += p[i]
		}
		total += p[i]
	}
	fmt.Println(total)
}
