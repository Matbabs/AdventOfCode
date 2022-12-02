package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	ribbon := 0
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		surfaces := strings.Split(scanner.Text(), "x")
		l, _ := strconv.Atoi(surfaces[0])
		w, _ := strconv.Atoi(surfaces[1])
		h, _ := strconv.Atoi(surfaces[2])
		sides := []int{l, w, h}
		sort.Ints(sides)
		ribbon += 2*sides[0] + 2*sides[1] + l*w*h
	}
	fmt.Println(ribbon)
}
