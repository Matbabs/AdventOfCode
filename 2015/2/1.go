package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	surface := 0
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		surfaces := strings.Split(scanner.Text(), "x")
		l, _ := strconv.Atoi(surfaces[0])
		w, _ := strconv.Atoi(surfaces[1])
		h, _ := strconv.Atoi(surfaces[2])
		sides := []int{l * w, w * h, h * l}
		min := sides[0]
		for _, side := range sides {
			if side < min {
				min = side
			}
		}
		surface += 2*sides[0] + 2*sides[1] + 2*sides[2] + min
	}
	fmt.Println(surface)
}
