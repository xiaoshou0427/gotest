package main

import (
	"fmt"
	"math/rand"
)

//归并排序实际上是一个不停递归的过程
//每次都是先分裂，再合并这么一个过程
func mergeSort(arr []int) []int {
	//切分不断的切分
	//这里定义两个新数组，一个是前半部分，一个是后半部分
	left, right := (arr)[:len(arr)/2], (arr)[len(arr)/2:]
	//这是第一次分，后面继续迭代的分，直到分完，进行merge操作
	if len(arr) <= 2 { // 入参的长度 ，就剩下两个数了，已经见底了，可以去merge了！ 这里是arr 不是left和right
		return mergeArr(left, right) //调用下面的函数，进行排序，并返回数组
	} else { //不为2的长度，继续分数组，把左边的继续分，分到小于等于2的长度（递归），然后执行right的分数组，直到小于等于2的长度
		//sortedLeft := mergeSort(left)   //继续迭代，把左边的数组传入到arr进行拆分
		//sortedRight := mergeSort(right) //把右边的数组传入arr进行拆分
		////这里迭代完了，会出现两个数组哦，一个是排列的左边，一个是排列的右边，下面是把者两个数组再排序
		//mergeSortedArr := mergeArr(sortedLeft, sortedRight)
		//return mergeSortedArr //返回结果
		//上面四行可以简化成如下的一行：
		return mergeArr(mergeSort(left), mergeSort(right))

	}
}

//两个排好序的队列，进行合并
func mergeArr(left, right []int) []int {
	out := []int{}
	leftI, rightI := 0, 0
	for {
		//if leftI == len(left) {
		//	break //下面不断的加，这里要给一个范围退出 ，即可索引最大不能超过长度的len-1
		//} //当等于这个长度了，就已经超出范围了，可以break掉了！
		//if rightI == len(right) {
		//	break //总会有一个先到头
		//}
		//精简写法：
		if leftI == len(left) || rightI == len(right) {
			break //总会有一个先到头
		}
		if left[leftI] < right[rightI] {
			out = append(out, left[leftI])
			leftI++  //左边小于右边的，那就追加进去，并且左边索引+1
			continue //继续进入下一次for循环，就是拿左边的所有数字跟右边的比较
		} else {
			//如果上述不成立，也就是右边的小于左边的，那就把右边的追加进去
			out = append(out, right[rightI])
			rightI++ //不断跟右边的比
			continue
		}
	}
	//极限情况 (妙啊)
	//假设左边走到头，右边都没变，或者右边走到头了，左边都没变
	//比如左边一直append完了，这个时候leftI == len(left)，上面break了，那右边的还没进去呢！
	//此时做个for循环，当leftI 小于这个长度，也就是下标不超出范围，才执行下面的第一个for循环
	//很显然，极限情况出现了，那么就不会执行下面第一个for循环。
	//上面break了，右边的咋办嘞，就执行第二个for循环，不断往里加进去！完事！
	//反之亦然！ 右边的都加进去了，左边还没加进去，那就用下面的for ;leftI
	//简单理解就是左手没到头，就把左手追加进去，反之亦然
	for ; leftI < len(left); leftI++ {
		out = append(out, left[leftI])
	}
	for ; rightI < len(right); rightI++ {
		out = append(out, right[rightI])
	}
	return out
}

func main() {
	arrSize := 10
	arr := []int{}
	for i := 0; i < arrSize; i++ {
		arr = append(arr, rand.Intn(50))
	} //50以内 随机加10个数字进去
	fmt.Println("原生状态：", arr)
	sorted := mergeSort(arr)
	fmt.Println("排序状态：", sorted)
}
