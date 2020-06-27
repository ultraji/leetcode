package main

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */
// math

import (
	"math"
)

// @lc code=start
func reverse(x int) int {
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
	if r > math.MaxInt32 || r < math.MinInt32 {
		return 0
	}
	return r
}

// @lc code=end
// func main() {
// 	fmt.Println(reverse(-123))
// }
