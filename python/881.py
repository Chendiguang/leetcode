'''
881. Boats to Save People

The i-th person has weight people[i], and each boat can carry a maximum weight of limit.
Each boat carries at most 2 people at the same time, provided the sum of the weight of those people is at most limit.
Return the minimum number of boats to carry every given person.  (It is guaranteed each person can be carried by a boat.)
'''

class Solution:
    def numRescueBoats(self, people, limit):
        """
        :type people: List[int]
        :type limit: int
        :rtype: int
        """
        people.sort()
        i, j = 0, len(people)-1
        cnt = 0
        while i <= j:
            cnt += 1
            if people[i] + people[j] <= limit:
                i += 1
            j -= 1
        return cnt


def numRescueBoats(people, limit):
    people.sort()
    i, j = 0, len(people) - 1
    cnt = 0
    while i <= j:
        cnt += 1
        if people[i] + people[j] <= limit:
            i += 1
        j -= 1
    return cnt