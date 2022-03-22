package main

import (
	"fmt"
)

//用船装

type shipLegend struct {

}

func (*shipLegend) OpenTheDoorOfRefrigerator(Refrigerator) error{
	fmt.Println("用 ship 做 OpenTheDoorOfRefrigerator")
	return nil
}

func (*shipLegend) PutElephantIntoRefrigerator(Elephant,Refrigerator ) error{
	fmt.Println("用 ship 做 PutElephantIntoRefrigerator")
	return nil
}

func (*shipLegend) CloseTheDoorOfRefrigerator(Refrigerator) error{
	fmt.Println("用 ship 做 CloseTheDoorOfRefrigerator")
	return nil
}