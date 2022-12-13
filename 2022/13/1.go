package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func parseArray(input string) []any {
	var list []any
	json.Unmarshal([]byte(input), &list)
	return list
}

func compare(a1 []any, a2 []any) int {
	for i := 0; i < len(a1); i++ {
		inside_res := 0
		if i == len(a2) {
			return -1
		}
		t1 := reflect.TypeOf(a1[i])
		t2 := reflect.TypeOf(a2[i])
		if t1 == t2 && t1.Name() == "float64" {
			if a1[i].(float64) < a2[i].(float64) {
				return 1
			} else if a1[i].(float64) > a2[i].(float64) {
				return -1
			}
		} else if t1 == t2 {
			inside_res = compare(a1[i].([]any), a2[i].([]any))
			if inside_res != 0 {
				return inside_res
			}
		} else {
			if t1.Name() == "float64" {
				inside_res = compare([]any{a1[i]}, a2[i].([]any))
			} else {
				inside_res = compare(a1[i].([]any), []any{a2[i]})
			}
			if inside_res != 0 {
				return inside_res
			}
		}
	}
	if len(a1) < len(a2) {
		return 1
	}
	return 0
}

func isOrdered(l1, l2 string) int {
	a1 := parseArray(l1)
	a2 := parseArray(l2)
	return compare(a1, a2)
}

func main() {
	l1, l2 := "", ""
	pair_id := 1
	pairs_sum := 0
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if l1 == "" {
			l1 = scanner.Text()
		} else if l2 == "" {
			l2 = scanner.Text()
		}
		if l1 != "" && l2 != "" {
			if isOrdered(l1, l2) == 1 {
				pairs_sum += pair_id
			}
			l1, l2 = "", ""
			pair_id++
		}
	}
	fmt.Println(pairs_sum)
}
