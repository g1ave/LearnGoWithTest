package platesbetweencandles

func PlatesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	left, right := make([]int, n), make([]int, n)
	prefixSum := make([]int, n)
	if s[0] == '|' {
		left[0] = 0
		prefixSum[0] = 1
	} else {
		left[0] = -1
		prefixSum[0] = 0
	}
	if s[n-1] == '|' {
		right[n-1] = n - 1
	} else {
		right[n-1] = -1
	}
	for i := 1; i < n; i++ {
		if s[i] == '*' {
			left[i] = left[i-1]
			prefixSum[i] = prefixSum[i-1]
		} else {
			left[i] = i
			prefixSum[i] = prefixSum[i-1] + 1
		}
		if s[n-1-i] == '*' {
			right[n-1-i] = right[n-i]
		} else {
			right[n-1-i] = n - 1 - i
		}
	}
	ans := make([]int, len(queries))
	for i, query := range queries {
		l, r := right[query[0]], left[query[1]]
		if l >= query[1] || l == -1 || r == -1 {
			ans[i] = 0
		} else {
			ans[i] = r - l - (prefixSum[r] - prefixSum[l])
		}
	}
	return ans
}

// 思路：对于每一个查询 如果能确定距离左右边界最近的蜡烛[left, right]以及整个区间的蜡烛数量candles，
// 即可确定盘子数量 = right - left - candles
// 左右边界最近的蜡烛可以用两次遍历确定
// 整个区间的蜡烛数量可以用前缀和确定
