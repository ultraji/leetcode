package main

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */
// string|backtracking

import "fmt"

// @lc code=start
func letterCombinations(digits string) []string {
	ans := make([]string, 0)
	if digits == "" {
		return ans
	}
	mp := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	tmp := mp[int(digits[0]-'0')]
	for i := 0; i < len(tmp); i++ {
		ans = append(ans, string(tmp[i]))
	}
	for i := 1; i < len(digits); i++ {
		t := ans
		ans = make([]string, 0)
		for j := 0; j < len(t); j++ {
			tmp = mp[int(digits[i]-'0')]
			for k := 0; k < len(tmp); k++ {
				ans = append(ans, t[j]+string(tmp[k]))
			}
		}
	}
	return ans
}

// @lc code=end
func main() {
	fmt.Println(letterCombinations("23"))
}
