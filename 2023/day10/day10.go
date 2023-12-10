package main

import (
	"fmt"
	"os"
	"strings"
)

type Coordinates struct {
	x int
	y int
}

var correctDirections map[string]string

func init() {
	correctDirections = map[string]string{
		"north": "SJL|",
		"south": "S|7F",
		"east":  "SJ7-",
		"west":  "SLF-",
	}
}

func oppositeDirection(dir string) string {
	if dir == "north" {
		return "south"
	} else if dir == "south" {
		return "north"
	} else if dir == "east" {
		return "west"
	} else {
		return "east"
	}
}

func validDirection(pipe byte, dir string) bool {
	if len(dir) == 0 {
		return true
	}

	return strings.Contains(correctDirections[oppositeDirection(dir)], string(pipe))
}

func dfs(graph []string, pos Coordinates, visited map[Coordinates]struct{}, dir string) (int, bool) {
	if pos.x < 0 || pos.x >= len(graph) || pos.y < 0 || pos.y >= len(graph[0]) || graph[pos.x][pos.y] == '.' {
		return 0, false
	}

	if !validDirection(graph[pos.x][pos.y], dir) {
		return 0, false
	}

	if _, ok := visited[pos]; ok {
		if graph[pos.x][pos.y] == 'S' {
			fmt.Printf("0: %v S\n", pos)
		}
		return 0, graph[pos.x][pos.y] == 'S'
	}

	visited[pos] = struct{}{}
	north := Coordinates{pos.x - 1, pos.y}
	south := Coordinates{pos.x + 1, pos.y}
	east := Coordinates{pos.x, pos.y - 1}
	west := Coordinates{pos.x, pos.y + 1}

	length := 0
	found := false
	if strings.Contains("SJL|", string(graph[pos.x][pos.y])) && dir != "south" {
		length, found = dfs(graph, north, visited, "north")
		if found {
			fmt.Printf("%v: %v %v\n", 1+length, pos, string(graph[pos.x][pos.y]))
		}

		if found || graph[pos.x][pos.y] != 'S' {
			return 1 + length, found
		}
	}

	if strings.Contains("S7F|", string(graph[pos.x][pos.y])) && dir != "north" {
		length, found = dfs(graph, south, visited, "south")
		if found {
			fmt.Printf("%v: %v %v\n", 1+length, pos, string(graph[pos.x][pos.y]))
		}

		if found || graph[pos.x][pos.y] != 'S' {
			return 1 + length, found
		}
	}

	if strings.Contains("SJ7-", string(graph[pos.x][pos.y])) && dir != "west" {
		length, found = dfs(graph, east, visited, "east")
		if found {
			fmt.Printf("%v: %v %v\n", 1+length, pos, string(graph[pos.x][pos.y]))
		}

		if found || graph[pos.x][pos.y] != 'S' {
			return 1 + length, found
		}
	}

	if strings.Contains("SLF-", string(graph[pos.x][pos.y])) && dir != "east" {
		length, found = dfs(graph, west, visited, "west")
		if found {
			fmt.Printf("%v: %v %v\n", 1+length, pos, string(graph[pos.x][pos.y]))
		}

		if found || graph[pos.x][pos.y] != 'S' {
			return 1 + length, found
		}
	}

	return 0, false
}

func main() {
	body, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("failed to read file: %v\n", err.Error())
		return
	}

	lines := strings.Split(strings.Trim(string(body), "\n"), "\n")
	start := Coordinates{0, 0}
	for i, line := range lines {
		if j := strings.Index(line, "S"); j != -1 {
			start = Coordinates{i, j}
			break
		}
	}

	visited := map[Coordinates]struct{}{}
	length, _ := dfs(lines, start, visited, "")

	fmt.Println(length / 2)
}
