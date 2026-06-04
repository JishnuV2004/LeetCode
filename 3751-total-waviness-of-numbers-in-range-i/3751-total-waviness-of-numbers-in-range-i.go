func totalWaviness(num1 int, num2 int) int {
    count := 0
    for i:=num1; i<=num2; i++ {
        intStr := strconv.Itoa(i)
        for j:=1; j<len(intStr)-1; j++ {
            if intStr[j-1] < intStr[j] && intStr[j+1] < intStr[j] {
                count++
            } else if intStr[j-1] > intStr[j] && intStr[j+1] > intStr[j]{
                count++
            }
        }
    }
    return count
}