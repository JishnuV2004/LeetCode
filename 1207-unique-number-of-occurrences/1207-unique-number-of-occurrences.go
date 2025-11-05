func uniqueOccurrences(arr []int) bool {
    freq:=make(map[int]int)
    for _,val:= range arr{
        freq[val]++
    }
    count:=make(map[int]bool)
    for _,val:= range freq{
        if count[val]{
            return false
        }
        count[val]=true
    }
    return true
}