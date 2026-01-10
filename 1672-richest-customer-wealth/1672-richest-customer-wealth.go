func maximumWealth(accounts [][]int) int {
    richest := 0
    for i:=0; i<len(accounts); i++ {
        count:= 0
        for _, v := range accounts[i] {
            count += v
            if richest < count {
                richest = count
            }
        }
    }
    return richest
}