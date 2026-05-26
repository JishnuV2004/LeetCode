func concatWithReverse(nums []int) []int {
    n := len(nums)
    arr := make([]int, n*2)

    for i:=0; i<n; i++ {
        arr[i] = nums[i]
        arr[n*2-1-i] = nums[i]
    }
    return arr
}