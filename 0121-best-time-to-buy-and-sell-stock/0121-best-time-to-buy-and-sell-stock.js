/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function(prices) {
    min=prices[0]
    profit=0
    for(n of prices){
        if( n<min){
            min=n
        }else if(n-min > profit){
            profit=n-min
        }
    }
    return profit
};