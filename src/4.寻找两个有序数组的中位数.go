package main

/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个有序数组的中位数
 */
// binary-search|divide-and-conquer

import (
	"fmt"
	"math"
)

// @lc code=start
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

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append([]int{math.MinInt32}, nums1...)
	nums2 = append([]int{math.MinInt32}, nums2...)
	nums1 = append(nums1, math.MaxInt32)
	nums2 = append(nums2, math.MaxInt32)

	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m, n := len(nums1), len(nums2)

	left, right := 0, m
	totalLeft := (m + n + 1) / 2
	for left < right {
		i := left + (right-left+1)/2
		j := totalLeft - i
		if nums2[j] < nums1[i-1] {
			right = i - 1
		} else {
			left = i
		}
	}
	i, j := left, totalLeft-left
	ml, mr, nl, nr := nums1[i-1], nums1[i], nums2[j-1], nums2[j]

	if (m+n)%2 == 0 {
		return float64(max(ml, nl)+min(mr, nr)) / 2
	}
	return float64(max(ml, nl))
}

// @lc code=end
func main() {
	fmt.Println(findMedianSortedArrays([]int{3}, []int{-2, -1}))
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
