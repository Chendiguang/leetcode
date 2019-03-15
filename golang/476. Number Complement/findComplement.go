/*
Given a positive integer, output its complement number.
The complement strategy is to flip the bits of its binary representation.

用异或^来求解，比如101 ^ 111 = 010。那么怎么得到111呢？考虑111 + 1 = 1000，
而1000又是最小的大于101的只有一位是1的二进制数
*/
package findComplement

func findComplement(num int) int {
	mask, tmp := 1, num
	for tmp > 0 {
		mask <<= 1
		tmp >>= 1
	}
	return num ^ (mask - 1)
}

func findComplement1(num int) int {
	mask := 1
	for mask <= num {
		mask <<= 1
	}
	return num ^ (mask - 1)
}
