func smallerNumbersThanCurrent(nums []int) []int {
    arr:=[]int{}
    for _,val1:= range nums{
        count:=0
        for _,val2:= range nums{
            if val1 > val2{
                count++
            }
        }
        arr=append(arr,count)
    }
    return arr
}