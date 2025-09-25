/**
 * @param {number[]} nums
 * @return {number[]}
 */
var runningSum = function(nums) {
    let temp=0
    let arr=[]
    for(num of nums){ 
        temp+=num
        arr.push(temp)
    }
    return arr
};