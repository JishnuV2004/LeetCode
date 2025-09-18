/**
 * @param {number[]} nums
 * @return {number[]}
 */
var smallerNumbersThanCurrent = function(nums) {
    let arr=[]
    for (i of nums){
        let count=0
        for(j of nums){
            if(i > j){
                count++
            }
        }
        arr.push(count)
    }
    return arr
};