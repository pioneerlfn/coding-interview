/*
@Time : 2019-11-15 17:53
@Author : lfn
@File : cut_rope
*/

package _14_cut_rope

func cutRope(length int) int {
	if length < 2 {
		return 0
	}
	if length == 2 {
		return 1
	}
	if length == 3 {
		return 2
	}
	dp := make([]int, length+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	for i := 4; i <= length; i++ {
		for j := 1; j <= i/2; j++ {
			dp[i] = max(dp[i], dp[j]*dp[i-j])
		}
	}
	return dp[length]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}