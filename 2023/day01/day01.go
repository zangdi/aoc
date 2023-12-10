package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	body, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("failed to read file: %v\n", err.Error())
		return
	}

	lines := strings.Split(string(body), "\n")

	fmt.Printf("length of lines: %v\n", len(lines))

	sum := 0
	first := -1
	last := 0
	for _, line := range lines {
		for _, c := range line {
			if unicode.IsDigit(c) {
				last = int(c) - int('0')
				if first == -1 {
					first = last
				}
			}
		}

		if first == -1 {
			continue
		}

		currNum := 10*first + last
		fmt.Println(currNum)
		sum += currNum
		first = -1
	}

	fmt.Println(sum)
}
