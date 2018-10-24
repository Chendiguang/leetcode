package main

import (
	"container/heap"
	"log"
	"math"
)

/*
You have k lists of sorted integers in ascending order.
Find the smallest range that includes at least one number from each of the k lists.
We define the range [a,b] is smaller than range [c,d]
if b-a < d-c or a < c if b-a == d-c.
*/

type PQ [][]int

func (p *PQ) Len() int {
	return len(*p)
}

func (p *PQ) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *PQ) Less(i, j int) bool {
	return (*p)[i][0] < (*p)[j][0]
}

func (p *PQ) Push(x interface{}) {
	*p = append(*p, x.([]int))
}

func (p *PQ) Pop() interface{} {
	v := (*p)[p.Len()-1]
	*p = (*p)[:p.Len()-1]
	return v
}

func smallestRange(nums [][]int) []int {
	res := []int{math.MinInt32, math.MaxInt32}
	max := math.MinInt32

	p := make(PQ, len(nums))
	for i := range nums {
		if len(nums) == 0 {
			return res
		}
		p[i] = nums[i]
		if nums[i][0] > max {
			max = nums[i][0]
		}
	}
	heap.Init(&p)
	for {
		curr := heap.Pop(&p).([]int)
		if max-curr[0] < res[1]-res[0] {
			res[0] = curr[0]
			res[1] = max
		}
		if len(curr) == 1 {
			return res
		}

		heap.Push(&p, curr[1:])
		if curr[1] > max {
			max = curr[1]
		}
	}
}

func main() {
	d := make([][]int, 3)
	d[0] = []int{4, 10, 15, 24, 26}
	d[1] = []int{0, 9, 12, 20}
	d[2] = []int{5, 18, 22, 30}
	res := smallestRange(d)
	log.Println(res)
}
