func containsNearbyDuplicate(nums []int, k int) bool {
    freq := make(map[int]int)
    for i:=0; i<len(nums); i++ {
        _, ok := freq[nums[i]]
        if !ok {
            freq[nums[i]] = i
            continue
        } else if i - freq[nums[i]] <= k {
            return true
        } else {
            freq[nums[i]] = i 
        }
    }
    return false
}