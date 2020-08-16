//this is memoized version of n choose k this program will work for m and n = 5000
package main

import "fmt"

var arr [5001][5001]int

func main() {
	var n int
	var m int
	fillelement(&arr)
	fmt.Println("enter m and n")
	fmt.Scanf("%d", &m)
	fmt.Scanf("%d", &n)
	fmt.Println(recur(m, n))
}
func recur(m int, n int) int {
	if n == 1 {
		return m
	}
	if m == n || n == 0 {
		return 1
	}
	if arr[m][n] == -1 {
		val := recur(m-1, n-1) + recur(m-1, n)
		arr[m][n] = val
		return val
	} else {
		return arr[m][n]
	}
}
func fillelement(arr *[5001][5001]int) {
	for i := 0; i <= 5000; i++ {
		for j := 0; j <= 5000; j++ {
			arr[i][j] = -1
		}
	}
}
