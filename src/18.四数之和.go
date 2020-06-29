package main

/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */
// hash-table|two-pointers

import (
	"fmt"
	"sort"
)

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0)
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k, l := j+1, n-1; k < l; {
				if k > j+1 && nums[k] == nums[k-1] {
					k++
					continue
				}
				if nums[i]+nums[j]+nums[k]+nums[l] == target {
					ans = append(ans, []int{nums[i], nums[j], nums[k], nums[l]})
					k++
				} else if nums[i]+nums[j]+nums[k]+nums[l] < target {
					k++
				} else {
					l--
				}
			}
		}
	}
	return ans
}

// @lc code=end
func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}
