func sortPeople(names []string, heights []int) []string {
    arr := make([]string, 0, len(names))
    freq := make(map[int]string)

    for i, v := range names {
        freq[heights[i]] = v
    }

    sort.Slice(heights,func(i,j int)bool{
        return heights[i]>heights[j]
    })

    for _,v:= range heights{
        arr=append(arr,freq[v])
    }
    return arr
}