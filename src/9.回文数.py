#
# @lc app=leetcode.cn id=9 lang=python3
#
# [9] 回文数
#

# @lc code=start
class Solution:
    def isPalindrome(self, x: int) -> bool:
        s1 = str(x)
        return s1 == s1[::-1]
# @lc code=end

if __name__ == "__main__":
    s = Solution()
    print(s.isPalindrome(123))
    print(s.isPalindrome(-121))
    print(s.isPalindrome(121))
    print(s.isPalindrome(0))

