func maximumDifference(nums []int) int {
    diff := 0
    for i:=0; i<len(nums)-1; i++ {
        for j:=i+1; j<len(nums); j++ {
            if nums[i] < nums[j] && diff < nums[j] - nums[i] {
                diff = nums[j] - nums[i]
            }
        }
    }
    if diff > 0 {
        return diff
    }
    return -1
}