package main

import (
	"fmt"
	"time"
)

//计数
var totalCompare int = 0

func main() {
	/*	//size := 1000
		//arr := sampleData
		//fmt.Println(arr)
		//1。 501 在不在里面？
		//2。888 在不在？
		//3。900 在不在？
		//4. 3 在不在*/

	//复制一份新的array，避免修改原始数组,这里要加...
	newArr := append([]int64{}, sampleData...)

	startTime := time.Now() //时间应该放在排序前面

	quickSort(&newArr, 0, len(newArr)-1) //这里要减一
	// todo 排序

	for i := 0; i < 1000000; i++ { //跑这么多次查询 //185.184094ms
		search(&newArr, 501)
		search(&newArr, 888)
		search(&newArr, 900)
		search(&newArr, 3)
		search(&newArr, 2472)
	}
	finishTime := time.Now()
	fmt.Println("查找完毕的总次数：", totalCompare)
	fmt.Println("总用时：", finishTime.Sub(startTime))
}

//arrP 不传arr过来，而是传arr的指针过来！ 注意边界问题！ 要减一
func search(arrP *[]int64, targetNum int64) bool {
	return searchHarf(arrP, 0, len(*arrP)-1, targetNum)
}

//这里调用下面的函数，从0到arrP的长度

//二分搜索功能，找数组的中间数字进行操作！
func searchHarf(arrP *[]int64, left, right int, targetNum int64) bool {
	middleIndex := (left + right) / 2 //中间的坐标
	data := (*arrP)[middleIndex]      //中间坐标的值

	totalCompare++ //增加计数器数值，每次比较的时候计数

	if data < targetNum {
		if left == right {
			return false //极端情况，已经到最右边了，还是没找到
		}
		return searchHarf(arrP, middleIndex+1, right, targetNum)
		//如果拿到的中间的数字比目标数字小(前提是已经俳好序，那就去一半的右边继续找)
		//这里进行了递归，这里return bool，这里加1了，从这个中间数字右边开始！
	} else if data > targetNum {
		if left == right {
			return false //极限情况，已经到了最左边，还是没找到
		}
		return searchHarf(arrP, left, middleIndex-1, targetNum)
		//这里没有做减一处理，因为左边最后一个值的时候减一就out of range了
		//这里也要减一
	} else {
		return true //不大不小，刚刚好 return true
	}
}

//这里面存在边界问题，一直往左或者一直往右！ 记得要减1 len(*arrP)-1，函数middleIndex+1/-1
