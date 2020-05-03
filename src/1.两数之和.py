#
# @lc app=leetcode.cn id=1 lang=python3
#
# [1] 两数之和
# hash-table

from typing import List

# @lc code=start
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        tmp, mp, res = 0, set(), [0, 0]
        for i, num in enumerate(nums):
            if num in mp:
                res[1], tmp = i, num
                break
            t = target - num
            mp.add(t)
        res[0] = nums.index(target-tmp)

        return res

# @lc code=end

if __name__ == "__main__":
    s = Solution()
    print(s.twoSum([2,11,7,15], 9))