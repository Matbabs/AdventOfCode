package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	score := 0
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		round := strings.Split(scanner.Text(), " ")
		elf, you := round[0], round[1]
		if you == "X" && elf == "A" || you == "Y" && elf == "B" || you == "Z" && elf == "C" {
			score += 3
		} else if you == "X" && elf == "C" || you == "Y" && elf == "A" || you == "Z" && elf == "B" {
			score += 6
		}
		switch you {
		case "X":
			score += 1
		case "Y":
			score += 2
		case "Z":
			score += 3
		}
	}
	fmt.Println(score)
}
