/**
 * @param {number[]} nums
 * @return {number[]}
 */
var decompressRLElist = function(nums) {
    let arr=[]
    for (let i=0; i<nums.length; i+=2){
        let temp=nums[i+1]
        for(let j=0; j<nums[i]; j++){
            arr.push(temp)
        }
    }
    return arr
};