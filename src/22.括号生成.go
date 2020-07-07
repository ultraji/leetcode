package main

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */
// backtracking

import "fmt"

// @lc code=start
func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	ans := make([]string, 0)
	for _, item := range generateParenthesis(n - 1) {
		ans = append(ans, "("+item+")")
	}
	for i := 1; i < n; i++ {
		for _, x := range generateParenthesis(i) {
			for _, y := range generateParenthesis(n - i) {
				ans = append(ans, x+y)
			}
		}
	}
	mp := make(map[string]int, 0)
	res := make([]string, 0)
	for _, item := range ans {
		l := len(mp)
		mp[item] = 0
		if len(mp) != l {
			res = append(res, item)
		}
	}
	return res
}

// @lc code=end

func main() {
	fmt.Println(generateParenthesis(3))
}
