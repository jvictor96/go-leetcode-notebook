package main

import (
    "fmt"
)

func triangle(nums []int) int {
    for i:= len(nums) - 1;i > 0; i-- {
		fmt.Println(i)
	}
	return 1
}

func main() {
    fmt.Println(triangle([]int{1,2,2,5,1,3}))
}