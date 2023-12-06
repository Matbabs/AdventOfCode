package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var (
		res, tms, dst int
	)
	re := regexp.MustCompile("\\d+")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, _ := strconv.Atoi(strings.Join(re.FindAllString(scanner.Text(), -1), ""))
		if tms == 0 {
			tms = n
		} else {
			dst = n
		}
	}
	for t := 0; t <= tms; t++ {
		if d := (tms - t) * t; d > dst {
			res++
		}
	}
	fmt.Println(res)
}
