/**
 * @param {string[]} operations
 * @return {number}
 */
var finalValueAfterOperations = function(operations) {
    let val=0
    for (i of operations){
        switch(i){
            case "--X" : --val
            break;
            case "X--" : val--
            break;
            case "++X" : ++val
            break;
            case "X++" : val++
            break;
            default : i
        }
    }
    return val
};