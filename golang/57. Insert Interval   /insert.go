package insert

/**
 * Definition for an interval.
Example 1:

Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]
*/
type Interval struct {
	Start int
	End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	l := len(intervals)
	res := make([]Interval, 0, l+1)
	if l == 0 {
		res[0] = newInterval
		return res
	}
	start, end := newInterval.Start, newInterval.End
	i := 0
	for i < l {
		if start <= intervals[i].End {
			if end < intervals[i].Start {
				break
			}
			if intervals[i].Start < start {
				start = intervals[i].Start
			}
			if intervals[i].End > end {
				end = intervals[i].End
			}
		} else {
			res[i] = intervals[i]
		}
		i++
	}
	res[i] = Interval{Start: start, End: end}
	for j := i; j < l+1; j++ {
		res[j+1] = intervals[j]
	}
	return res
}
