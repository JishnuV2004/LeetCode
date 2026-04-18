func mirrorDistance(n int) int {
    currentNvalue:=n
    reversed := 0
    for n != 0 {
        reversed = reversed * 10 + n % 10
        n/=10
    }
    sum := reversed - currentNvalue
    if sum < 0 {
        return -sum
    }
    return sum
}