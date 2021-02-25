package leetcode1225

// 滑动窗口法
func maxSatisfied(customers []int, grumpy []int, X int) int {
	sum := 0
	for i := 0; i < len(customers); i++ {
		sum += customers[i] * (grumpy[i] ^ 1)
	}
	increase := 0
	right := 0
	for ; right < X; right++ {
		increase += customers[right] * grumpy[right]
	}
	maxInc := increase
	for ; right < len(customers); right++ {
		increase = increase - customers[right-X]*grumpy[right-X] + customers[right]*grumpy[right]
		maxInc = max(maxInc, increase)
	}
	return sum + maxInc
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
