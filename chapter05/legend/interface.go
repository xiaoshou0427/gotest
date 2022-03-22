package main

//这是一个接口，里面的函数是没有变量名字的，只是！！！类型！！！ 这些类型都是结构体，对象！
type PutElephantIntoRefrigerator interface {
	OpenTheDoorOfRefrigerator(Refrigerator) error
	PutElephantIntoRefrigerator(Elephant,Refrigerator ) error
	CloseTheDoorOfRefrigerator(Refrigerator) error
}

//定义结构体
type Refrigerator struct {
	Size string
}

type Elephant struct {
	Name string
}


