#
# @lc app=leetcode.cn id=8 lang=python3
#
# [8] 字符串转换整数 (atoi)
#

# @lc code=start
class Solution:
    def myAtoi(self, str: str) -> int:
        if not str:
            return 0
        INT_MIN, INT_MAX = -2**31, 2**31-1
        i, syb, num = 0, '', ''
        while i < len(str) and str[i] == ' ': i = i+1

        bg = i
        while i < len(str) and str[i] in ['-','+']: i = i+1
        syb = str[bg:i]

        while i < len(str)and str[i] == '0': i = i+1
        
        bg = i
        while i < len(str) and str[i] >= '0' and str[i] <= '9': i = i+1
        num = str[bg:i]
        
        if not num or len(syb) > 1:
            return 0
        num = eval(syb + num)
        if num > INT_MAX:
            return INT_MAX
        elif num < INT_MIN:
            return INT_MIN
        else:
            return num
        

# @lc code=end
s = Solution()
print(s.myAtoi(""))
print(s.myAtoi(" "))
print(s.myAtoi("   -00000042.3453"))
print(s.myAtoi("   -42.3453"))
print(s.myAtoi("words and 987"))
print(s.myAtoi("4193 with words"))


