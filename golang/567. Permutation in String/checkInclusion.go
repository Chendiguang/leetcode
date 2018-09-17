package CheckInclusion

/*
Example 1:
Input:s1 = "ab" s2 = "eidbaooo"
Output:True
Explanation: s2 contains one permutation of s1 ("ba").
Example 2:
Input:s1= "ab" s2 = "eidboaoo"
Output: False

滑动窗口：
	只需要考虑s2中包含s1同样个数的的字符，并且这些字符是连在一起的。
	使用一个滑动窗口在s2上滑动，窗口保证了字符是连在一起的，现在只需要统计比较他们的字符的数目是否一致。
*/

// O(l2 + (l2-l1)), more quickly
func checkInclusion(s1 string, s2 string) bool {
	l1, l2 := len(s1), len(s2)
	if l1 > l2 {
		return false
	}

	// Suppose all character in lower
	mp1 := make([]int, 26)
	mp2 := make([]int, 26)
	for i := 0; i < l1; i++ {
		mp1[s1[i]-'a']++
		mp2[s2[i]-'a']++
	}

	cnt := 0
	for i := 0; i < 26; i++ {
		if mp1[i] == mp2[i] {
			cnt++
		}
	}

	for j := 0; j < l2-l1; j++ {
		l, r := s2[j]-'a', s2[j+l1]-'a'
		if cnt == 26 {
			return true
		}
		mp2[r]++
		if mp1[r] == mp2[r] {
			cnt++
		} else if mp2[r] == mp1[r]+1 {
			cnt--
		}
		mp2[l]--
		if mp1[l] == mp2[l] {
			cnt++
		} else if mp2[l] == mp1[l]-1 {
			cnt--
		}
	}
	return cnt == 26
}

// O(l1 + 26*(l2-l1))
func checkInclusion2(s1 string, s2 string) bool {
	l1, l2 := len(s1), len(s2)
	if l1 > l2 {
		return false
	}
	mp1 := make([]int, 26)
	mp2 := make([]int, 26)
	for i := 0; i < l1; i++ {
		mp1[s1[i]-'a']++
		mp2[s2[i]-'a']++
	}
	if IsEqual(mp1, mp2) {
		return true
	}

	for i := 0; i < l2-l1; i++ {
		mp2[s2[i]-'a']--
		mp1[s2[l1+i]-'a']++
		if IsEqual(mp1, mp2) {
			return true
		}
	}
	return false
}

// Suppose len(a) == len(b), O(26)
func IsEqual(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
