#
# @lc app=leetcode.cn id=3 lang=python3
#
# [3] 无重复字符的最长子串
#

# @lc code=start
class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        maxn, j, mp = 0, 0, set()
        for i, c in enumerate(s):
            while c in mp:
                mp.remove(s[j])
                j += 1
            mp.add(c)
            if i - j + 1 > maxn:
                maxn = i - j + 1
        return maxn
            
# @lc code=end

if __name__ == "__main__":
    s = Solution()
    print(s.lengthOfLongestSubstring("bbbbb"))