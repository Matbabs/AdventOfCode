package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	score := 0
	win_table := map[string]int{"A": 2, "B": 3, "C": 1}
	loose_table := map[string]int{"A": 3, "B": 1, "C": 2}
	draw_table := map[string]int{"A": 1, "B": 2, "C": 3}
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		round := strings.Split(scanner.Text(), " ")
		elf, you := round[0], round[1]
		switch you {
		case "X":
			score += loose_table[elf]
		case "Y":
			score += 3 + draw_table[elf]
		case "Z":
			score += 6 + win_table[elf]
		}
	}
	fmt.Println(score)
}
