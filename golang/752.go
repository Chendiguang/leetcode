package leetcode

/*
752. Open the Lock
Example 1:
Input: deadends = ["0201","0101","0102","1212","2002"], target = "0202"
Output: 6
Explanation:
A sequence of valid moves would be "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202".
Note that a sequence like "0000" -> "0001" -> "0002" -> "0102" -> "0202" would be invalid,
because the wheels of the lock become stuck after the display becomes the dead end "0102".

Note:
The length of deadends will be in the range [1, 500].
target will not be in the list deadends.
Every string in deadends and the string target will be a string of 4 digits from the 10,000 possibilities '0000' to '9999'.
要求是：有一个四位的密码锁，每次只能旋转一位，每位的数字的在0-9循环变化，避开给出的
死锁，求出开锁的最小步数。
*/

func openLock(deadends []string, target string) int {
	tmp, tar := [4]rune{}, [4]rune{}
	for i, item := range target {
		tar[i] = item
	}

}
