package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	panoramic := 0
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
			if !(i == 0 || j == 0 || i == len(_map)-1 || j == len(_map[i])-1) {
				score := 0
				t, r, b, l := 0, 0, 0, 0
				for d := i - 1; d >= 0; d-- {
					t++
					if _map[d][j] >= _map[i][j] {
						break
					}
				}
				for d := j + 1; d < len(_map[i]); d++ {
					r++
					if _map[i][d] >= _map[i][j] {
						break
					}
				}
				for d := i + 1; d < len(_map); d++ {
					b++
					if _map[d][j] >= _map[i][j] {
						break
					}
				}
				for d := j - 1; d >= 0; d-- {
					l++
					if _map[i][d] >= _map[i][j] {
						break
					}
				}
				score = t * r * b * l
				if score > panoramic {
					panoramic = score
				}
			}
		}
	}
	fmt.Println(panoramic)
}
