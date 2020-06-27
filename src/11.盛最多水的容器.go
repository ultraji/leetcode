package main

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */
// two-pointers

// @lc code=sbtart
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	res := 0
	for i, j := 0, len(height)-1; i < j; {
		vol := min(height[i], height[j]) * (j - i)
		res = max(res, vol)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return res
}

// @lc code=end
// func main() {
// 	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
// }
