func maxSubArray(nums []int) int {
    currSum := nums[0]
    maxSum := nums[0]

    for i:=1; i < len(nums); i++ {
        if currSum+nums[i] > nums[i] {
            currSum = currSum + nums[i]
        } else {
            currSum = nums[i]
        }

        if maxSum < currSum {
            maxSum = currSum
        }
    }

    return maxSum
}