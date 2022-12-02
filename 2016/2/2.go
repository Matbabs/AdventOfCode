package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	maps := [][]rune{
		[]rune{' ', ' ', '1', ' ', ' '},
		[]rune{' ', '2', '3', '4', ' '},
		[]rune{'5', '6', '7', '8', '9'},
		[]rune{' ', 'A', 'B', 'C', ' '},
		[]rune{' ', ' ', 'D', ' ', ' '},
	}
	code := ""
	x := 0
	y := 2
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		op := scanner.Text()
		for _, c := range op {
			switch c {
			case 'U':
				if y > 0 && maps[y-1][x] != ' ' {
					y--
				}
			case 'D':
				if y < 4 && maps[y+1][x] != ' ' {
					y++
				}
			case 'L':
				if x > 0 && maps[y][x-1] != ' ' {
					x--
				}
			case 'R':
				if x < 4 && maps[y][x+1] != ' ' {
					x++
				}
			}
		}
		code += fmt.Sprintf("%c", maps[y][x])
	}
	fmt.Println(code)
}
