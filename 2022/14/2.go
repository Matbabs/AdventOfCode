package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const SAND_X = 500
const SAND_Y = 0
const FLOOR = 200

func main() {
	depth := 0
	_map := make(map[string]rune)
	_map[fmt.Sprintf("%d,%d", SAND_X, SAND_Y)] = '+'
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		cols := strings.Split(row, " -> ")
		for i := 0; i < len(cols)-1; i++ {
			xy1 := strings.Split(cols[i], ",")
			xy2 := strings.Split(cols[i+1], ",")
			x1, _ := strconv.Atoi(xy1[0])
			y1, _ := strconv.Atoi(xy1[1])
			x2, _ := strconv.Atoi(xy2[0])
			y2, _ := strconv.Atoi(xy2[1])
			if x1 == x2 {
				if y1 > y2 {
					y1, y2 = y2, y1
				}
				for y := y1; y <= y2; y++ {
					_map[fmt.Sprintf("%d,%d", x1, y)] = '#'
				}
			} else if y1 == y2 {
				if x1 > x2 {
					x1, x2 = x2, x1
				}
				for x := x1; x <= x2; x++ {
					_map[fmt.Sprintf("%d,%d", x, y1)] = '#'
				}
			}
			if depth < y2 {
				depth = y2
			}
		}
	}
	depth += 2
	for x := SAND_X - FLOOR; x < SAND_X+FLOOR; x++ {
		_map[fmt.Sprintf("%d,%d", x, depth)] = '#'
	}
	steps := 0
	isReachedTop := false
	for !isReachedTop {
		i := SAND_Y
		j := SAND_X
		isImmobilized := false
		for !isImmobilized {
			for _map[fmt.Sprintf("%d,%d", j, i+1)] != '#' && i < depth {
				i++
			}
			_, left := _map[fmt.Sprintf("%d,%d", j-1, i+1)]
			_, right := _map[fmt.Sprintf("%d,%d", j+1, i+1)]
			if !left {
				j--
				continue
			}
			if !right {
				j++
				continue
			}
			_map[fmt.Sprintf("%d,%d", j, i)] = '#'
			isImmobilized = true
			steps++
			if i == SAND_Y && j == SAND_X {
				isReachedTop = true
				break
			}
		}
	}
	fmt.Println(steps)
}
