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

func readLines(filename string) ([][]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := make([][]byte, 0)
	for scanner.Scan() {
		lines = append(lines, []byte(scanner.Text()))
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

func stringToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func isSymbol(b byte) bool {
	if b == '+' || b == '=' || b == '-' || b == '$' || b == '&' || b == '@' || b == '*' || b == '/' || b == '#' || b == '%' {
		return true
	}
	return false
}

func isDigit(b byte) bool {
	if b == '0' || b == '1' || b == '2' || b == '3' || b == '4' || b == '5' || b == '6' || b == '7' || b == '8' || b == '9' {
		return true
	}
	return false
}

func getSurroundingParts(i, j int, schematics *[][]byte) []int {
	parts := make([]int, 0)
	if i-1 >= 0 {
		row := make([]byte, 0)
		row = append(row, (*schematics)[i-1][j])
		for k := j - 1; k >= 0 && isDigit((*schematics)[i-1][k]); k -= 1 {
			row = prepend(row, (*schematics)[i-1][k])
		}
		for k := j + 1; k <= len((*schematics)[i-1])-1 && isDigit((*schematics)[i-1][k]); k += 1 {
			row = append(row, (*schematics)[i-1][k])
		}
		// Edge case I tested for:
		// for i := range row {
		// 	if !isDigit(row[i]) {
		// 		row[i] = '.'
		// 	}
		// }
		for _, number := range strings.Split(string(row), ".") {
			if len(number) > 0 {
				parts = append(parts, stringToInt(number))
			}
		}
	}
	if j-1 >= 0 {
		number := make([]byte, 0)
		for k := j - 1; k >= 0 && isDigit((*schematics)[i][k]); k -= 1 {
			number = prepend(number, (*schematics)[i][k])
		}
		if len(number) > 0 {
			parts = append(parts, stringToInt(string(number)))
		}
	}
	if j+1 <= len((*schematics)[i])-1 {
		number := make([]byte, 0)
		for k := j + 1; k <= len((*schematics)[i])-1 && isDigit((*schematics)[i][k]); k += 1 {
			number = append(number, (*schematics)[i][k])
		}
		if len(number) > 0 {
			parts = append(parts, stringToInt(string(number)))
		}
	}
	if i+1 <= len(*schematics)-1 {
		row := make([]byte, 0)
		row = append(row, (*schematics)[i+1][j])
		for k := j - 1; k >= 0 && isDigit((*schematics)[i+1][k]); k -= 1 {
			row = prepend(row, (*schematics)[i+1][k])
		}
		for k := j + 1; k <= len((*schematics)[i+1])-1 && isDigit((*schematics)[i+1][k]); k += 1 {
			row = append(row, (*schematics)[i+1][k])
		}
		// for i := range row {
		// 	if !isDigit(row[i]) {
		// 		row[i] = '.'
		// 	}
		// }
		for _, number := range strings.Split(string(row), ".") {
			if len(number) > 0 {
				parts = append(parts, stringToInt(number))
			}
		}
	}
	return parts
}

func part1() {
	schematics, err := readLines("input.txt")
	if err != nil {
		panic(err)
	}
	parts := make([]int, 0)
	for i := 0; i < len(schematics); i += 1 {
		for j := 0; j < len(schematics[i]); j += 1 {
			if isSymbol(schematics[i][j]) {
				for _, part := range getSurroundingParts(i, j, &schematics) {
					parts = append(parts, part)
				}
			}
		}
	}
	fmt.Println(sum(parts))
}

func part2() {
	schematics, err := readLines("input.txt")
	if err != nil {
		panic(err)
	}
	gearRatios := make([]int, 0)
	for i := 0; i < len(schematics); i += 1 {
		for j := 0; j < len(schematics[i]); j += 1 {
			if isSymbol(schematics[i][j]) && schematics[i][j] == '*' {
				parts := getSurroundingParts(i, j, &schematics)
				if len(parts) == 2 {
					gearRatios = append(gearRatios, parts[0]*parts[1])
				}
			}
		}
	}
	fmt.Println(sum(gearRatios))
}

