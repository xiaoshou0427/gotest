package main

import "fmt"

func panicAndRecover() {
	//defer coverPanicUpgrade()
	defer func() {
		coverPanicUpgrade()
	}()
	//defer coverPanic()
	var nameScore map[string]int = nil
	nameScore["锦文"] = 100
	fmt.Println(test)
	fmt.newPrinter
}

func coverPanic() { //未能抓出panic
	func() {
		if r := recover(); r != nil {
			fmt.Println("系统出严重故障： ", r)
		}
	}()
}

func coverPanicUpgrade() { //这种是可以抓住严重错误的
	if r := recover(); r != nil {
		fmt.Println("系统出严重故障： ", r)
	}
}
