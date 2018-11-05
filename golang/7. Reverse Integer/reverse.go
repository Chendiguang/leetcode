package reverse

func reverse(x int) int {
	var pre = -1
	if x >= 10 {
		pre = 1
	} else if x <= -10 {
		x = -x
	} else {
		return x
	}
	res := 0
	for x > 0 {
		res, x = res*10+x%10, x/10
	}
	if res > 1<<31-1 {
		return 0
	}
	return res * pre
}
