/**
 * @param {number[]} nums
 * @return {number}
 */
var differenceOfSum = function(nums) { 
    let sum2=nums.reduce((a,b)=>a+b)
    let result=nums.join("").split("").map(Number)
    let sum1=result.reduce((a,b)=>a+b)
    return sum2-sum1
};