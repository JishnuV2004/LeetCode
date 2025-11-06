func runningSum(nums []int) []int {
    result:=[]int{}
    sum:=0
    for _,val:= range nums{
        sum += val
        result=append(result, sum)
    }
    return result
}