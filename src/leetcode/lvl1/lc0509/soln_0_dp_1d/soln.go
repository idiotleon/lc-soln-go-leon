// @author: Leon
// https://leetcode.com/problems/fibonacci-number/
//
// Time Complexity:		O(`n`)
// Space Complexity:	O(`n`)
package lc0509

func fib(n int) int {
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
