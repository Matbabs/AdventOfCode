package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	visited := make(map[string]bool)
	x := 0
	y := 0
	x_visited := 0
	y_visited := 0
	facing := 0
	visit := func(x, y int) bool {
		if _, ok := visited[fmt.Sprintf("%d,%d", x, y)]; !ok {
			visited[fmt.Sprintf("%d,%d", x, y_visited)] = true
			return false
		}
		fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
		os.Exit(0)
		return true
	}
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		steps := strings.Split(scanner.Text(), ", ")
		for _, step := range steps {
			if step[0] == 'R' {
				facing = (facing + 1) % 4
			} else {
				facing = (facing + 3) % 4
			}
			dist, _ := strconv.Atoi(step[1:])
			switch facing {
			case 0:
				y -= dist
				for y_visited > y+1 {
					y_visited--
					visit(x, y_visited)
				}
			case 1:
				x += dist
				for x_visited < x-1 {
					x_visited++
					visit(x_visited, y)
				}
			case 2:
				y += dist
				for y_visited < y-1 {
					y_visited++
					visit(x, y_visited)
				}
			case 3:
				x -= dist
				for x_visited > x+1 {
					x_visited--
					visit(x_visited, y)
				}
			}
			visit(x, y)
			x_visited = x
			y_visited = y
		}
	}
}
