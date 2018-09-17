import collections

class Solution(object):
    def uniqueLetterString(self, S):
        mp = collections.defaultdict(list)
        for i, c in enumerate(S):
            mp[c].append(i)

        res = 0
        for lst in mp.values():
            lst = [-1] + lst + [len(S)]
            for i in range(1, len(lst)-1):
                res += (lst[i+1]-lst[i]) * (lst[i]-lst[i-1])
        return res