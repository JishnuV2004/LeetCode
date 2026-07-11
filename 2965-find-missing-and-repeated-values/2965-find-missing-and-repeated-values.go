func findMissingAndRepeatedValues(grid [][]int) []int {
    freq := make(map[int]int)
    n := len(grid)
    result := make([]int, 2)

    for _, p := range grid {
        for _, v := range p {
            freq[v]++
        }
    }

    for i:=1; i<=n*n; i++ {
        if freq[i] == 2 {
            result[0] = i
        } else if freq[i] == 0 {
            result[1] = i
        }
    }
    return result
}