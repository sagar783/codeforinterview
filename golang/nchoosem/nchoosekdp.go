package main

import "fmt"

var arr [51][51]int

func main() {
	pascaltriangle(&arr)
	var n int
	var m int
	fmt.Println("enter m and n")
	fmt.Scanf("%d",&m)
	fmt.Scanf("%d",&n)
	if n > m{
	fmt.Println(0)
	}else if n< 0{
		fmt.Println(0)
	}else{
		fmt.Println(arr[m+1][n+1])

	}
	

}

func pascaltriangle(arr *[51][51]int) {
	for i := 0; i <= 50; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				arr[i][j] = 0
			} else if j == i {
				arr[i][j] = 1
			} else {
				arr[i][j] = arr[i-1][j-1] + arr[i-1][j]
			}

		}

	}

}
