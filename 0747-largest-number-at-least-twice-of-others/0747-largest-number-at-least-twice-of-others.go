func dominantIndex(nums []int) int {
    large := 0
    index := 0
    for i:=0; i<len(nums); i++ {
        if nums[i] > large {
            large = nums[i]
            index = i
        }
    }
    for j:=0; j<len(nums); j++ {
        if nums[j] != large && nums[j]*2 > large {
            return -1
        }
    }
    return index
}