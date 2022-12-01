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
				switch field {
				case "byr":
					byr := strings.Split(line, "byr:")[1]
					byr = strings.Split(byr, " ")[0]
					if len(byr) == 4 {
						byr_int := 0
						fmt.Sscanf(byr, "%d", &byr_int)
						if !(byr_int < 1920 || byr_int > 2002) {
							fields_detected[field] = true
						}
					}
				case "iyr":
					iyr := strings.Split(line, "iyr:")[1]
					iyr = strings.Split(iyr, " ")[0]
					if len(iyr) == 4 {
						iyr_int := 0
						fmt.Sscanf(iyr, "%d", &iyr_int)
						if !(iyr_int < 2010 || iyr_int > 2020) {
							fields_detected[field] = true
						}
					}
				case "eyr":
					eyr := strings.Split(line, "eyr:")[1]
					eyr = strings.Split(eyr, " ")[0]
					if len(eyr) == 4 {
						eyr_int := 0
						fmt.Sscanf(eyr, "%d", &eyr_int)
						if !(eyr_int < 2020 || eyr_int > 2030) {
							fields_detected[field] = true
						}
					}
				case "hgt":
					hgt := strings.Split(line, "hgt:")[1]
					hgt = strings.Split(hgt, " ")[0]
					if strings.HasSuffix(hgt, "cm") {

						hgt_int := 0
						fmt.Sscanf(hgt, "%dcm", &hgt_int)
						if !(hgt_int < 150 || hgt_int > 193) {
							fields_detected[field] = true
						}
					}
					if strings.HasSuffix(hgt, "in") {
						hgt_int := 0
						fmt.Sscanf(hgt, "%din", &hgt_int)
						if !(hgt_int < 59 || hgt_int > 76) {
							fields_detected[field] = true
						}
					}
				case "hcl":
					hcl := strings.Split(line, "hcl:")[1]
					hcl = strings.Split(hcl, " ")[0]
					if len(hcl) == 7 {
						if !(hcl[0] != '#') {
							fields_detected[field] = true
						}
						for i := 1; i < 7; i++ {
							if !(hcl[i] < '0' || hcl[i] > '9') && (hcl[i] < 'a' || hcl[i] > 'f') {
								fields_detected[field] = true
							}
						}
					}
				case "ecl":
					ecl := strings.Split(line, "ecl:")[1]
					ecl = strings.Split(ecl, " ")[0]
					if !(ecl != "amb" && ecl != "blu" && ecl != "brn" && ecl != "gry" && ecl != "grn" && ecl != "hzl" && ecl != "oth") {
						fields_detected[field] = true
					}
				case "pid":
					pid := strings.Split(line, "pid:")[1]
					pid = strings.Split(pid, " ")[0]
					if !(len(pid) != 9) {
						fields_detected[field] = true
					}
					if len(pid) == 9 {
						pid_int := 0
						fmt.Sscanf(pid, "%d", &pid_int)
						if pid_int > 0 {
							fields_detected[field] = true
						}
					}
				}
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
