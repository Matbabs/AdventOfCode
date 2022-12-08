package main

import (
	"bufio"
	"fmt"
	"os"
)

const GREEN = "\033[32m"

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
				tv, tr, tb, tl := true, true, true, true
				for d := i - 1; d >= 0; d-- {
					t++
					if _map[d][j] >= _map[i][j] {
						tv = false
						break
					}
				}
				for d := j + 1; d < len(_map[i]); d++ {
					r++
					if _map[i][d] >= _map[i][j] {
						tr = false
						break
					}
				}
				for d := i + 1; d < len(_map); d++ {
					b++
					if _map[d][j] >= _map[i][j] {
						tb = false
						break
					}
				}
				for d := j - 1; d >= 0; d-- {
					l++
					if _map[i][d] >= _map[i][j] {
						tl = false
						break
					}
				}
				score = t * r * b * l
				if score > panoramic {
					panoramic = score
				}
				if tv || tr || tb || tl {
					if score > 100 {
						fmt.Print(GREEN, "@")
					} else if score > 75 {
						fmt.Print(GREEN, "0")
					} else if score > 50 {
						fmt.Print(GREEN, "O")
					} else if score > 25 {
						fmt.Print(GREEN, "o")
					} else {
						fmt.Print(GREEN, "°")
					}
				} else {
					fmt.Print(".")
				}
			} else {
				fmt.Print(GREEN, "°")
			}
		}
		fmt.Println()
	}
}
