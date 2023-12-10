package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	body, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("failed to read file: %v\n", err.Error())
		return
	}

	lines := strings.Split(strings.Trim(string(body), "\n"), "\n")
	sum := 0
	r, _ := regexp.Compile(`\d+`)
	for _, line := range lines {
		line = line[8:]
		cards := strings.Split(line, "|")

		winningsStr := r.FindAllString(cards[0], -1)
		winnings := map[string]struct{}{}
		for _, number := range winningsStr {
			winnings[number] = struct{}{}
		}

		currValue := 0
		cardStr := r.FindAllString(cards[1], -1)
		for _, number := range cardStr {
			if _, ok := winnings[number]; ok {
				if currValue == 0 {
					currValue = 1
				} else {
					currValue *= 2
				}
			}
		}

		fmt.Println(currValue)
		sum += currValue
	}

	fmt.Println(sum)
}
