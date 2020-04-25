#
# @lc app=leetcode.cn id=7 lang=python3
#
# [7] 整数反转
#

# @lc code=start
class Solution:
    def reverse(self, x: int) -> int:
        if x == 0:
            return 0
        s = ""
        if x < 0:
            x = -x
            s += "-"
        while x%10 == 0:
            x //= 10
        while x != 0:
            s, x = s+str(x%10), x//10
        x = eval(s)
        if x > 2**31-1 or x < -2**31:
            x = 0
        return x
# @lc code=end

if __name__ == "__main__":
    s = Solution()
    print(s.reverse(123))
    print(s.reverse(-123))
    print(s.reverse(-120))
    print(s.reverse(-0))
