package leetcode

/*
210. Course Schedule II

There are a total of n courses you have to take, labeled from 0 to n-1.
Some courses may have prerequisites, for example to take course 0 you have to first take course 1,
which is expressed as a pair: [0,1]
Given the total number of courses and a list of prerequisite pairs,
return the ordering of courses you should take to finish all courses.
There may be multiple correct orders, you just need to return one of them.
If it is impossible to finish all courses, return an empty array.

Input: 2, [[1,0]]
Output: [0,1]
Explanation: There are a total of 2 courses to take. To take course 1 you should have finished
             course 0. So the correct course order is [0,1] .
*/

// BFS topological sorting.
func findOrder(numCourses int, prerequisites [][]int) []int {
	mps := make([]int, numCourses)
	res := make([]int, numCourses)
	idx := 0
	// 用一个数组存储每个节点的入度
	for i := 0; i < len(prerequisites); i++ {
		mps[prerequisites[i][0]]++
	}

	queue := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		if mps[i] == 0 {
			queue = append(queue, i)
			res[idx] = i
			idx++
		}
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for i := 0; i < len(prerequisites); i++ {
			from, to := prerequisites[i][1], prerequisites[i][0]
			if cur == from {
				mps[to]--
				if mps[to] == 0 {
					queue = append(queue, to)
					res[idx] = to
					idx++
				}
			}
		}
	}
	if idx == numCourses {
		return res
	}
	return []int{0}
}

func findOrder1(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	indegree := make([]int, numCourses)
	queue, stack := []int{}, []int{}

	// 用一个数组存储每个节点的入度
	for _, pair := range prerequisites {
		from, to := pair[0], pair[1]
		graph[from] = append(graph[from], to)
		indegree[to]++
	}
	for i, v := range indegree {
		//store courses that have no prerequisites
		if v == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		stack = append(stack, node)

		for _, neighbor := range graph[node] {
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	// reverse
	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}

	if len(stack) < numCourses {
		return []int{}
	}
	return stack
}
