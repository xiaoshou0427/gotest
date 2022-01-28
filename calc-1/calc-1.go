package main

import "fmt"

func main() {
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
	var bmi float64 = weight / (tall * tall)
	var sexWeight int
	if sex == "男" { //这里只会是男和女，就一个if/else即可
		sexWeight = 1
	} else {
		sexWeight = 0
	}
	var fateRate float64 = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	fmt.Printf("您好，%v，体脂率为: %v\n", name, fateRate)
	if sex == "男" {
		if age >= 18 && age <= 39 {
			if fateRate <= 0.1 { // 不同区间得到不同的结果
				fmt.Println("您的体重有些偏瘦")
			} else if fateRate > 0.1 && fateRate <= 0.16 {
				fmt.Println("您的体重很标准")
			} else if fateRate > 0.16 && fateRate <= 0.21 {
				fmt.Println("您的体重有点偏重")
			} else if fateRate > 0.21 && fateRate <= 0.26 {
				fmt.Println("您的体重属于肥胖")
			} else {
				fmt.Println("您的体重严重肥胖了")
			}
		} else if age >= 40 && age <= 59 {
			//todo
		} else if age >= 60 {
			//todo
		} else {
			fmt.Println("不判断未成年人的体脂率的区间范围，变化太大，无法评判")
		}
	} else {
		//todo
	}
}
