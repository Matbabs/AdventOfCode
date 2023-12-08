package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	res := 0
	a := "AAA"
	p := ""
	l := 0
	n := make(map[string][]string)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if l == 0 {
			p = scanner.Text()
			l++
		} else if scanner.Text() != "" {
			nd := strings.Split(strings.ReplaceAll(scanner.Text(), " ", ""), "=")
			pp := strings.Split(nd[1], ",")
			n[nd[0]] = []string{pp[0][1:], pp[1][:3]}
		}
	}
	for a != "ZZZ" {
		for _, r := range p {
			a = n[a][strings.IndexRune("LR", r)]
			res++
			if a == "ZZZ" {
				break
			}
		}
	}
	fmt.Println(res)
}
