package _053_max_sub_sum

//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
//示例:
//
//输入: [-2,1,-3,4,-1,2,1,-5,4],
//输出: 6
//解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
//定义dp[i]  表示已num[i]结尾的最大和
//
// dp[i] = max(  num[i] 自身
//                 dp[i-1] + num[i])
//dp[0] = num[0]

func maxSubArray(nums []int) int {
	lenNum := len(nums)
	if lenNum == 0 {
		return 0
	}
	dp := make([]int, lenNum)
	dp[0] = nums[0]
	for i := 1; i < lenNum; i++ {
		if dp[i-1]+nums[i] > nums[i] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
	}
	max := dp[0]
	for i := 1; i < lenNum; i++ {
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
