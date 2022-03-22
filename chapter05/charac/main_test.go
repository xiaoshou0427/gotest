package main

import (
	"fmt"
	"testing"
)

func TestAssertion(t *testing.T) {
	r := TestBox{}  //你这有个接口实现，有个方法close 也没错！
	var c Close = r //这个定义也没有问题，也有实现，但是下面的r2就不一定是你想要的
	//上述等同于 var c Close = TestBox{}
	//switch cDetail := c.(type) { // 定义一个变量等于c的类型
	//case Refrigerator: //判断cDetail 的值是不是Refrigerator（Refrigerator是个类型）
	//	fmt.Println("是预期的冰箱")
	//	fmt.Println(cDetail.Size) //是这个类型就print size
	//case Box: //类型是不是Box
	//	fmt.Println("这是一个box，不能当冰箱用")
	//case TestBox: //用Close结果也是一样的，TestBox是Close接口的实现
	//	fmt.Println("这是一个TestBox", cDetail)
	//	//这个判断的过程就叫做断言
	//}
	refrigerator,ok := checkIsRefrigerator(c) //跑一下，看c是不是想要的结果！此时c并不是Refrigerator
	if ok {
		fmt.Println("是个冰箱，开门装大象",refrigerator)
	}else {
		fmt.Println("这不是一个冰箱",refrigerator)
	}
}

func checkIsRefrigerator(c Close) (Refrigerator, bool) {
	r, ok := c.(Refrigerator) //断言，判断类型
	return r, ok //这样就拿到是不是冰箱
}

type TestBox struct {
}

func (tb TestBox) Close() error {
	return nil
}
