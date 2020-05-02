#
# @lc app=leetcode.cn id=11 lang=python3
#
# [11] 盛最多水的容器
#

"""
O(n)
"""

from typing import List

# @lc code=start
class Solution:
    def maxArea(self, height: List[int]) -> int:
        i, j, res = 0, len(height)-1, 0
        while i < j:
            vol = min(height[i], height[j]) * (j-i)
            res = max(res, vol)
            if height[i] < height[j]:
                i += 1
            else:
                j -= 1
        return res
# @lc code=end

s = Solution()
print(s.maxArea([1,8,6,2,5,4,8,3,7]))