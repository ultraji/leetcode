package main

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */
// hash-table

import "fmt"

// @lc code=start
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, num := range nums {
		if v, ok := mp[num]; ok {
			return []int{v, i}
		}
		mp[target-num] = i
	}
	return nil
}

// @lc code=end

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
