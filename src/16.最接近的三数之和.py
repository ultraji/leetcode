#
# @lc app=leetcode.cn id=16 lang=python3
#
# [16] 最接近的三数之和
# two-pointers

from math import fabs
from typing import List

# @lc code=start
class Solution:
    def threeSumClosest(self, nums: List[int], target: int) -> int:
        nums, n = sorted(nums), len(nums)
        minx, res = 1000000000, 0
        i = 0
        while i < n-2:
            j, k = i+1, i+2
            while k < n-1 and nums[i] + nums[j] + nums[k] < target:
                k += 1
            while j < k:
                tmp = nums[i] + nums[j] + nums[k]
                x = fabs(tmp-target)
                if x < minx:
                    minx, res = x, tmp
                if tmp > target:
                    k -= 1
                elif tmp < target:
                    j += 1
                elif tmp == target:
                    return target
            i += 1
        return res

# @lc code=end

s = Solution()
print(s.threeSumClosest([1,2,4,8,16,32,64,128], 82))
print(s.threeSumClosest([0,0,0], 1))
print(s.threeSumClosest([-1,2,1,-4], 1))
