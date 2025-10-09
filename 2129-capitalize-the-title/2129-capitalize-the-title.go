func capitalizeTitle(title string) string {
    words:=strings.Split(title , " ")
    for i,word:= range words{
        word=strings.ToLower(word)
        if len(word)>2{
            word=strings.Title(word)
        }
        words[i]=word
    }
    return strings.Join(words," ")
}