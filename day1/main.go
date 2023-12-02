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

func isNumber(r rune) bool {
	if r < 48 || r > 57 {
		return false
	}
	return true
}

func runeToInt(r rune) int {
	if r < 48 || r > 57 {
		panic(string(r) + "is not a number")
	}
	return int(r - '0')
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
		panic(err)
	}
	calibration_values := make([]int, 0)
	for _, line := range lines {
		calibration_value := make([]rune, 0)
		runes := []rune(line)
		for i := 0; i < len(runes); i += 1 {
			if isNumber(runes[i]) {
				calibration_value = append(calibration_value, runes[i])
				break
			}
		}
		for i := len(runes) - 1; i >= 0; i -= 1 {
			if isNumber(runes[i]) {
				calibration_value = append(calibration_value, runes[i])
				break
			}
		}
		calibration_values = append(calibration_values, stringToInt(string(calibration_value)))
	}
	fmt.Println(sum(calibration_values))
}

func part2() {
	lines, err := readLines("input.txt")
	if err != nil {
		panic(err)
	}
	numbers := map[string]rune{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}
	calibration_values := make([]int, 0)
	for _, line := range lines {
		calibration_value := make([]rune, 0)
		runes := []rune(line)
	L1:
		for i := 0; i < len(runes); i += 1 {
			if !isNumber(runes[i]) {
				for number := range numbers {
					if strings.HasPrefix(string(runes[i:]), number) {
						calibration_value = append(calibration_value, numbers[number])
						break L1
					}
				}
			} else {
				calibration_value = append(calibration_value, runes[i])
				break L1
			}
		}
	L2:
		for i := len(runes) - 1; i >= 0; i -= 1 {
			if !isNumber(runes[i]) {
				for number := range numbers {
					if strings.HasSuffix(string(runes[:i+1]), number) {
						calibration_value = append(calibration_value, numbers[number])
						break L2
					}
				}
			} else {
				calibration_value = append(calibration_value, runes[i])
				break L2
			}
		}
		calibration_values = append(calibration_values, stringToInt(string(calibration_value)))
	}
	fmt.Println(sum(calibration_values))
}
