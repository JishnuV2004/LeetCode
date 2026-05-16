func findMin(nums []int) int {
    for i:=0; i<len(nums); i++ {
        swapped := false
        for j:=0; j<len(nums)-1; j++ {
            if nums[j] > nums[j+1] {
                nums[j], nums[j+1] = nums[j+1], nums[j]
                swapped = true
            }
            
        }
        if !swapped {
                break
            }
    }
    return nums[0]
}