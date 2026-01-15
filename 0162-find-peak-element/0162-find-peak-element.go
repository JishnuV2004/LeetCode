func findPeakElement(nums []int) int {
    index := 0
    for i, v := range nums {
        if nums[index] < v {
            index = i
        }
    }
    return index
}