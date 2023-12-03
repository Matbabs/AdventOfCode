package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	res := 0
	gears := make(map[string]int)
	pairs := make(map[string]bool)
	_map := []string{}
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		_map = append(_map, scanner.Text())
	}
	for i := range _map {
		for j := 0; j < len(_map[i]); {
			if unicode.IsDigit(rune(_map[i][j])) {
				k := 0
				for j+k < len(_map[i]) && unicode.IsDigit(rune(_map[i][j+k])) {
					k++
				}
				n, _ := strconv.Atoi(_map[i][j : j+k])
				for y := i - 1; y < i+2; y++ {
					for x := j - 1; x < j+k+1; x++ {
						if x > 0 && x < len(_map[i]) && y > 0 && y < len(_map) && _map[y][x] == '*' && !unicode.IsDigit(rune(_map[y][x])) {
							key := fmt.Sprintf("%d:%d", y, x)
							_, p := pairs[key]
							if _, f := gears[key]; f && !p {
								gears[key] *= n
								pairs[key] = true
							} else {
								gears[key] = n
							}
						}
					}
				}
				j += k
			} else {
				j++
			}
		}
	}
	for k, m := range gears {
		if _, v := pairs[k]; v {
			res += m
		}
	}
	fmt.Println(res)
}
