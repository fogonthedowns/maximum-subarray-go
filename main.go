package main

func maxSubArray(nums []int) int {
	currSum := nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		currSum = max(nums[i], currSum+nums[i])
		maxSum = max(maxSum, currSum)
	}
	return maxSum
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
