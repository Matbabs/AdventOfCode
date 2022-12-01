package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var bags_insiders = make(map[string]map[string]int64)

func visit(key string) int64 {
	count := int64(0)
	for i := range bags_insiders[key] {
		count += bags_insiders[key][i] + bags_insiders[key][i]*visit(i)
	}
	return count
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		container := strings.Join(words[0:2], " ")
		for i := 4; i < len(words); i++ {
			if words[i][0] >= '0' && words[i][0] <= '9' {
				child := strings.Join(words[i+1:i+3], " ")
				if _, ok := bags_insiders[container]; !ok {
					bags_insiders[container] = make(map[string]int64)
				}
				bags_insiders[container][child] = int64(words[i][0] - '0')
			}
		}
	}
	fmt.Println(visit("shiny gold"))
}
