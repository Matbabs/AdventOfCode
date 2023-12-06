package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	res := 1
	tms := make([]int, 0)
	dst := make([]int, 0)
	re := regexp.MustCompile("\\d+")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ns := []int{}
		nbs := re.FindAllString(scanner.Text(), -1)
		ns = make([]int, len(nbs))
		for i, nb := range nbs {
			ns[i], _ = strconv.Atoi(nb)
		}
		if len(tms) == 0 {
			tms = ns
		} else {
			dst = ns
		}
	}
	for r := 0; r < len(tms); r++ {
		subRes := 0
		for t := 0; t <= tms[r]; t++ {
			if d := (tms[r] - t) * t; d > dst[r] {
				subRes++
			}
		}
		res *= subRes
	}
	fmt.Println(res)
}
