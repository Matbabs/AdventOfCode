package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		max_cls = make([]int, 3)
		cls     = 0
	)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if cls > max_cls[2] {
				max_cls[0], max_cls[1], max_cls[2] = max_cls[1], max_cls[2], cls
			} else if cls > max_cls[1] {
				max_cls[0], max_cls[1] = max_cls[1], cls
			} else if cls > max_cls[0] {
				max_cls[0] = cls
			}
			cls = 0
			continue
		}
		kcl, _ := strconv.Atoi(line)
		cls += kcl
	}
	fmt.Println(max_cls[2] + max_cls[1] + max_cls[0])
}
