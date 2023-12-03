package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"unicode"
)

// y increases downward
type number struct {
	value      int
	startX     int
	stopX      int
	y          int
	isTouching bool
}

type symbol struct {
	x            int
	y            int
	isStar       bool
	numAdjacent  int
	adjacentProd int64
}

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

func getNumbersInLine(lineNum int, line string) ([]number, []symbol) {
	var numbers []number
	var symbols []symbol

	currentNumber := ""
	numberStart := math.MaxInt

	for i, l := range line {
		if unicode.IsDigit(l) {
			if i < numberStart {
				numberStart = i
			}

			currentNumber = currentNumber + string(l)
		} else if numberStart != math.MaxInt {
			numberEnd := i
			numValue, _ := strconv.Atoi(line[numberStart:numberEnd])

			newNumber := number{startX: numberStart, stopX: numberEnd, y: lineNum, value: numValue, isTouching: false}
			numbers = append(numbers, newNumber)

			numberStart = math.MaxInt
			currentNumber = ""
		}

		if l == '*' && !unicode.IsDigit(l) {
			newSymbol := symbol{x: i, y: lineNum, isStar: true, numAdjacent: 0, adjacentProd: 1}
			symbols = append(symbols, newSymbol)
		} else if l != '.' && !unicode.IsDigit(l) {
			newSymbol := symbol{x: i, y: lineNum, isStar: false, numAdjacent: 0, adjacentProd: 1}
			symbols = append(symbols, newSymbol)
		}
	}

	if numberStart != math.MaxInt {
		numberEnd := len(line)
		numValue, _ := strconv.Atoi(line[numberStart:numberEnd])

		newNumber := number{startX: numberStart, stopX: numberEnd, y: lineNum, value: numValue, isTouching: false}

		numbers = append(numbers, newNumber)
	}

	return numbers, symbols
}

func checkTouching(num *number, symbols []symbol) bool {
	for i, sym := range symbols {
		numStart := num.startX
		numEnd := num.stopX
		numY := num.y

		if sym.x >= numStart-1 && sym.x <= numEnd && sym.y <= numY+1 && sym.y >= numY-1 {
			symbols[i].numAdjacent += 1
			symbols[i].adjacentProd *= int64(num.value)

			return true
		}
	}
	return false
}

func main() {
	fileContents := scanFile("in.txt")

	var allNumbers []number
	var allSymbols []symbol

	for i, each_ln := range fileContents {
		numbersInLine, symbolsInLine := getNumbersInLine(i, each_ln)

		allNumbers = append(allNumbers, numbersInLine...)
		allSymbols = append(allSymbols, symbolsInLine...)
	}

	for i, num := range allNumbers {
		touching := checkTouching(&num, allSymbols)
		allNumbers[i].isTouching = touching
	}

	sum := 0
	for _, num := range allNumbers {
		if num.isTouching {
			sum = sum + num.value
		}
	}

	fmt.Println(sum)

	sum = 0

	for _, sym := range allSymbols {
		if sym.numAdjacent == 2 && sym.isStar {
			sum += int(sym.adjacentProd)
		}
	}

	fmt.Println(sum)
}
