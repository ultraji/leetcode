package main

/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */
// binary-search

import (
	"fmt"
	"math"
)

// @lc code=start
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	positive := true
	if dividend < 0 {
		positive, dividend = !positive, -dividend
	}
	if divisor < 0 {
		positive, divisor = !positive, -divisor
	}
	res := 0
	for dividend >= divisor {
		p := divisor
		ans := 1
		for ; p <= dividend; p <<= 1 {
			ans += ans
		}
		dividend -= p >> 1
		res += ans >> 1
	}
	if !positive {
		res = -res
	}
	return res
}

// @lc code=end
func main() {
	fmt.Println(divide(-2147483648, -1))
}
