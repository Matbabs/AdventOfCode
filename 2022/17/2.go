package main

import (
	"bufio"
	"fmt"
	"os"
)

const STEPS = 100000
const WIDTH = 7
const START_X = 2
const START_Y = 3

var shapes = [][][]int{
	{{0, 0}, {1, 0}, {2, 0}, {3, 0}},             // -
	{{1, 0}, {0, -1}, {1, -1}, {2, -1}, {1, -2}}, // +
	{{2, 0}, {2, -1}, {0, -2}, {1, -2}, {2, -2}}, // _|
	{{0, 0}, {0, -1}, {0, -2}, {0, -3}},          // |
	{{0, 0}, {1, 0}, {0, -1}, {1, -1}},           // =
}

var size_shapes = []int{1, 3, 3, 4, 2}

func main() {
	jets := ""
	steps := 0
	num_jets := 0
	max_height := -1
	sx, sy := 0, 0
	actual_shape := -1
	isMoving := false
	tower := make(map[string]bool)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		jets += scanner.Text()
	}
	for steps < STEPS {
		num_jets = num_jets % len(jets)
		if !isMoving {
			actual_shape++
			actual_shape = actual_shape % len(shapes)
			sx = START_X
			sy = START_Y + max_height + size_shapes[actual_shape]
			isMoving = true
		}
		canMove := true
		move := 0
		if jets[num_jets] == '>' {
			move = 1
		}
		if jets[num_jets] == '<' {
			move = -1
		}
		for _, block := range shapes[actual_shape] {
			px, py := sx+block[0]+move, sy+block[1]
			_, ok := tower[fmt.Sprintf("%d,%d", px, py)]
			if px < 0 || px > WIDTH-1 || ok {
				canMove = false
				break
			}
		}
		if canMove {
			sx += move
		}
		canDown := true
		for _, block := range shapes[actual_shape] {
			px, py := sx+block[0], sy+block[1]+1
			_, ok := tower[fmt.Sprintf("%d,%d", px, py-2)]
			if py <= 1 || ok {
				canDown = false
				break
			}
		}
		if canDown {
			sy--
		}
		if !canDown {
			isMoving = false
			new_max_height := max_height
			for _, block := range shapes[actual_shape] {
				px, py := sx+block[0], sy+block[1]
				tower[fmt.Sprintf("%d,%d", px, py)] = true
				if py > new_max_height {
					new_max_height = py
				}
			}
			max_height = new_max_height
			steps++
		}
		num_jets++
	}
	max_height++
	fmt.Println(max_height)
}
