package main

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */
// stack

import "fmt"

// @lc code=start
func isValid(s string) bool {
	left := map[string]int{"(": 1, "[": 2, "{": 3}
	right := map[string]int{")": 1, "]": 2, "}": 3}
	all := map[string]int{"()": 1, "[]": 2, "{}": 3}

	st := make([]string, 0)
	for i := 0; i < len(s); i++ {
		if _, ok := left[string(s[i])]; ok {
			st = append(st, string(s[i]))
		} else if _, ok := right[string(s[i])]; ok {
			if len(st) == 0 {
				return false
			}
			t := st[len(st)-1]
			st = st[:len(st)-1]
			if _, ok := all[t+string(s[i])]; !ok {
				return false
			}
		}
	}
	if len(st) != 0 {
		return false
	}
	return true
}

// @lc code=end
func main() {
	fmt.Println(isValid("[](){}"))
}
