import heapq
class ExamRoom:

    def __init__(self, N):
        """
        :type N: int
        """
        self.pq=[(-((N-1)//2),0,N-1)]
        self.N=N
        # self.pq = [(-((N-1)//2), 0, N-1)]
        # self.N = N

    def seat(self):
        """
        :rtype: int
        """
        _, start, end = heapq.heappop(self.pq)
        if start == 0:
            heapq.heappush(self.pq, (-((end-start-1)//2),start+1, end))
            return start
        if end == self.N-1:
            heapq.heappush(self.pq, (-((end-start-1)//2),start, end-1))
            return end
        mid = (start+end)//2
        if mid-1 >= start:
            heapq.heappush(self.pq, (-((mid-start-1)//2),start, mid-1))
        if mid+1 <= end:
            heapq.heappush(self.pq, (-((end-mid-1)//2),mid+1, end))
        return mid
        
    def leave(self, p):
        """
        :type p: int
        :rtype: void
        """
        newStart, newEnd=p,p
        left,right=None,None
        for seg in self.pq:
            if seg[1]==p+1:
                right=seg
                newEnd=seg[2]
            if seg[2]==p-1:
                left=seg
                newStart=seg[1]
        if left:
            self.pq.remove(left)
        if right:
            self.pq.remove(right)
        if newStart==0:
            heapq.heappush(self.pq, (-(newEnd),newStart,newEnd))
        elif newEnd==self.N-1:
            heapq.heappush(self.pq, (-(newEnd-newStart),newStart,newEnd))
        else:
            heapq.heappush(self.pq, (-((newEnd-newStart)//2),newStart,newEnd))


# Your ExamRoom object will be instantiated and called as such:
# obj = ExamRoom(N)
# param_1 = obj.seat()
# obj.leave(p)