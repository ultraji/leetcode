package main

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	f := func(x int) int {
		if x == 0 {
			return 0
		}
		flag, r := 1, 0
		if x < 0 {
			x, flag = -x, -1
		}
		for ; x%10 == 0; x /= 10 {
		}
		for ; x != 0; x /= 10 {
			r = r*10 + x%10
		}
		r *= flag
		return r
	}
	return x == f(x)
}

// @lc code=end
// func main() {
// 	fmt.Print(isPalindrome(-123))
// }
