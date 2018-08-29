package leetcode

import "sort"

/*
881. Boats to Save People

The i-th person has weight people[i], and each boat can carry a maximum weight of limit.
Each boat carries at most 2 people at the same time, provided the sum of the weight of those people is at most limit.
Return the minimum number of boats to carry every given person.  (It is guaranteed each person can be carried by a boat.)
*/
// two pointers
func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	i, j := 0, len(people)-1
	cnt := 0
	for i <= j {
		cnt++
		if people[i]+people[j] <= limit {
			i++
		}
		j--
	}
	return cnt
}
