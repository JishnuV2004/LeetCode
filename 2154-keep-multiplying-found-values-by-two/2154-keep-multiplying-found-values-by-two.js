/**
 * @param {number[]} nums
 * @param {number} original
 * @return {number}
 */
var findFinalValue = function(nums, original) {
    let mult=original
    for(let i=0; i<nums.length; i++){
        for(let j=0; j<nums.length; j++){
            if (nums[j] == mult){
                mult=mult*2
            }
        }
    }
    return mult
};