func majorityElement(nums []int) int {
    freq := make(map[int]int)
    i:=0
    for i=0; i<len(nums)-1; i++ {
        freq[nums[i]]++
        if freq[nums[i]] > len(nums)/2 {
            break
        }
    }
    return nums[i]
}