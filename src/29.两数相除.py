#
# @lc app=leetcode.cn id=29 lang=python3
#
# [29] 两数相除
#

# @lc code=start
class Solution:
    def divide(self, dividend: int, divisor: int) -> int:
        if dividend == 0:
            return 0
        if divisor == 0:
            return 2**31-1
        sg = (dividend > 0) ^ (divisor > 0)
        dd, dr = abs(dividend), abs(divisor)
        res = 0
        t, dr = 1<<31, dr<<31
        while dd > 0:
            while dr > dd:
                t, dr = t>>1, dr>>1
            res, dd = res+t, dd-dr
        if dd != 0:
            res -= 1
        if sg:
            res = -res
        if res > 2**31-1 or res < -2**31:
            return 2**31-1
        return res
# @lc code=end

