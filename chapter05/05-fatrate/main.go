package main

import "fmt"

func main() {
	//实例化结构体
	frSvc := &fatRateService{s: getFatRateSuggestion(),input: &fakeInput{}} // todo

	// 永远循环下去：
	for {
		p := frSvc.input.GetInput()
		fmt.Println(frSvc.GiveSuggestionToPerson(&p))
	}
}



//录入一个假的信息，这个在ut的时候调用这个方法来做剩下所有的计算了：
//参考：https://zhuanlan.zhihu.com/p/46673861
//再理解一下 *Person 类似于 *int，*float 是一个指针类型的Person，预期产出是一个指针类型的Persion!
//就像你定义函数的时候，预期产出可以是一个int 或者float 类似，这里预期产出是一个指针类型的结构体Person!
//而&Person{xxx: xxxx} ===> & Person{xxx:xxxx} return返回一个指针变量，而这个指针变量指向Person{xxx: xxxx}的内存地址
//而预期产出是一个指针类型的Person ,也就是指向Person的内存地址，符合预期，返回这个已经给成员变量赋值的结构体的内存地址！
//func getFakePersonInfo() *Person {
//	return &Person{
//		name:   "小强",
//		sex:    "男",
//		tall:   1.7,
//		weight: 70,
//		age:    35,
//	}
//}

//下面的修改案例更好说明，上面是简写的，预期产出是一个指针类型的结构体（也可以叫指针变量吧？），return 返回真实的值给到预期
// xyz 为指针变量（是指针类型），指向结构体person{小强 男 1.7 70 35 24.221453287197235 0.20565743944636683}的内存地址
//给到person := getFakePersonInfo() 这个调用函数，直接将person 设置为指针类型的变量！
//值为结构体erson{小强 男 1.7 70 35 24.221453287197235 0.20565743944636683}的内存地址
//func getFakePersonInfo() *Person {
//	xyz := &Person{
//		name:   "小强",
//		sex:    "男",
//		tall:   1.7,
//		weight: 70,
//		age:    35,
//	}
//	return xyz
//}
