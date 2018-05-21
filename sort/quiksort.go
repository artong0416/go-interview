/*
时间：2018/5/19
功能：快排(单线程)
*/
package sort

import "fmt"

func QuickSort(array []int, left, right int) {
	//递归退出条件
	if left >= right {
		return
	}
	mid := partition(array, left, right)
	QuickSort(array, left, mid -1)
	QuickSort(array, mid + 1, right)
}

func partition(array []int, left, right int) int {
	//随机选取基准--最左边元素
	midNum := array[left]
	for left < right  {
		//从右向左找第一个小于基准元素的项
		for array[right] > midNum && left < right  {
			right --
		}
		//放到当前未比较的最左
		array[left] = array[right]
		//从左向右找第一个大于基准元素的项
		for array[left] <= midNum && left < right{
			left ++
		}
		//放到当前未比较的最右
		array[right] = array[left]
	}
	//退出循环时候的 right == left 放基准元素
	array[left] = midNum
	fmt.Println(array)
	return left
}