package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func hash(s string) int {
	v := 0
	for _, r := range s {
		v += int(r)
		v *= 17
		v %= 256
	}
	return v
}

func main() {
	res := 0
	seq := []string{}
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		seq = strings.Split(scanner.Text(), ",")
	}
	for _, s := range seq {
		res += hash(s)
	}
	fmt.Println(res)
}
