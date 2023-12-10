package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func extractNums(line string) []int {
	r, _ := regexp.Compile(`\d+`)
	numsStrings := r.FindAllString(line, -1)
	nums := make([]int, 0)

	for _, numString := range numsStrings {
		val, _ := strconv.Atoi(numString)
		nums = append(nums, val)
	}

	return nums
}

func calculateDistance(timeHeld, totalTime int) int {
	return timeHeld * (totalTime - timeHeld)
}

func binarySearchImpl(start, end, total, distance int, left bool) int {
	fmt.Printf("start: %v, end: %v, left: %v\n", start, end, left)

	if start >= end {
		if left {
			return start
		} else {
			return end
		}
	}

	if end-start == 1 {
		if left {
			if calculateDistance(start, total) > distance {
				return start
			} else {
				return end
			}
		} else {
			if calculateDistance(end, total) > distance {
				return end
			} else {
				return start
			}
		}
	}

	mid := (start + end) / 2
	val := calculateDistance(mid, total)

	if val == distance+1 {
		return mid
	}

	if val == distance {
		if left {
			return mid + 1
		} else {
			return mid - 1
		}
	}

	if val > distance {
		if left {
			return binarySearchImpl(start, mid, total, distance, left)
		} else {
			return binarySearchImpl(mid, end, total, distance, left)
		}
	} else {
		if left {
			return binarySearchImpl(mid, end, total, distance, left)
		} else {
			return binarySearchImpl(start, mid, total, distance, left)
		}
	}
}

func binarySearch(time, distance int) int {
	mid := time / 2
	if calculateDistance(mid, time) < distance {
		return 0
	}

	left := binarySearchImpl(0, mid, time, distance, true)
	right := binarySearchImpl(mid, time, time, distance, false)

	fmt.Printf("left: %v, right: %v\n", left, right)

	return right + 1 - left
}

func main() {
	body, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("failed to read file: %v\n", err.Error())
		return
	}

	lines := strings.Split(strings.Trim(string(body), "\n"), "\n")
	times := extractNums(strings.ReplaceAll(lines[0], " ", ""))
	distances := extractNums(strings.ReplaceAll(lines[1], " ", ""))

	total := 1
	for i := 0; i < len(times); i++ {
		currVal := binarySearch(times[i], distances[i])
		total *= currVal
	}

	fmt.Println(total)
}
