func increasingTriplet(nums []int) bool {
    first := int(^uint(0) >> 1)  
    second := int(^uint(0) >> 1)  
    for _, n := range nums {
        if n <= first {
            first = n          
        } else if n <= second {
            second = n         
        } else {
            return true        
        }
    }
    return false
}
