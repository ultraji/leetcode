package main

/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	mp1 := map[string]int {"I":1, "V":5, "X":10, "L":50, "C":100, "D":500, "M":1000}
	mp2 := map[string]int {"IV":4, "IX":9, "XL":40, "XC":90, "CD":400, "CM":900}
	res := 0
	for k := range mp2 {
		if strings.Index(s, k) != -1 {
			s = strings.Replace(s, k, "", 1)
			res += mp2[k]
		}
	}
	for k := range mp1 {
		for ; strings.Index(s, k) != -1; {
			s = strings.Replace(s, k, "", 1)
			res += mp1[k]
		}
	}
	return res
}
// @lc code=end

