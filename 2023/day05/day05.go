package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/emirpasic/gods/maps/treemap"
)

type destRange struct {
	dest int
	rang int
}

func main() {
	body, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("failed to read file: %v\n", err.Error())
		return
	}

	lines := strings.Split(strings.Trim(string(body), "\n"), "\n")
	r, _ := regexp.Compile(`\d+`)
	seeds := r.FindAllString(lines[0], -1)

	converters := make([]*treemap.Map, 0)
	currMap := treemap.NewWithIntComparator()
	for _, line := range lines[3:] {
		if len(line) == 0 {
			converters = append(converters, currMap)
			currMap = treemap.NewWithIntComparator()
		} else if unicode.IsDigit(rune(line[0])) {
			vals := r.FindAllString(line, -1)
			dest, _ := strconv.Atoi(vals[0])
			source, _ := strconv.Atoi(vals[1])
			rang, _ := strconv.Atoi(vals[2])
			currMap.Put(source, destRange{dest: dest, rang: rang})
		}
	}

	converters = append(converters, currMap)

	min := -1
	for i := 0; i < len(seeds); i += 2 {
		seedStart, _ := strconv.Atoi(seeds[i])
		seedCount, _ := strconv.Atoi(seeds[i+1])
		// currVal, _ := strconv.Atoi(seed)
		for j := 0; j < seedCount; j++ {
			currVal := seedStart + j
			for _, m := range converters {
				k, v := m.Floor(currVal)
				dr, _ := v.(destRange)
				if k != nil {
					sc, _ := k.(int)
					if currVal < sc+dr.rang {
						currVal = dr.dest + currVal - sc
					}
				}
			}

			if min == -1 || currVal < min {
				min = currVal
			}
		}
	}

	fmt.Println(min)
}
