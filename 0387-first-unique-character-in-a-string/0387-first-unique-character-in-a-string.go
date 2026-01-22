func firstUniqChar(s string) int {
    x := len(s)
    r := []rune(s)
    freq := make(map[rune]int)
    for i:=0; i<x; i++ {
        freq[r[i]]++
    }
    for i:=0; i<x; i++ {
        if freq[r[i]] == 1 {
            return i
        }
    }
    return -1
}