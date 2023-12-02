package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func parseColours(subsetsStr string) [][3]int {
	var coloursPerLine [][3]int

	split := strings.Split(subsetsStr, ";")

	for _, subs := range split {
		redCubes := 0
		greenCubes := 0
		blueCubes := 0

		colours := strings.Split(subs, ",")

		for _, col := range colours {
			col = strings.TrimSpace(col)
			spl := strings.Split(col, " ")

			num, _ := strconv.Atoi(spl[0])
			colour := spl[1]

			if colour == "red" {
				redCubes = num
			} else if colour == "green" {
				greenCubes = num
			} else {
				blueCubes = num
			}
		}

		coloursPerLine = append(coloursPerLine, [3]int{redCubes, greenCubes, blueCubes})
	}

	return coloursPerLine
}

func max(colours [][3]int) (int, int, int) {

	mRed := 0
	mGreen := 0
	mBlue := 0

	for _, tupl := range colours {
		if tupl[0] > mRed {
			mRed = tupl[0]
		}

		if tupl[1] > mGreen {
			mGreen = tupl[1]
		}

		if tupl[2] > mBlue {
			mBlue = tupl[2]
		}
	}

	return mRed, mGreen, mBlue
}

func main() {
	fileContents := scanFile("in.txt")

	sum := 0

	for _, each_ln := range fileContents {
		spl := strings.Split(each_ln, ":")

		game := spl[0]
		game = strings.Split(game, " ")[1]
		// game_id, _ := strconv.Atoi(game)

		colours := parseColours(spl[1])
		red, green, blue := max(colours)

		// if red < 13 && green < 14 && blue < 15 {
		// 	sum = sum + game_id
		// }

		sum = sum + red*green*blue
	}
	fmt.Println(sum)
}
