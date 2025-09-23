/**
 * @param {number[]} nums
 * @return {number}
 */
var differenceOfSum = function(nums) {
    let sum1=0
    let sum2=0
    result=nums.join("").split("").map(Number)
    for(num1 of result){
        sum1+=num1
    }
    for(num2 of nums){
        sum2+=num2
    }
    return sum2-sum1
};