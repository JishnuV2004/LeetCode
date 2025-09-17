/**
 * @param {string} s
 * @return {string}
 */
var reverseWords = function(s) {
    let arr=s.split(" ")
    let str=[]
    for(text of arr){
        str.push(text.split("").reverse().join(""))
    }
    return str.join(" ")
};