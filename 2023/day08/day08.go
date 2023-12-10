package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type graphNode struct {
	name  string
	left  *graphNode
	right *graphNode
}

func getOrCreate(name string, nodeMap map[string]*graphNode) *graphNode {
	if _, ok := nodeMap[name]; !ok {
		nodeMap[name] = &graphNode{name: name, left: nil, right: nil}
	}

	return nodeMap[name]
}

func main1() {
	body, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("failed to read file: %v\n", err.Error())
		return
	}

	r, _ := regexp.Compile(`\w+`)
	lines := strings.Split(strings.Trim(string(body), "\n"), "\n")
	instructions := lines[0]
	nodeMap := map[string]*graphNode{}
	for _, line := range lines[2:] {
		components := r.FindAllString(line, -1)
		node := getOrCreate(components[0], nodeMap)
		left := getOrCreate(components[1], nodeMap)
		right := getOrCreate(components[2], nodeMap)

		node.left = left
		node.right = right
	}

	curr, _ := nodeMap["AAA"]
	steps := 0
	for ; curr.name != "ZZZ"; steps++ {
		idx := steps % len(instructions)
		fmt.Printf("at %v, steps: %v,  index: %v, ", curr.name, steps, idx)
		if instructions[idx] == 'L' {
			curr = curr.left
			fmt.Printf("left to: %v\n", curr)
		} else {
			curr = curr.right
			fmt.Printf("right to: %v\n", curr)
		}
	}

	fmt.Println(steps)
}

func allNodesEndWithZ(currNodes []*graphNode) bool {
	for _, node := range currNodes {
		if node.name[2] != 'Z' {
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

	r, _ := regexp.Compile(`\w+`)
	lines := strings.Split(strings.Trim(string(body), "\n"), "\n")
	instructions := lines[0]
	nodeMap := map[string]*graphNode{}
	currNodes := make([]*graphNode, 0, 6)
	for _, line := range lines[2:] {
		components := r.FindAllString(line, -1)
		node := getOrCreate(components[0], nodeMap)
		left := getOrCreate(components[1], nodeMap)
		right := getOrCreate(components[2], nodeMap)

		node.left = left
		node.right = right

		if node.name[2] == 'A' {
			currNodes = append(currNodes, node)
		}
	}

	steps := 0
	for ; !allNodesEndWithZ(currNodes); steps++ {
		idx := steps % len(instructions)
		if instructions[idx] == 'L' {
			for i, node := range currNodes {
				currNodes[i] = node.left
			}
		} else {
			for i, node := range currNodes {
				currNodes[i] = node.right
			}
		}
	}

	fmt.Println(steps)
}
