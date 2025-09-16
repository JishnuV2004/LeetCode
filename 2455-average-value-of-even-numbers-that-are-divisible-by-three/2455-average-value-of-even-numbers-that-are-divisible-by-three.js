/**
 * @param {number[]} nums
 * @return {number}
 */
var averageValue = function(nums) {
    let sum=0
    let count=0
    for(num of nums){
        if(num%2 == 0 && num % 3== 0){
            sum=sum+num
            count++
        }
    }
    if(sum>0){
        return Math.floor(sum/count)
    }else{
        return 0
    }
};