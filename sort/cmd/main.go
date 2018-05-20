/*
时间：2018/5/20
功能：
*/
package main

import (
	"github.com/artong0416/go-interview/sort"
	"fmt"
)

func main()  {
	arr :=[]int{9,12,3,14,28,9,2,4,6,89}
	sort.QuickSort(arr, 0, len(arr) - 1)
	fmt.Println(arr)
}
