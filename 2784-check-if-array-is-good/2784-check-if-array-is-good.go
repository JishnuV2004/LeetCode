func isGood(nums []int) bool {
    freq := make(map[int]int)
    count := 0
    for i:=0; i<len(nums); i++ {
        if nums[i] > len(nums){
            return false
        }
        freq[nums[i]]++
        if freq[nums[i]] == 2{
            count++
        }
    }
    if count > 1 {
        return false
    }
    for key, val := range freq {
        if len(nums)-1 == key && val == 2 {
            return true
        }
    }
    return false
}