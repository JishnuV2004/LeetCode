func gcdOfOddEvenSums(n int) int {
    sumodd:=0
    sumeven:=0
    for i:=1; i<=n; i++ {
            sumodd += (i*2)-1
            sumeven += i*2
    }
    return sumeven - sumodd
}