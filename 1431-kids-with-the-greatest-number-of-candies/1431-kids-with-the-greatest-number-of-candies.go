func kidsWithCandies(candies []int, extraCandies int) []bool {
    max := slices.Max(candies)
    var arr []bool
    for _,val:= range candies{
        if val + extraCandies >= max {
            arr=append(arr,true)
        }else {
            arr=append(arr,false)
        }
    }
    return arr
}