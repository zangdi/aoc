package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	body, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("failed to read file: %v\n", err.Error())
		return
	}

	bagSet := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	lines := strings.Split(strings.Trim(string(body), "\n"), "\n")
	sum := 0
	for _, line := range lines {
		line = line[5:]
		colonSeparated := strings.Split(line, ": ")
		gameId, _ := strconv.Atoi(colonSeparated[0])
		ballSets := strings.Split(colonSeparated[1], "; ")
		isValid := true
		for _, ballSet := range ballSets {
			balls := strings.Split(ballSet, ", ")
			for _, ball := range balls {
				spaceSeparated := strings.Split(ball, " ")
				amount, _ := strconv.Atoi(spaceSeparated[0])
				if amount > bagSet[spaceSeparated[1]] {
					isValid = false
					break
				}
			}

			if !isValid {
				break
			}
		}

		if isValid {
			sum += gameId
		}
	}

	fmt.Println(sum)
}
