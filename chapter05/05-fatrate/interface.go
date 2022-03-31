package main

//创建一个录入的接口
type InputService interface {
	GetInput() Person
}

type OutputService interface {
	Output(Person, string) //保存人的信息和建议信息
}
