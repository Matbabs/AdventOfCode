package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ids := make([]string, 0)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ids = append(ids, scanner.Text())
	}
	for i := 0; i < len(ids); i++ {
		for j := i + 1; j < len(ids); j++ {
			same := ""
			for k := range ids[i] {
				if ids[i][k] == ids[j][k] {
					same += string(ids[i][k])
				}
			}
			if len(same) == len(ids[i])-1 {
				fmt.Println(same)
				return
			}
		}
	}
}
