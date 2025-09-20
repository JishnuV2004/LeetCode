/**
 * @param {number[]} nums
 * @return {number}
 */
var singleNumber = function(nums) {
    let sort=nums.sort()
    for(let i=0; i<sort.length; i+=2){
        if(sort[i] !== sort[i+1]){
            return sort[i]
        }
    }
    return uniq
};