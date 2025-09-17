/**
 * @param {string} text
 * @param {string} brokenLetters
 * @return {number}
 */
var canBeTypedWords = function(text, brokenLetters) {
    arr=text.split(" ")
    count=arr.length
    for( word of arr ){
        for(l of brokenLetters)
        if(word.includes(l)){
            count--
            break
        }

    }
    return count
};