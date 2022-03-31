package main

import (
	gobmi "github.com/armstrongli/go-bmi"
	"log"
)

type Person struct {
	name   string
	sex    string
	tall   float64
	weight float64
	age    int

	bmi     float64
	fatRate float64
}

//定义一个成员函数，返回error即可，因为调用另外一个包gobmi
func (p *Person) calcBmi() error {
	bmi, err := gobmi.BMI(p.weight, p.tall)
	if err != nil {
		log.Printf("Error when calculate BMI for Person[%s]: %v:", p.name,err)
		//这里是在离错误最近的地方（贴近报错）把log打印出来
		return err
	}
	p.bmi = bmi
	return nil
}

func (p *Person) FatRate() {
	p.fatRate = gobmi.FatRateCalc(p.bmi,p.age,p.sex)
	//此处省略return
}
