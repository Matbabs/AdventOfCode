package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	code := ""
	x := 0
	y := 0
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		op := scanner.Text()
		for _, c := range op {
			switch c {
			case 'U':
				if y > 0 {
					y--
				}
			case 'D':
				if y < 2 {
					y++
				}
			case 'L':
				if x > 0 {
					x--
				}
			case 'R':
				if x < 2 {
					x++
				}
			}
		}
		code += fmt.Sprintf("%d", y*3+x+1)
	}
	fmt.Println(code)
}
