package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := 0
	_map := make([][]int, 0)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := make([]int, 0)
		for _, c := range scanner.Text() {
			row = append(row, int(c-'0'))
		}
		_map = append(_map, row)
	}
	for i := 0; i < len(_map); i++ {
		for j := 0; j < len(_map[i]); j++ {
			if i == 0 || j == 0 || i == len(_map)-1 || j == len(_map[i])-1 {
				count++
			} else {
				t := true
				for d := i - 1; d >= 0; d-- {
					if _map[d][j] >= _map[i][j] {
						t = false
						break
					}
				}
				r := true
				for d := j + 1; d < len(_map[i]); d++ {
					if _map[i][d] >= _map[i][j] {
						r = false
						break
					}
				}
				b := true
				for d := i + 1; d < len(_map); d++ {
					if _map[d][j] >= _map[i][j] {
						b = false
						break
					}
				}
				l := true
				for d := j - 1; d >= 0; d-- {
					if _map[i][d] >= _map[i][j] {
						l = false
						break
					}
				}
				if t || r || b || l {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
