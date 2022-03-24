package main

import (
	"fmt"
	"testing"
)

type Change interface {
	ChangeName(newName string)
	ChangeAge(newAge int)
}

type Student struct {
	Name string
	Age int
}

////都在指针上
//func (s *Student) ChangeName(newName string)  { //理解为：定义s指针变量，指向结构体的内存地址（编号）
//	s.Name = newName
//}
//
////都在指针上
//func (s *Student) ChangeAge(newAge int)  { //理解为：定义s指针变量，指向结构体的内存地址（编号）
//	s.Age = newAge
//}

//都在对象上
func (s Student) ChangeName(newName string)  { //理解为：定义s指针变量，指向结构体的内存地址（编号）
	s.Name = newName
}

//都在对象上
func (s Student) ChangeAge(newAge int)  { //理解为：定义s指针变量，指向结构体的内存地址（编号）
	s.Age = newAge
}

func TestVal(t *testing.T) {
	var stdChg Change
	stdChg = &Student{ //实例化接口
		Name: "Tom", //这两个值跟成员函数没有任何关系，其实没有调用成员函数！
		Age: 0,
	}
	//stdChg = Student{} //实例化接口
	fmt.Println(stdChg)
}

//func TestName(t *testing.T) {
//	s := Student{Name: "Tom"}
//	fmt.Println(s.Name)
//	s.ChangeName("Jerry") //调用成员函数修改结构体的成员变量
//	fmt.Println(s.Name)
//}
