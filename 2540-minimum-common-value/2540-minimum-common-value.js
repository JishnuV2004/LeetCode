/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
var getCommon = function(nums1, nums2) {
    for(num1 of nums1){
        for(num2 of nums2){
        if( num1===num2 ){
            return num1
        }
    }
    }
    return -1
};