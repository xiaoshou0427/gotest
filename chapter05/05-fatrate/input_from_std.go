package main

import "fmt"

type inputFromStd struct {

}

func (inputFromStd) GetInput() Person {
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
	return Person{
		name:   name,
		sex:    sex,
		tall:   tall,
		weight: weight,
		age:    age,
	}
}
