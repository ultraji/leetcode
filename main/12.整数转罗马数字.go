package main

/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */
// hash-table

// @lc code=start
func intToRoman(num int) string {
	mp := map[int]string{1: "I", 5: "V", 10: "X", 50: "L", 100: "C", 500: "D", 1000: "M",
		4: "IV", 9: "IX", 40: "XL", 90: "XC", 400: "CD", 900: "CM"}
	row, roman := []int{1000, 100, 10, 1}, ""
	for _, r := range row {
		if num >= r {
			t := num / r
			num %= r
			if t == 4 || t == 5 || t == 9 {
				roman += mp[t*r]
			} else {
				if t > 5 {
					roman += mp[5*r]
					t -= 5
				}
				for ; t > 0; t-- {
					roman += mp[r]
				}
			}
		}
	}
	return roman
}

// @lc code=end
