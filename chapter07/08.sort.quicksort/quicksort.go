package main

import (
	"fmt"
	"math/rand"
)

//有数组以外还需要位置信息，左半边，右半边的位置信息
//给一个数组，左半边，右半边
func quickSort(arr *[]int, start, end int) {
	//todo 确认终止条件，否则将无限递归下去 （当剩下一个数字的时候还需要拆吗？不用了吧 即左边和右边指向同一个数字，就停了吧）

	//选择好支点
	pivotIdx := (start + end) / 2 //获取支点位置，用中间的点位！作为支点！
	pivotV := (*arr)[pivotIdx]    //获取数组的Value

	//左右开始对比,用for 循环
	l, r := start, end
	for l < r { //当左边小于右边，也就是下标左边小于右边
		for (*arr)[l] < pivotV { //左手的值小于支点，也就是小于中间这个值，左边的下一个值再去对比
			l++
		}
		for (*arr)[r] > pivotV { //右手的值大于支点，也就是大于中间的值，右边的笑一个值进行对比
			r--
		}
		//！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！
		//这一段是最后才加的！ 发现了bug
		if l >= r { //这里 下标已经过了中间点，跳出！
			break //
		}
		//！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！
		//如果不满足上面两个条件，意味着左边的值大于中间值，右边的值大于中间值
		//就需要做交换

		(*arr)[l], (*arr)[r] = (*arr)[r], (*arr)[l]
		//然后退到最外面的for循环

		//这一步要做！！！！ 如果不做的话后面程序会卡死！！！
		l++ //
		r-- //
	}
	fmt.Println("l:", l, "r:", r)
	fmt.Println(*arr) //看看处理完是啥样子
	//做递归！！！
	//start = 0 ， end = 4，上面循环完47 这个数字就好了！其他的呢？
	//这里，会变成死循环！加判断
	if l == r {
		l++ //
		r-- //
	} //让r和l 走过去，如果都等于3的话，l++ = 4 ， r-- = 2
	if r > start { //当3 大于0 ，成立，递归！
		quickSort(arr, start, r) //这里递归是（0 3）第四个数字不用管了，【31 37 31 9 】这几个开始
	}
	if l < end { //当5小于4 不成立，不做下面的递归
		quickSort(arr, l, end)
	}
}

//还有一种可能是永远循环下去

func main() {
	arrSize := 100
	arr := []int{}
	for i := 0; i < arrSize; i++ {
		arr = append(arr, rand.Intn(60)) //这里随机数是假的
	} //50以内 随机加10个数字进去
	fmt.Println("原生状态：", arr)
	quickSort(&arr, 0, arrSize-1) //right 只能是4，用变量！！不要用具体的值
	fmt.Println("排序状态：", arr)
}
