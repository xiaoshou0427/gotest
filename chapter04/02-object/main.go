package main

import (
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
)

func main() {

	persons := []Person{ //这是一个slice 切片, 这里很像二维数组
		{ //录入一个人的信息
			"小强",
			"男",
			1.7,
			71,
			34,
		},
	}
	fmt.Println(persons[0])
	for _, item := range persons { //提取值，不需要index
		bmi, err := gobmi.BMI(item.weight, item.tall)
		fmt.Println("BMI:", bmi, "err:", err)
	}
	a := new(Person)
	fmt.Println(a)
	var b = &Person{}
	fmt.Println(b)
}

type Person struct {
	name   string
	sex    string
	tall   float64
	weight float64
	age    int
}
