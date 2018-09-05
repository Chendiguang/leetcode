class Solution(object):
    def pushDominoes(self, dominoes):
        """
        :type dominoes: str
        :rtype: str
        """
        d = 'L' + dominoes + 'R'
        i, n = 0, len(d)
        ans = []
        for j in range(1, n):
            if d[j] == '.':
                continue
            if i != 0:
                ans.append(d[i])
            m = j - i - 1
            if d[i] == d[j]: 
                # R....R => RRRRRR
                # L....L => LLLLLL
                for k in range(m-1):
                    ans.append(d[i])
            elif d[i] == 'L' and d[j] == 'R': # L....R => L....R
                for k in range(m-1):
                    ans.append('.')
            else:
                # R...L => RRLL or RR.LL
                for k in range(m//2):
                    ans.append('R')
                if m % 2 == 1:
                    ans.append('.')
                for k in range(m//2):
                    ans.append('L')
            i = j
        return '.'.join(ans)

                
