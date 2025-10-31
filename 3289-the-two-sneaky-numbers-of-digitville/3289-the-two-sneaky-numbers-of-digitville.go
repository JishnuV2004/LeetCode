func getSneakyNumbers(nums []int) []int {
    freq := make(map[int]int)
    var result []int
    for _,val:= range nums{
       freq[val]++
       if freq[val] == 2 {
        result= append(result, val)
       }
    }
    return result
}