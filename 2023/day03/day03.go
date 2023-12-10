package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func checkCell(lines []string, x, y int) bool {
	if x < 0 || x >= 140 || y < 0 || y >= 140 {
		return false
	}

	currValue := rune(lines[x][y])
	return !unicode.IsDigit(currValue) && currValue != '.'
}

func checkSurrounds(lines []string, x, y int) bool {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			if checkCell(lines, x+i, y+j) {
				return true
			}
		}
	}

	return false
}

func main() {
	body, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("failed to read file: %v\n", err.Error())
		return
	}

	lines := strings.Split(string(body), "\n")
	sum := 0
	for x, line := range lines {
		currNum := 0
		isPart := false
		for y, ch := range line {
			if !unicode.IsDigit(ch) {
				if isPart {
					sum += currNum
				}

				currNum = 0
				isPart = false
				continue
			}

			digit := int(ch) - int('0')
			currNum = currNum*10 + digit
			if !isPart {
				isPart = checkSurrounds(lines, x, y)
			}
		}

		if isPart {
			sum += curr514969Num
		}
	}

	fmt.Println(sum)
}
