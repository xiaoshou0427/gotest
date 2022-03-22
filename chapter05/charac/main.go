package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var refrigerator Refrigerator
	//fmt.Println(refrigerator.Size)
	//
	//var elephant Elephant
	//fmt.Println(elephant.Name)
	//
	//var putER PutElephantIntoRefrigerator //默认为nil
	//putER.OpenTheDoorOfRefrigerator(refrigerator)

	var refrigerator Refrigerator
	var elephant Elephant
	var legend PutElephantIntoRefrigerator = &PutElephantIntoRefrigeratorImpl{} //指向实现
	legend.OpenTheDoorOfRefrigerator(refrigerator)
	legend.PutElephantIntoRefrigerator(elephant,refrigerator)
	legend.CloseTheDoorOfRefrigerator(refrigerator)
	// todo show the elephant in refrigerator

	var c Close = Refrigerator{}
	var b Box = Refrigerator{}
	fmt.Println(b,c)
	c = b //从范围小的向范围大的转 会成功
	//b = c //从范围大的向范围小的转 会失败

	var i interface{}
	i = 3
	fmt.Println(reflect.TypeOf(i),"value:",i)
	i = 3.3234
	fmt.Println(reflect.TypeOf(i),"value:",i)
	i = Refrigerator{}
	fmt.Println(reflect.TypeOf(i),"value:",i)
}

//这是一个接口
type PutElephantIntoRefrigerator interface {
	OpenTheDoorOfRefrigerator(Refrigerator) error
	PutElephantIntoRefrigerator(Elephant,Refrigerator ) error
	CloseTheDoorOfRefrigerator(Refrigerator) error
}

type TestTypeImplInterface func() //定义一个函数类型的TestTypeImplInterface 对象

func (t TestTypeImplInterface) OpenTheDoorOfRefrigerator(_ Refrigerator) error {
	return nil
}
func (t TestTypeImplInterface) PutElephantIntoRefrigerator(_ Elephant,_ Refrigerator) error {
	return nil
}
func (t TestTypeImplInterface) CloseTheDoorOfRefrigerator(_ Refrigerator) error {
	return nil
}

//定义一个对象
type PutElephantIntoRefrigeratorImpl struct {
}

//对象的成员函数
func (legent *PutElephantIntoRefrigeratorImpl) OpenTheDoorOfRefrigerator(refrigerator Refrigerator) error {
	// todo
	fmt.Println("打开冰箱们")
	return nil
}
func (legent *PutElephantIntoRefrigeratorImpl) PutElephantIntoRefrigerator(elephant Elephant, refrigerator Refrigerator) error {
	// todo
	fmt.Println("将大象放入冰箱")
	return nil
}
func (legent *PutElephantIntoRefrigeratorImpl) CloseTheDoorOfRefrigerator(refrigerator Refrigerator) error {
	// todo
	fmt.Println("关闭冰箱")
	return nil
}


type Open interface {
	Open() error
}
type Close interface {
	Close() error
}

type Box interface {
	Open
	Close
}


type Refrigerator struct {
	Size string
}

func (Refrigerator) Open() error {
	return nil
}

func (Refrigerator) Close() error {
	return nil
}

type Elephant struct {
	Name string
}

