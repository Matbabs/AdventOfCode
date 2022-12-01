package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	fields_detected := make(map[string]bool)
	valid := 0

	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	validate := func() {
		if len(fields_detected) == len(fields) {
			valid++
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		for _, field := range fields {
			if strings.Count(line, field+":") > 0 {
				fields_detected[field] = true
			}
		}
		if line == "" {
			validate()
			fields_detected = make(map[string]bool)
		}
	}

	validate()
	fmt.Println(valid)
}
