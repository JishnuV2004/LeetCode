func maxIceCream(costs []int, coins int) int {
 sort.Ints(costs) 
 res:=0
 for _,n:=range costs{
    if coins-n >= 0{
        coins-=n 
        res++
    }
 }   
 return res
}