func vowelStrings(words []string, left int, right int) int {
    count := 0
    vowel := "aeiou"
    for i:=left; i<=right; i++ {
        increment := 0
        word := words[i]
        for i := range vowel {
            first := word[0]
            last := word[len(word)-1]
            if first == vowel[i] {
                increment++
            }
            if last == vowel[i] {
                increment++
            }
        }
        if increment == 2 {
                count++
            }
    }
    return count
}