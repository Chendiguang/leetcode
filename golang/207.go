package leetcode

// topo sort, 20ms
func canFinish(numCourses int, prerequisites [][]int) bool {
	indegrees := make([]int, numCourses) // 每个节点的入度
	visited := make([]bool, numCourses)  // 标记结点是否已被排序
	graph := make([][]int, numCourses)   // 邻接表
	pick, findOne := 0, false

	// initialization
	for i := 0; i < numCourses; i++ {
		tmp := []int{}
		graph[i] = tmp
	}

	// 用一个数组存储每个节点的入度，
	// 顶点的出边条数称为该顶点的出度，顶点的入边条数称为该项点的入度
	for _, edge := range prerequisites {
		// 对于有向边(i,j)，j是边的起点，i是边的终点。因此处理是将i加入到j的vector内，同时i的入度加1
		i, j := edge[0], edge[1]
		graph[j] = append(graph[j], i)
		indegrees[i]++
	}

	// topo sort
	// 在每一次循环中，找出一个入度为0的结点，找出后标记为已排序，
	// 并遍历该结点的邻接表，将没有被排序的结点的入度减1，继续循环；
	// 如果没有找到，直接结束循环并返回false。
	// 在所有结点都标记为已排序后返回true
	for pick < numCourses {
		findOne = false
		for i := 0; i < numCourses && !findOne; i++ {
			if indegrees[i] == 0 && !visited[i] {
				for j := 0; j < len(graph[i]); j++ {
					if !visited[graph[i][j]] {
						indegrees[graph[i][j]]--
					}
				}
				visited[i] = true
				findOne = true
				pick++
				continue
			}
		}
		if !findOne {
			return false
		}
	}
	return true
}

// BFS, 12ms
func canFinishBfs(numCourses int, prerequisites [][]int) bool {
	if numCourses <= 0 {
		return false
	}
	if prerequisites == nil || len(prerequisites) == 0 {
		return true
	}

	degrees := make([]int, numCourses, numCourses)
	graph := make([][]int, numCourses)
	queue := make([]int, 0, numCourses)
	removed := 0

	for _, edge := range prerequisites {
		start, end := edge[0], edge[1]
		if start == end {
			return false
		}
		degrees[start]++
		graph[end] = append(graph[end], start)
	}

	for i, n := range degrees {
		//store courses that have no prerequisites
		if n == 0 {
			queue = append(queue, i)
			removed++
		}
	}

	for len(queue) > 0 {
		n := queue[0]
		for _, out := range graph[n] {
			degrees[out]--
			if degrees[out] == 0 {
				queue = append(queue, out)
				removed++
			}
		}
		if len(queue) <= 1 {
			break
		}
		queue = queue[1:]
	}
	return removed == numCourses
}
