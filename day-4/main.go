package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

func scanFile(filename string) []string {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text
}

func main() {
	fileContents := scanFile("in.txt")

	var winnings []int

	for _, each_ln := range fileContents {
		card := strings.Split(each_ln, ": ")[1]
		parts := strings.Split(card, " | ")

		winning := strings.Split(parts[0], " ")
		chosen := strings.Split(parts[1], " ")

		winning_set := mapset.NewSet[int]()
		chosen_set := mapset.NewSet[int]()

		for _, w := range winning {
			if w != "" {
				w_i, _ := strconv.Atoi(w)
				winning_set.Add(w_i)
			}
		}

		for _, c := range chosen {
			if c != "" {
				c_i, _ := strconv.Atoi(c)
				chosen_set.Add(c_i)
			}
		}

		intersect := chosen_set.Intersect(winning_set)
		numIntersect := intersect.Cardinality()

		winnings = append(winnings, numIntersect)
	}

	numCards := make([]int, len(winnings))
	for i := range numCards {
		numCards[i] = 1
	}

	for i := range numCards {
		numWon := winnings[i]
		numCurr := numCards[i]

		endIndex := numWon + i + 1

		if endIndex > len(winnings) {
			endIndex = len(winnings)
		}

		for j := i + 1; j < endIndex; j++ {
			numCards[j] += numCurr
		}
	}

	sum := 0
	for _, a := range numCards {
		sum += a
	}

	fmt.Println(winnings)
	fmt.Println(numCards)
	fmt.Println(sum)
}
