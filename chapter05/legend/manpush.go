package main

import (
	"fmt"
)

//用人力装

type manLegend struct {
}

//这里没有定义结构体内部的成员，可以不用关联结构体时命名，并且入参也只是类型，而没有变量！！！
func (*manLegend) OpenTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用 manLegend 做 OpenTheDoorOfRefrigerator")
	return nil
}
func (*manLegend) PutElephantIntoRefrigerator(Elephant, Refrigerator) error {
	fmt.Println("用 manLegend 做 PutElephantIntoRefrigerator")
	return nil
}
func (*manLegend) CloseTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用 manLegend 做 CloseTheDoorOfRefrigerator")
	return nil
}