package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
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

func replaceImLazy(line string) string {
	line = strings.Replace(line, "one", "one1one", -1)
	line = strings.Replace(line, "two", "two2two", -1)
	line = strings.Replace(line, "three", "three3three", -1)
	line = strings.Replace(line, "four", "four4four", -1)
	line = strings.Replace(line, "five", "five5five", -1)
	line = strings.Replace(line, "six", "six6six", -1)
	line = strings.Replace(line, "seven", "seven7seven", -1)
	line = strings.Replace(line, "eight", "eight8eight", -1)
	line = strings.Replace(line, "nine", "nine9nine", -1)

	return line
}

func main() {
	fileContents := scanFile("in.txt")

	sum := 0

	for _, each_ln := range fileContents {
		each_ln := replaceImLazy(each_ln)

		firstDigit := ""
		lastDigit := ""

		for _, l := range each_ln {
			if unicode.IsDigit(l) {
				l := string(l)

				if firstDigit == "" {
					firstDigit = l
				}

				lastDigit = l
			}
		}

		numString := firstDigit + lastDigit
		num, _ := strconv.Atoi(numString)

		sum = sum + num
	}

	fmt.Println(sum)
}
