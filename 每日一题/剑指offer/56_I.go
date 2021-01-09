package main

func singleNumbers(nums []int) []int {
	// xor用来计算nums的异或和
	xor := 0

	// 计算异或和 并存到xor
	// e.g. [2,4,2,3,3,6] 异或和：(2^2)^(3^3)^(4^6)=2=010
	for i := 0; i < len(nums); i++ {
		xor ^= nums[i]
	}

	//设置mask为1，则二进制为0001
	// mask是一个二进制数，且其中只有一位是1，其他位全是0，
	// 比如000010，表示我们用倒数第二位作为分组标准，倒数第二位是0的数字分到一组，倒数第二位是1的分到另一组
	mask := 1

	// 找到第一位不是0的
	//& operator只有1&1时等于1 其余等于0
	// 用上面的e.g. 4和6的二进制是不同的 我们从右到左找到第一个不同的位就可以分组 4=0100 6=0110
	// 根据e.g. 010 & 001 = 000 = 0则 mask=010
	// 010 & 010 != 0 所以mask=010
	// 之后就可以用mask来将数组里的两个数分区分开
	for (xor & mask) == 0 {
		mask <<= 1
	}

	//两个只出现一次的数字
	a := 0
	b := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]&mask == 0 {
			a ^= nums[i]
		} else {
			b ^= nums[i]
		}
	}
	return []int{a, b}
}

func main() {

}
