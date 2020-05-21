package dp

import "fmt"

//N = 3, W = 4
//wt = [2, 1, 3]
//val = [4, 2, 3]
//背包问题
func MaxValue(captity int, cap []int, values []int, num int) int {
	dp := make([][]int, num+1)
	for i := 0; i <= num; i++ {
		dp[i] = make([]int, captity+1)
	}
	for i := 0; i <= captity; i++ {
		dp[0][i] = 0
	}
	for i := 0; i <= num; i++ {
		dp[i][0] = 0
	}

	for i := 1; i <= num; i++ {
		for w := 1; w <= captity; w++ {
			if w-cap[i-1] < 0 {
				dp[i][w] = dp[i-1][w]
			} else {
				dp[i][w] = max(dp[i-1][w-cap[i-1]]+values[i-1], dp[i-1][w])
			}
		}
	}
	fmt.Println(dp)
	return dp[num][captity]
}

func max(i, j int) int {
	if i >= j {
		return i
	} else {
		return j
	}
}

//子集背包问题  https://labuladong.gitbook.io/algo/dong-tai-gui-hua-xi-lie/bei-bao-zi-ji
//value=[1,4,5,10]
func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	if sum%2 != 0 {
		return false
	}

	halfSum := sum / 2
	num := len(nums)
	dp := make([][]bool, num+1)
	for i := 0; i <= num; i++ {
		dp[i] = make([]bool, halfSum+1)
	}
	for i := 0; i <= halfSum; i++ {
		dp[0][i] = true
	}

	for i := 1; i <= num; i++ {
		for w := 1; w <= halfSum; w++ {
			if w-nums[i-1] < 0 {
				dp[i][w] = dp[i-1][w]
			} else {
				dp[i][w] = dp[i-1][w-nums[i-1]] || dp[i-1][w]
			}
		}
	}
	return dp[num][halfSum]
}

//完全背包问题
func change(amount int, coins []int) int {
	// dp[i]表示总额为i时的方案数.
	// 转移方程: dp[i] = Σdp[i - coins[j]]; 表示 总额为i时的方案数 = 总额为i-coins[j]的方案数的加和.
	dp := make([]int, amount+1)
	// 记得初始化dp[0] = 1; 表示总额为0时方案数为1.
	dp[0] = 1

	for _, coin := range coins {
		for i := coin; i <= amount; i++ { // 从coin开始遍历，小于coin的值没有意义
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}
