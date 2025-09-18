/**
 * @param {number[]} nums
 * @return {number[]}
 */
var smallerNumbersThanCurrent = function(nums) {
    let arr=[]
       let count=0
    for (i of nums){
        for(j of nums){
            if(i > j){
                count++
               
            }
             
        }
        arr.push(count)
        count=0
    }
    return arr
};