package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	res := 0
	re := regexp.MustCompile("\\d")
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nbs := re.FindAllString(line, -1)
		ns, _ := strconv.Atoi(nbs[0] + nbs[len(nbs)-1])
		res += ns
	}
	fmt.Println(res)
}
