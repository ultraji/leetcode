package main

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */
// two-pointers

import (
	"fmt"
	"sort"
)

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, n-1; j < k; {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if nums[i]+nums[j]+nums[k] == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				j++
			} else if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return ans
}

// @lc code=end
func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
