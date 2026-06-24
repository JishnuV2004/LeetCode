func digitFrequencyScore(n int) int {
    sum := 0
    str := strconv.Itoa(n)
    for _, v := range str {
        num := int(v - '0')
        sum += num
    }
    return sum
}