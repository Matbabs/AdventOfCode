package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	size_total := 0
	check_size := 100000
	pwd := make([]string, 0)
	dirs_size := make(map[string]int)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		log := strings.Split(scanner.Text(), " ")
		if log[1] == "cd" && log[2] != ".." {
			pwd = append(pwd, log[2])
		} else if log[1] == "cd" && log[2] == ".." {
			pwd = pwd[:len(pwd)-1]
		} else {
			size, err := strconv.Atoi(log[0])
			if err == nil {
				path := ""
				for _, dir := range pwd {
					path += dir + "/"
					dirs_size[path] += size
				}
			}
		}
	}
	for dir, size := range dirs_size {
		if dir != "/" && size <= check_size {
			size_total += size
		}
	}
	fmt.Println(size_total)
}
