package main

func main()  {
	legendary(&manLegend{},Refrigerator{},Elephant{}) //接口需要用& 指针！

	//fmt.Println("传说中，装大象可以这么装：")
	//这里定义了变量和类型
	//var legend PutElephantIntoRefrigerator = &manLegend{} //这个变量的类型是接口，指向的是（存放的是）对象manLegend的内存地址？
	//var r Refrigerator
	//var e Elephant
	//
	//legend.OpenTheDoorOfRefrigerator(r) //这里r就是一个变量，形参！类型为结构体Refrigerator!
	//legend.PutElephantIntoRefrigerator(e,r)
	//legend.CloseTheDoorOfRefrigerator(r)
	//legendary(legend,r,e) //这个要定义变量
	//fmt.Println("This is a legendary.")

}
