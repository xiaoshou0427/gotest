package main

import "fmt"

func main() {
	//实例化结构体
	frSvc := &fatRateService{s: getFatRateSuggestion()}
	//获取一个人的信息测试
	fakePerson := getFakePersonInfo()
	fmt.Println(frSvc.GiveSuggestionToPerson(fakePerson))
	// 永远循环下去：
	for {
		p:=getPersonInfoFromInput()
		fmt.Println(frSvc.GiveSuggestionToPerson(p))
	}
}


//命令行录入真人的信息：*person 打开盒子将内容放进去！ 指针类型的！用指针节省资源？
func getPersonInfoFromInput() *Person {
	var name string
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)
	var sex string
	fmt.Print("请输入性别（男/女）：")
	fmt.Scanln(&sex)
	var tall float64
	fmt.Print("请输入身高（单位：米，例如：1.70）：")
	fmt.Scanln(&tall)
	var weight float64
	fmt.Print("请输入体重（单位：千克，例如：71.5)：")
	fmt.Scanln(&weight)
	var age int
	fmt.Print("请输入年龄：")
	fmt.Scanln(&age)
	return &Person{
		name:   name,
		sex:    sex,
		tall:   tall,
		weight: weight,
		age:    age,
	}
	//return 一定是指向某个值给到 *Person 那么就是指向Pserson的内存地址，所以才用&符号指向person 与*person呼应
	// person 里面的bmi和fatrate 是计算出来的，不用写在这里
}

//录入一个假的信息，这个在ut的时候调用这个方法来做剩下所有的计算了：
//参考：https://zhuanlan.zhihu.com/p/46673861
//再理解一下 *Person 类似于 *int，*float 是一个指针类型的Person，预期产出是一个指针类型的Persion!
//就像你定义函数的时候，预期产出可以是一个int 或者float 类似，这里预期产出是一个指针类型的结构体Person!
//而&Person{xxx: xxxx} ===> & Person{xxx:xxxx} return返回一个指针变量，而这个指针变量指向Person{xxx: xxxx}的内存地址
//而预期产出是一个指针类型的Person ,也就是指向Person的内存地址，符合预期，返回这个已经给成员变量赋值的结构体的内存地址！
func getFakePersonInfo() *Person {
	return &Person{
		name:   "小强",
		sex:    "男",
		tall:   1.7,
		weight: 70,
		age:    35,
	}
}

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
