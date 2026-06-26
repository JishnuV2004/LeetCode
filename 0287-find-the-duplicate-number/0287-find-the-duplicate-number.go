func findDuplicate(nums []int) int {
    freq := make(map[int]int)
    for _, v := range nums {
        freq[v]++
        if freq[v] > 1 {
            return v
        }
    }
    return 0
}