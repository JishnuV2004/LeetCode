func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
   count := 0
    for _, a := range arr1 {
        valid := true
        for _, b := range arr2 {
            dff := a - b
            if dff < 0 {
                dff = - dff
            }
            if dff <= d {
                valid = false
                break
            }
        }
        if valid {
            count++
        }
    }
    return count
}