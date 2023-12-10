package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func extractNums(line string) []int {
	r, _ := regexp.Compile(`-?\d+`)
	numsStrings := r.FindAllString(line, -1)
	nums := make([]int, 0)

	for _, numString := range numsStrings {
		val, _ := strconv.Atoi(numString)
		nums = append(nums, val)
	}

	return nums
}

func allSame(nums []int) bool {
	if len(nums) == 0 {
		return true
	}

	first := nums[0]
	for _, num := range nums[1:] {
		if num != first {
			return false
		}
	}

	return true
}

func main() {
	body, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("failed to read file: %v\n", err.Error())
		return
	}

	sum := 0
	lines := strings.Split(strings.Trim(string(body), "\n"), "\n")
	for _, line := range lines {
		nums := extractNums(line)
		lastNums := make([]int, 0)
		diffs := nums
		fmt.Println(nums)
		lastNums = append(lastNums, nums[len(nums)-1])
		for !allSame(diffs) {
			prev := diffs
			diffs = make([]int, len(diffs)-1)
			for i := range prev[1:] {
				diffs[i] = prev[i+1] - prev[i]
			}

			fmt.Println(diffs)
			lastNums = append(lastNums, diffs[len(diffs)-1])
		}

		total := 0
		for _, num := range lastNums {
			total += num
		}

		fmt.Println(total)
		sum += total
	}

	fmt.Println(sum)
}
