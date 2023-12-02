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
	maxRed, maxGreen, maxBlue := 12, 13, 14
	possibleIDs := make([]int, 0)
	for _, game := range lines {
		parts := strings.Split(game, ": ")
		gameID := stringToInt(strings.Split(parts[0], " ")[1])
		possible := true
		for _, round := range strings.Split(parts[1], "; ") {
			red, green, blue := 0, 0, 0
			cubes := strings.Split(round, ", ")
			for _, cube := range cubes {
				cubeParts := strings.Split(cube, " ")
				switch cubeParts[1] {
				case "red":
					red += stringToInt(cubeParts[0])
				case "green":
					green += stringToInt(cubeParts[0])
				case "blue":
					blue += stringToInt(cubeParts[0])
				}
			}
			if red > maxRed || green > maxGreen || blue > maxBlue {
				possible = false
				break
			}
		}
		if possible {
			possibleIDs = append(possibleIDs, gameID)
		}
	}
	fmt.Println(sum(possibleIDs))
}

func part2() {
	lines, err := readLines("input.txt")
	if err != nil {
		panic(err)
	}
	powers := make([]int, 0)
	for _, game := range lines {
		minRed, minGreen, minBlue := 0,0,0
		for _, round := range strings.Split(strings.Split(game, ": ")[1], "; ") {
			cubes := strings.Split(round, ", ")
			for _, cube := range cubes {
				cubeParts := strings.Split(cube, " ")
				switch cubeParts[1] {
				case "red":
					red := stringToInt(cubeParts[0])
					if red >minRed {
						minRed = red
					}
				case "green":
					green := stringToInt(cubeParts[0])
					if green >minGreen {
						minGreen = green
					}
				case "blue":
					blue := stringToInt(cubeParts[0])
					if blue >minBlue {
						minBlue = blue
					}
				}
			}
		}
		powers = append(powers, minRed*minGreen*minBlue)
	}
	fmt.Println(sum(powers))
}
