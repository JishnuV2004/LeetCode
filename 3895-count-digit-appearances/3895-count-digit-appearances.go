func countDigitOccurrences(nums []int, digit int) int {
    count := 0
    strDigit := strconv.Itoa(digit)
    for _, num := range nums {
        str := strconv.Itoa(num)
        for _, v := range str {
            if string(v) == string(strDigit) {
                count++
            }
        }
    }
    return count
}