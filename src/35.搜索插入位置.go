package main

/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */
// binary-search

import "fmt"

// @lc code=start
func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}

// @lc code=end
func main() {
	fmt.Println(searchInsert([]int{1, 2, 3, 4, 5, 6, 7}, 6))
}
