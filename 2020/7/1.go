package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var bags_containers = make(map[string][]string)
var visited = make(map[string]bool)

func visit(key string) {
	for _, v := range bags_containers[key] {
		if !visited[v] {
			visited[v] = true
			if _, ok := bags_containers[v]; ok {
				visit(v)
			}
		}
	}
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
				bags_containers[child] = append(bags_containers[child], container)
			}

		}
	}
	visit("shiny gold")
	fmt.Println(len(visited))
}
