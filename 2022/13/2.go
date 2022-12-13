package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

const KEY_1 = "[[2]]"
const KEY_2 = "[[6]]"

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

func main() {
	packets := make([][]any, 0)
	key_id1, key_id2 := -1, -1
	key_p1, key_p2 := parseArray(KEY_1), parseArray(KEY_2)
	packets = append(packets, key_p1)
	packets = append(packets, key_p2)
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			packets = append(packets, parseArray(line))
		}
	}
	for i := 0; i < len(packets); i++ {
		for j := 0; j < len(packets); j++ {
			if compare(packets[i], packets[j]) == 1 {
				packets[i], packets[j] = packets[j], packets[i]
			}
		}
	}
	for i := 0; i < len(packets); i++ {
		if compare(packets[i], key_p1) == 0 {
			key_id1 = i + 1
		}
		if compare(packets[i], key_p2) == 0 {
			key_id2 = i + 1
		}
		if key_id1 != -1 && key_id2 != -1 {
			break
		}
	}
	fmt.Println(key_id1 * key_id2)
}
