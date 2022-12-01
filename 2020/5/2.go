package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	id_max := 0
	seats := make(map[int]bool)
	for scanner.Scan() {
		line := scanner.Text()
		row_min := 0
		row_max := 127
		col_min := 0
		col_max := 7
		for _, c := range line {
			switch c {
			case 'F':
				row_max = (row_max + row_min) / 2
			case 'B':
				row_min = (row_max + row_min) / 2
			case 'L':
				col_max = (col_max + col_min) / 2
			case 'R':
				col_min = (col_max + col_min) / 2
			}
		}
		seat_id := row_max*8 + col_max
		seats[seat_id] = true
		if seat_id > id_max {
			id_max = seat_id
		}
	}
	for i := 0; i < id_max; i++ {
		if !seats[i] && seats[i-1] && seats[i+1] {
			fmt.Println(i)
			break
		}
	}
}
