func findMaxConsecutiveOnes(nums []int) int {
    count := 0
    c := 0
    for _, v := range nums {
        if v == 1 {
            c++
            if c > count {
                count = c
            }
        }else{
            c = 0
        }
    }
    return count
}