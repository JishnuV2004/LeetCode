func maxIceCream(costs []int, coins int) int {
 sort.Ints(costs) 
 res:=0
 for _,n:=range costs{
    if n > coins {
      break
    }
    coins-=n 
    res++
 }   
 return res
}