# two pointers
class Solution1(object):
    def longestMountain(self, A):
        n = len(A)
        ans = base = 0
        while base < n:
            end = base 
            # 在第一个符合条件的山脉起点
            if end + 1 < n and A[end] > A[end+1]:
                # 在到达山脉峰顶之前
                while end + 1 < n and A[end] > A[end+1]:
                    end += 1
                
                # if end is really a peak.
                if end + 1 < n and A[end] < A[end+1]:
                    # 在到达山脉右边脚之前
                    while end + 1 < n and A[end] < A[end+1]:
                        end += 1
                    ans = max(ans, end-base+1)

            base = max(end, base+1)
        return ans