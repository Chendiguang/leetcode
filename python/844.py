# 利用栈的用法, O(M+N)
class Solution:
    def backspaceCompare(self, S, T):
        """
        :type S: str
        :type T: str
        :rtype: bool
        """
        def build(s):
            res = []
            for c in S:
                if c != '#':
                    res.append(c)
                elif res:
                    res.pop()
            return ''.join(res)
        return build(S) == build(T)