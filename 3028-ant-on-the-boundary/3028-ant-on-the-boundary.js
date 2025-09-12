/**
 * @param {number[]} nums
 * @return {number}
 */
var returnToBoundaryCount = function(nums) {
    let count=0
    let sum=0
    for(num of nums){
        sum=sum+num
        if(sum === 0){
            count++
        }
    }
    return count
};