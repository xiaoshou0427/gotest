package main

import "fmt"

func legendary(legend PutElephantIntoRefrigerator , r Refrigerator , e Elephant)  {
	fmt.Println("传说中，装大象可以这么装：")
	//todo 如果是人工的操作就给警告
	//利用类型断言来判断特殊实现：如果是某种特定的实现（如人工操作），就给一个警告！！！
	// 这里legend 是legendary方法的形参变量，main调用这个方法，给了个形参变量：&manLegend{}，这个指向结构体manLegend
	//而*manLegend 指向的也是结构体manLegend，两个类型是一样的！断言成功！输出print!
	//这里不用思考legend的类型是个接口类型，这个接口实例化，就是结构体/对象：manLegend
	//相当于var legend PutElephantIntoRefrigerator = &manLegend{} 实例化！而不是接口和对象 做类型比较！
	if _,ok := legend.(*manLegend); ok{
		fmt.Println("WARNING: 现在还在用人工，效率太低！",legend,manLegend{})
	}

	legend.OpenTheDoorOfRefrigerator(r) //这里r就是一个变量，形参！类型为结构体Refrigerator!
	legend.PutElephantIntoRefrigerator(e,r)
	legend.CloseTheDoorOfRefrigerator(r)

	fmt.Println("This is a legendary.")

}
