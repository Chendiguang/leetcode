package merge

import "sort"

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	l := len(intervals)
	if intervals == nil || l == 0 {
		return intervals
	}

	// 排序好
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	var res []Interval
	start, end := intervals[0].Start, intervals[0].End
	for i := 1; i < l; i++ {
		if intervals[i].Start > end {
			res = append(res, Interval{start, end})
			start, end = intervals[i].Start, intervals[i].End
		} else {
			if end < intervals[i].Start {
				break
			}
			if end < intervals[i].End {
				end = intervals[i].End
			}
			if intervals[i].Start < start {
				start = intervals[i].Start
			}
		}
	}
	res = append(res, Interval{start, end})
	return res
}
