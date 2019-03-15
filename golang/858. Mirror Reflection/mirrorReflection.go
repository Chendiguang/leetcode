/*
 * 有个正方形的玻璃房子，从左下角射出一个光线，
 * 光线第一次打在右边墙上的高度是q。正方形边长是p，
 * 求最终这个光线射在了0,1,2的哪个位置。
 *
 * 利用镜子的反射原理, 把正方形看成一对无限平行的无限长的镜子
 * 求第一次满足m * p = n * q的情况， 其中m为房子的拓展次数，n为反射的拓展次数。
 * 不能同时为偶数，考虑: 2xp = 2yq ==> xp=yq，会提前相遇
 *（1）m为奇数，n为偶数
 *（2）m为偶数，n为奇数
 *（3）m、n均为奇数

	m is even & n is odd => return 0.
	m is odd & n is odd => return 1.      ===> 1 - p%2 + q%2
	m is odd & n is even => return 2.
*/

package mirrorReflection

func mirrorReflection(p int, q int) int {
	m, n := q, p
	for m%2 == 0 && n%2 == 0 {
		m /= 2
		n /= 2
	}
	if m%2 == 0 && n%2 == 1 {
		return 0
	} else if m%2 == 1 && n%2 == 1 {
		return 1
	} else if m%2 == 1 && n%2 == 0 {
		return 2
	}
	return 0
}

func mirrorReflection2(p int, q int) int {
	for p%2 == 0 && q%2 == 0 {
		p >>= 1
		q >>= 1
	}
	return 1 - p%2 + q%2
}
