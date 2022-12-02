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
	x := 0
	y := 0
	facing := 0
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
			case 1:
				x += dist
			case 2:
				y += dist
			case 3:
				x -= dist
			}
		}
	}
	fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
}
