package main
import "fmt"
func main(){
var n int
var m int
fmt.Println("enter m and n")
fmt.Scanf("%d",&m)
fmt.Scanf("%d",&n)
fmt.Println(recur(m,n))
}

func recur(m int,n int) int{
if n == 1{
return m
}
if m == n||n==0{
return 1
}

return recur(m-1,n-1) + recur(m-1,n)
} 