package main

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */
// two-pointers

import (
	"fmt"
	"math"
	"sort"
)

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	ans, minDiff := 0, math.MaxInt32
	for i := 0; i < n-2; i++ {
		for j, k := i+1, n-1; j < k; {
			if diff := target - (nums[i] + nums[j] + nums[k]); diff >= 0 {
				if diff < minDiff {
					ans, minDiff = nums[i]+nums[j]+nums[k], diff
				}
				j++
			} else {
				if -diff < minDiff {
					ans, minDiff = nums[i]+nums[j]+nums[k], -diff
				}
				k--
			}
		}
	}
	return ans
}

// @lc code=end
func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}
