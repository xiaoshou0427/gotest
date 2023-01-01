package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

//计数
var totalCompare int = 0

func main() {
	//size := 1000
	arr := sampleData //generateRandomData(size)
	//fmt.Println(arr)
	//1。 501 在不在里面？
	//2。888 在不在？
	//3。900 在不在？
	//4. 3 在不在
	startTime := time.Now()
	for i := 0; i < 1000000; i++ { //跑这么多次查询
		search(&arr, 501)
		search(&arr, 888)
		search(&arr, 900)
		search(&arr, 3)
		search(&arr, 2472)
	}
	finishTime := time.Now()
	fmt.Println("查找完毕的总次数：", totalCompare)
	fmt.Println("总用时：", finishTime.Sub(startTime))
}

//顺序查找
//arrP 不传arr过来，而是传arr的指针过来！
func search(arrP *[]int64, targetNum int64) bool {
	for _, v := range *arrP {
		totalCompare++ //每次查找都计数一次
		if v == targetNum {
			return true
		}
	}
	return false
}

func generateRandomData(size int) []int64 {
	arr := make([]int64, 0, size)

	for i := 0; i < size; i++ {
		//rand.Reader 可以理解为生成器，在数据类型为int64  50内随机生成数字赋值给i，error不管
		i, _ := rand.Int(rand.Reader, big.NewInt(3000)) //50以内随机int

		arr = append(arr, i.Int64()) //类型都要是int64
	}
	return arr
}
