package main

import (
	"math"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(str string) int {
	s := strings.TrimLeft(str, " ")
	if len(s) <= 0 {
		return 0
	}
	i, bg, syb, num := 0, 0, "", ""
	for ; i < len(s) && (s[i] == '-' || s[i] == '+'); i++ {
	}
	syb += s[bg:i]
	for ; i < len(s) && s[i] == '0'; i++ {
	}
	for bg = i; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
	}
	num += s[bg:i]

	if len(num) <= 0 || len(syb) > 1 {
		return 0
	}
	t, _ := strconv.Atoi(syb + num)
	if t > math.MaxInt32 {
		return math.MaxInt32
	} else if t < math.MinInt32 {
		return math.MinInt32
	}
	return t
}

// @lc code=end
// func main() {
// 	fmt.Println(myAtoi("sadfwe 4342"))
// }
