func searchRange(nums []int, target int) []int {
    arr := []int{}
    for i, v := range nums {
        if target == v {
            arr = append(arr, i)
        }
    }
    if len(arr) == 1 {
        arr = append(arr, arr[0])
    }else if len(arr) == 0 {
        arr = append(arr, -1, -1)
    }
    a := arr[len(arr)-1]
    arr = arr[:1]
    arr = append(arr, a)
    return arr
}