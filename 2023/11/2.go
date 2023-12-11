package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	y := 0
	res := 0
	exp := 1000000
	_map := []string{}
	g, ex, ey := [][]int{}, []int{}, []int{}
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		_map = append(_map, scanner.Text())
		for x := range _map[y] {
			if string(_map[y][x]) == "#" {
				g = append(g, []int{x, y})
			}
		}
		y++
	}
	for y := range _map {
		if !strings.Contains(_map[y], "#") {
			ey = append(ey, y)
		}
	}
	for x := range _map[0] {
		v := true
		for y := range _map {
			if string(_map[y][x]) == "#" {
				v = false
				break
			}
		}
		if v {
			ex = append(ex, x)
		}
	}
	for i := range g {
		cy, cx := 0, 0
		for _, x := range ex {
			if g[i][0] > x {
				cx += exp - 1
			}
		}
		for _, y := range ey {
			if g[i][1] > y {
				cy += exp - 1
			}
		}
		g[i][0] += cx
		g[i][1] += cy
	}
	for i := range g {
		for j := i + 1; j < len(g); j++ {
			res += int(math.Abs(float64(g[i][0])-float64(g[j][0])) + math.Abs(float64(g[i][1])-float64(g[j][1])))
		}
	}
	fmt.Println(res)
}
