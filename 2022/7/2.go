package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	disk := 70000000
	need := 30000000
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
	min_size := dirs_size["//"]
	check_size := need - (disk - dirs_size["//"])
	for dir, size := range dirs_size {
		if dir != "/" && size >= check_size {
			if size < min_size {
				min_size = size
			}
		}
	}
	fmt.Println(min_size)
}
