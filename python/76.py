class Solution(object):
    def minWindow(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: str
        """
        lt = len(t)
        if len(t) > len(s):
            return ""
        lst = [0] * 256
        for i in t:
            lst[ord(i)] += 1

        l, r , head, gap = 0, 0, -1, len(s)+1
        while r < len(s):
            lst[ord(s[r])] -= 1
            if lst[ord(s[r])] >= 0: # 当前字符在t里，并且不超出它的次数
                lt -= 1
            r += 1

            # lt=0, 证明此时s[l:r]包含了t的所有字符
            while lt == 0:
                if r - l < gap: # 更新结果的左边界head和区间大小gap
                    head = l
                    gap = r - l
                # 更新左边界
                lst[ord(s[l])] += 1
                if lst[ord(s[l])] > 0:
                    lt += 1
                l += 1
        if head == -1:
            return ""
        return s[head:head+gap]
  