package totalFruit

// 维护一个只有两个不同数字的窗口，求一个最长的区间长度
func totalFruit(tree []int) int {
	if tree == nil || len(tree) == 0 {
		return 0
	}
	l, m := 0, 0
	maps := make(map[int]int)
	// 遍历窗口的后延
	for r := 0; r < len(tree); r++ {
		maps[tree[r]]++
		// 维护窗口的种类数量为2
		for len(maps) > 2 {
			maps[tree[l]]--
			if maps[tree[l]] == 0 {
				delete(maps, tree[l])
			}
			l++
		}
		m = max(m, r-l+1)
	}
	return m
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
