package main

import (
	"fmt"
	"math/rand"
)

//用指针的方式可以对原来的数组进行修改！
func bubble(arr *[]int) {
	for i := 0; i < len(*arr); i++ { //做循环进行比较
		for j := 0; j < len(*arr)-i-1; j++ {
			//这里len(*arr)-i-1 ，其中-1 的目的是下面j+1了，已经到最后一个了
			//比如一共10个数字，到arr[8] 和arr[9]对比的时候，已经比完了，所以这里减1是避免j超出index range
			//减i 这个动作是 每轮找两个数字进行对比，当第一轮比完之后最后一个数是固定，也是最大的，就不用再比较了
			//所以每次外循环到内循环 就会把大的数都固定下来，所以减i就不会再去对比之前固定下来的大数字
			//当然这里不减i，也是可以的，只是后面会做几个无效的对比！！！！
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j] //做二元交换
			}
		}
		fmt.Println("中间状态：", *arr)
	}
	fmt.Println("最终状态", *arr)
}

func main() {
	arrSize := 10
	arr := []int{}
	for i := 0; i < arrSize; i++ {
		arr = append(arr, rand.Intn(50))
	} //50以内 随机加10个数字进去
	fmt.Println("原生状态：", arr)
	bubble(&arr)
}
