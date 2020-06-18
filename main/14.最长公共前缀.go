package main

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func lcp(s1, s2 string) string {
	index := 0
	slen := min(len(s1), len(s2))
	for i := 0; i < slen; i++ {
		if s1[i] != s2[i] {
			break
		}
		index++
	}
	return s1[:index]
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}
	p := strs[0]
	for i := 1; i < len(strs); i++ {
		p = lcp(p, strs[i])
	}
	return p
}
// @lc code=end

