package leetcode1225

// 一. 暴力法
func minKBitFlips(A []int, K int) int {
	n := len(A)
	res := 0
	for i := 0; i < n; i++ {
		if A[i] == 0 {
			if i+K > n {
				return -1
			}
			for j := i; j < i+K; j++ {
				A[j] ^= 1
			}
			res++
		}
	}
	return res
}

// 二. 差分数组法（常规解法）
func minKBitFlips2(A []int, K int) int {
	n := len(A)
	R := make([]int, n+1)
	count := 0
	res := 0
	for i := 0; i < n; i++ {
		count += R[i]
		if (count+A[i])%2 == 0 {
			if i+K > n {
				return -1
			}
			count++
			R[i+K]--
			res++
		}
	}
	return res
}

// 三. 差分数组（位运算解法）
func minKBitFlips3(A []int, K int) int {
	n := len(A)
	R := make([]int, n+1)
	count := 0
	res := 0
	for i := 0; i < n; i++ {
		count ^= R[i]
		if count == A[i] {
			if i+K > n {
				return -1
			}
			count ^= 1
			R[i+K] ^= 1
			res++
		}
	}
	return res
}

// 四. 差分数组（原数组修改解法）
func minKBitFlips(A []int, K int) int {
	n := len(A)
	count := 0
	res := 0
	for i := 0; i < n; i++ {
		if i >= K && A[i-K] > 1 {
			count ^= 1
			A[i-K] -= 2
		}
		if count == A[i] {
			if i+K > n {
				return -1
			}
			count ^= 1
			A[i] += 2
			res++
		}
	}
	return res
}

// 五. 滑动窗口法

func minKBitFlips5(A []int, K int) int {
	n := len(A)
	q := make([]int, 0, n)
	res := 0
	for i := 0; i < n; i++ {
		if len(q) > 0 && i-q[0] >= K {
			q = q[1:]
		}
		if (len(q)+A[i])%2 == 0 {
			if i+K > n {
				return -1
			}
			q = append(q, i)
			res++
		}
	}
	return res
}
