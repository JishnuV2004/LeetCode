func createTargetArray(nums []int, index []int) []int {
    arr := []int{}
    for i, v := range index {
        arr = append(arr[:v], append([]int{nums[i]}, arr[v:]...)...,)
    }
    return arr
}