#
# @lc app=leetcode.cn id=15 lang=python3
#
# [15] 三数之和
# two-pointers

"""
1. 首先对数据进行升序排序；
2. 第一二三的数的索引分别为i，j(i+1)，k(i+2)
3. 循环：
    a. 移动k指针：直到三数相加大于0位置或者k移动到数组末尾；
    b. 若j < k：若和小于0，j+=1；若和大于0， k-=1；若等于0，则保存数据
    c. 若j >= k：则自增i，同时保持i，j(i+1)，k(i+2)；重复abc步骤
"""

from typing import List

# @lc code=start
class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums, n = sorted(nums), len(nums)
        res = set()
        i = 0
        while i < n-2:
            j, k = i+1, i+2
            while k < n-1 and nums[i] + nums[j] + nums[k] < 0:
                k += 1
            while j < k:
                if nums[i] + nums[j] + nums[k] == 0:
                    res.add((nums[i], nums[j], nums[k]))
                    j += 1
                elif nums[i] + nums[j] + nums[k] < 0:
                    j += 1
                else:
                    k -= 1
            i += 1
        res = [list(item) for item in res]
        return res

# @lc code=end
s = Solution()
print(s.threeSum([0,-4,-1,-4,-2,-3,2]))
print(s.threeSum([-1, 0, 1, 2, -1, -4]))