/**
 * @param {number[]} nums
 * @return {number[]}
 */
var sortedSquares = function(nums) {
    let sqr=nums.map((val)=> val*val).sort((a,b)=>a-b)
    return sqr
};