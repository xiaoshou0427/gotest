package main

import (
	"fmt"
	"time"
)

func main() {
	guess(1, 100)
	fmt.Println("done calc, sleep somewhile")
	time.Sleep(10 * time.Second) //睡10秒？ why？ 这种写法
}

func guess(left, right uint) {
	guessed := (left + right) / 2 //给定个标准，你心中所想与这个进行对比！二分法（定标）
	var getFromInput string
	fmt.Println("人工智能猜测是：", guessed)
	fmt.Print("如果高了，输入1；如果低了，输入0；猜对了，输入9: ")
	fmt.Scanln(&getFromInput)
	switch getFromInput {
	case "1":
		if left == right { //防止你耍赖
			fmt.Println("你是不是改变注意了？")
			return
		}
		guess(left, guessed-1) //猜高了，就从你猜的值往下，定义新的范围,递归到函数
	case "0":
		if left == right { //防止你耍赖
			fmt.Println("你是不是改变注意了？")
			return
		}
		guess(guessed+1, right) //猜低了
	case "9":
		fmt.Println("你心里想的数字是：", guessed)
	}
}
