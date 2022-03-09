package main

import "fmt"

func main() {
	//获取个人信息
	person := getFakePersonInfo() //不做一项一项的录入了，假数据测试。接收函数返回的指针函数所指向结构体的内存地址！
	//剩下的是事情就是去做计算和获取建议了
	//开始计算
	c := Calc{}
	c.BMI(person)     //不需要赋值，因为结果不重要，要的是计算好的bmi 写回到person
	c.FatRate(person) //上面已经计算好bmi，写回到person，这时候person有了bmi，计算fatrate
	//调用完毕后，也写回到了person中
	fmt.Println("person:", *person) //person 不是一个指针哦，是一个变量，定义为从函数获取返回结果
	//获取的返回结果是指针类型的Person---*Person意味着打开盒子取出内容，盒子是一个指针类型的结构体，打开结构体指针
	//获得 Person的属性信息，之前c.xxx 已经做了写入操作，bmi和fatrate 就已经有了
	//根据计算结果，给出建议
	sugg := fatRateSuggestion{}
	suggestion := sugg.GetSuggestion(person)
	fmt.Println("Suggestion: ", suggestion)
	//整个逻辑没有变化都是：获取个人信息，计算体脂率，给出建议
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
