func sortColors(nums []int)  {
    for i:=0; i<len(nums)-1; i++ {
        sorted := false
        for j:=0; j<len(nums)-i-1; j++ {
            if nums[j+1] < nums[j] {
                nums[j+1], nums[j] = nums[j], nums[j+1]
                sorted = true
            }
        }
        if !sorted {
            return
        }
    }
}