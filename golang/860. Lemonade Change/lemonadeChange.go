package main

import (
	"fmt"
)

func lemonadeChange(bills []int) bool {
	mp := make(map[int]int)
	for i, bill := range bills {
		mp[bill]++
		switch bill {
		case 10:
			if mp[5] > 0 {
				mp[5]--
			} else {
				return false
			}
		case 20:
			if mp[10] > 0 && mp[5] > 0 {
				mp[5]--
				mp[10]--
			} else if mp[5] > 3 {
				mp[5] = mp[5] - 3
			} else {
				return false
			}
		}
	}
	fmt.Println(mp)
	return true
}

func main() {
	bills := []int{5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 20, 5}
	res := lemonadeChange(bills)
	fmt.Println(res)
}
