package main

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */
// two-pointers

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	maxn, mp := 0, map[uint8]int{}
	for i, j := 0, 0; i < len(s); i++ {
		for _, ok := mp[s[i]]; ok; _, ok = mp[s[i]] {
			delete(mp, s[j])
			j++
		}
		mp[s[i]] = 1
		if i-j+1 > maxn {
			maxn = i - j + 1
		}
	}
	return maxn
}

// @lc code=end

// func main() {
// 	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
// }
