func limitOccurrences(nums []int, k int) []int {
    freq := make(map[int]int)
    result := []int{}
    for _, v := range nums {
        freq[v]++
        if freq[v] <= k {
        result = append(result, v)
        }
    }
    return result
}