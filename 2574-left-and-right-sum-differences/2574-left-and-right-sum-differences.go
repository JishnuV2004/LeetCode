func leftRightDifference(nums []int) []int {
    leftSum := make([]int, len(nums))
    rightSum := make([]int, len(nums))

    for i := range nums {
        if i != 0{
            leftSum[i] = leftSum[i-1] + nums[i-1]
            rightSum[len(nums)-i-1] = rightSum[len(nums)-i] + nums[len(nums)-i]
        }    
    }
    for i := range nums {
        nums[i] = int(math.Abs(float64(leftSum[i] - rightSum[i])))
    }
    return nums
}