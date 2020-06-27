package main

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */
// string|dynamic-programming|manacher

import (
	"fmt"
)

// @lc code=start
func longestPalindrome(s string) string {
	if s == "" {
		return s
	}

	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	ans := ""
	for j := 0; j < n; j++ {
		for i := j; i >= 0; i-- {
			if s[j] == s[i] && (j-i <= 1 || dp[i+1][j-1]) {
				dp[i][j] = true
				if j-i+1 > len(ans) {
					ans = s[i : j+1]
				}
			}
		}
	}
	return ans
}

// @lc code=end
func main() {
	fmt.Println(longestPalindrome("aaaa"))
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
}
