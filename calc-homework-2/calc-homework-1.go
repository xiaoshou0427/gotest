package main

import (
	"fmt"
)

func main() {
	var p1x, p2x, p3x, p4x, p1y, p2y, p3y, p4y float64
	p1x, p2x, p1y, p2y = getLineInfo() //调用的时候，定义的每个变量，都从函数的每个预期（实参）中获得！
	p3x, p4x, p3y, p4y = getLineInfo()
	k1 := calculateK(p3x, p4x, p1y, p2y) //(p2y - p1y) * (p4x - p3x) 这里看好了！别写错了
	k2 := calculateK(p1x, p2x, p3y, p4y) //(p2x - p1x) * (p4y - p3y) 这里乘法，第一条线和第二条线的值相乘
	getResult(k1, k2)

}

func getLineInfo() (x1, x2, y1, y2 float64) { //预期是多个变量，并且float64，通过内部scanln 获取输入信息，给到这些变量
	fmt.Print("请输入直线的第一个点的x：")
	fmt.Scanln(&x1)
	fmt.Print("请输入直线的第一个点的y：")
	fmt.Scanln(&y2)
	fmt.Print("请输入直线的第二个点的x：")
	fmt.Scanln(&x2)
	fmt.Print("请输入直线的第二个点的y：")
	fmt.Scanln(&y2)
	return x1, x2, y1, y2 //返回内容给到变量！
}

func calculateK(x1, x2, y1, y2 float64) float64 {
	return (y2 - y1) * (x2 - x1) //这里返回的乘法，视频返回的是除法，没有区别，只是在调用的时候，投入（形参）要选好
}

func getResult(k1, k2 float64) {
	if k1 == k2 {
		fmt.Println("两条直线平行！")
	} else {
		fmt.Println("两条直线不平行")
	}
}
