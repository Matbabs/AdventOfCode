package main

import (
	"bufio"
	"fmt"
	"os"
)

var MOVES = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func elevation(x, y int, _map []string) byte {
	if _map[y][x] == 'S' {
		return 'a'
	}
	if _map[y][x] == 'E' {
		return 'z'
	}
	return _map[y][x]
}

func bfs(xs, ys int, _map []string) (map[string]string, int, int) {
	toVisit := make([][]int, 0)
	toVisit = append(toVisit, []int{xs, ys})
	visited := make(map[string]bool)
	visited[fmt.Sprintf("%d,%d", xs, ys)] = true
	paths := make(map[string]string)
	for len(toVisit) > 0 {
		x, y := toVisit[0][0], toVisit[0][1]
		toVisit = toVisit[1:]
		for _, move := range MOVES {
			nx, ny := x+move[0], y+move[1]
			_, isVisited := visited[fmt.Sprintf("%d,%d", nx, ny)]
			if nx >= 0 && ny >= 0 && nx < len(_map[0]) && ny < len(_map) && !isVisited {
				if elevation(nx, ny, _map) <= elevation(x, y, _map)+1 {
					toVisit = append(toVisit, []int{nx, ny})
					visited[fmt.Sprintf("%d,%d", nx, ny)] = true
					paths[fmt.Sprintf("%d,%d", nx, ny)] = fmt.Sprintf("%d,%d", x, y)
					if _map[ny][nx] == 'E' {
						return paths, nx, ny
					}
				}
			}
		}
	}
	return paths, -1, -1
}

func main() {
	starts := make([][]int, 0)
	minSteps := -1
	_map := make([]string, 0)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		_map = append(_map, line)
		for x, c := range line {
			if c == 'S' || c == 'a' {
				starts = append(starts, []int{x, len(_map) - 1})
			}
		}
	}
	for _, start := range starts {
		steps := 0
		xs, ys := start[0], start[1]
		paths, xe, ye := bfs(xs, ys, _map)
		if xe != -1 && ye != -1 {
			parent, hasParent := paths[fmt.Sprintf("%d,%d", xe, ye)]
			for hasParent {
				parent, hasParent = paths[parent]
				steps++
			}
			if minSteps == -1 || steps < minSteps {
				minSteps = steps
			}
		}
	}
	fmt.Println(minSteps)
}
