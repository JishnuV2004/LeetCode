func isAdjacentDiffAtMostTwo(s string) bool {
    for i:=1; i<len(s); i++ {
        deff := int(s[i]) - int(s[i-1])
        if deff < 0 {
            deff = -deff
        }
        if deff > 2 {
            return false
        }
    }
    return true
}