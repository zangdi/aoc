package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type wager struct {
	cards    string
	value    int
	strength int
}

func byteToVal(b byte) int {
	switch b {
	case 'A':
		return 13
	case 'K':
		return 12
	case 'Q':
		return 11
	case 'J':
		return 10
	case 'T':
		return 9
	default:
		num, _ := strconv.Atoi(string(b))
		return num - 1
	}
}

func main() {
	body, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("failed to read file: %v\n", err.Error())
		return
	}

	lines := strings.Split(strings.Trim(string(body), "\n"), "\n")
	wagers := make([]wager, 0)
	for _, line := range lines {
		components := strings.Split(line, " ")
		counts := map[rune]int{}

		for _, c := range components[0] {
			counts[c] += 1
		}

		sets := map[int]int{}
		for _, v := range counts {
			sets[v] += 1
		}

		strength := 0
		if sets[5] == 1 {
			strength = 6
		} else if sets[4] == 1 {
			strength = 5
		} else if sets[3] == 1 && sets[2] == 1 {
			strength = 4
		} else if sets[3] == 1 {
			strength = 3
		} else if sets[2] == 2 {
			strength = 2
		} else if sets[2] == 1 {
			strength = 1
		}

		value, _ := strconv.Atoi(components[1])
		wagers = append(wagers, wager{cards: components[0], value: value, strength: strength})
	}

	sort.SliceStable(wagers, func(i, j int) bool {
		a := wagers[i]
		b := wagers[j]

		if a.strength != b.strength {
			return a.strength < b.strength
		}

		for k := 0; k < 5; k++ {
			if a.cards[k] != b.cards[k] {
				return byteToVal(a.cards[k]) < byteToVal(b.cards[k])
			}
		}

		return true
	})

	fmt.Println(wagers)

	sum := 0
	for i, val := range wagers {
		sum += (i + 1) * val.value
	}

	fmt.Println(sum)
}
