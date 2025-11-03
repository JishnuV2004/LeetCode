func largestAltitude(gain []int) int {
    max := 0
    temp := 0
    for _,val:= range gain{
        temp += val
        if max < temp {
            max= temp
        }
    }
    return max
}