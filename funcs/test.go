package main

import "fmt"

var tall uint

//var weight float64

func main() {
	fmt.Println("全局变量赋值前")
	sample()
	//tall, weight = 1.70, 70.0
	fmt.Println("全局变量赋值后")
	sample()
	fmt.Println("重新定义重名的局部变量")
	tall, weight := 100.00, 70.0
	fmt.Println("重新定义重名局部变量：", tall, weight)
	sample()
	fmt.Println("重新赋值局部变量")
	tall, weight = 200, 70.0
	fmt.Println("重新赋值局部变量", tall, weight)
	sample()
	Sample2(tall, weight)

}

func sample() float64 {
	fmt.Println(tall + weight)
	return 0
}

func Sample2(tall, weight float64) float64 {
	fmt.Println(tall + weight)
	return 0
}
