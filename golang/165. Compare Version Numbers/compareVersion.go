package compareVersion

import (
	"strconv"
	"strings"
)

/*
Compare two version numbers version1 and version2.
If version1 > version2 return 1;
if version1 < version2 return -1;
otherwise return 0.
*/

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	for i := 0; i < len(v1) || i < len(v2); i++ {
		// 不足的比较位对应的数值默认为0
		vi1, vi2 := 0, 0
		if i < len(v1) {
			vi1, _ = strconv.Atoi(v1[i])
		}
		if i < len(v2) {
			vi2, _ = strconv.Atoi(v2[i])
		}
		if vi1 > vi2 {
			return 1
		} else if vi1 < vi2 {
			return -1
		}
	}
	return 0
}
