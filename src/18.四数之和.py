#
# @lc app=leetcode.cn id=18 lang=python3
#
# [18] 四数之和
# two-pointers

"""
思路和三数之和方法一致，动后两个数的指针，把时间复杂度从O(n^4)降到O(n^3)。
同时做适当的剪枝
"""

from typing import List

# @lc code=start
class Solution:
    def fourSum(self, nums: List[int], target: int) -> List[List[int]]:
        nums, n = sorted(nums), len(nums)
        res = set()
        i = 0
        while i < n-3:
            j = i+1
            while j < n-2:
                k, m = j+1, j+2
                if nums[i]+nums[j]+nums[k]+nums[m] > target:
                    break
                if nums[i]+nums[j]+nums[n-2]+nums[n-1] < target:
                    j += 1
                    continue
                while m < n-1 and nums[i]+nums[j]+nums[k]+nums[m] < target:
                    m += 1
                while k < m:
                    if nums[i]+nums[j]+nums[k]+nums[m] == target:
                        res.add((nums[i], nums[j], nums[k], nums[m]))
                        k += 1
                    elif nums[i]+nums[j]+nums[k]+nums[m] < target:
                        k += 1
                    else:
                        m -= 1
                j += 1
            i += 1
        res = [list(item) for item in res]
        return res

# @lc code=end
s = Solution()
print(s.fourSum([-1,0,1,2,-1,-4], -1))
print(s.fourSum([1,0,-1,0,-2,2], 0))
