func minimumCost(cost []int) int {
    sort.Slice(cost, func (a, b int) bool {
        return cost[a] > cost[b]
    })
    n := len(cost)
    count :=0

    for i:=0; i<n; i+=3 {
        count += cost[i]
        if i+1 < n {
            count += cost[i+1]
        }
    }
    return count
}