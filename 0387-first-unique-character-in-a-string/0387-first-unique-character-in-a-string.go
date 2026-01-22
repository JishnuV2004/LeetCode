func firstUniqChar(s string) int {
    x := len(s)
    freq := make(map[string]int)
    for i:=0; i<x; i++ {
        freq[string(s[i])]++
    }
    for i:=0; i<x; i++ {
        if freq[string(s[i])] == 1 {
            return i
        }
    }
    return -1
}