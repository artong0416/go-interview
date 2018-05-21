/*
作者：artong0416
时间：2018/5/21
功能：快排的多协程实现
*/
package sort


func QuikSortChan(array []int, ch chan int)  {
	if len(array) == 1 {
		ch <- array[0]
		close(ch)
		return
	}
	if len(array) == 0 {
		close(ch)
		return
	}
	leftarr := make([]int, 0)
	rightarr := make([]int, 0)
	basenum := array[0]
	newarray := array[1:]
	for _,v := range newarray{
		if v <= basenum {
			leftarr = append(leftarr, v)
		} else {
			rightarr = append(rightarr, v)
		}
	}
	leftch := make(chan int, len(leftarr))
	rightch := make(chan int, len(rightarr))
	go QuikSortChan(leftarr, leftch)
	go QuikSortChan(rightarr, rightch)
	for i := range leftch{
		ch <- i
	}
	ch <- basenum
	for i := range rightch{
		ch <- i
	}
	close(ch)
	return
}