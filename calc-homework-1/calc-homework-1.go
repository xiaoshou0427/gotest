package main

import "fmt"

func main() {
	var totalFatRate float64
	names := [3]string{}
	weights := [3]float64{}
	talls := [3]float64{}
	ages := [3]int{}
	bmis := [3]float64{}
	fatRates := [3]float64{}
	var sex string

	for i := 0; i < 3; i++ {
		name, sex, tall, weight, age := getinfoFromInput()
		bmi := calcBMI(tall, weight)
		bmis[i] = bmi
		var sexWeight int
		if sex == "男" { //这里只会是男和女，就一个if/else即可
			sexWeight = 1
		} else {
			sexWeight = 0
		}
		fatRates[i] = (1.2*bmis[i] + 0.23*float64(ages[i]) - 5.4 - 10.8*float64(sexWeight)) / 100
		fmt.Printf("您好，%v，体脂率为: %v\n", names[i], fatRates[i])
		if sex == "男" {
			if ages[i] >= 18 && ages[i] <= 39 {
				if fatRates[i] <= 0.1 { // 不同区间得到不同的结果
					fmt.Println("您的体重有些偏瘦")
				} else if fatRates[i] > 0.1 && fatRates[i] <= 0.16 {
					fmt.Println("您的体重很标准")
				} else if fatRates[i] > 0.16 && fatRates[i] <= 0.21 {
					fmt.Println("您的体重有点偏重")
				} else if fatRates[i] > 0.21 && fatRates[i] <= 0.26 {
					fmt.Println("您的体重属于肥胖")
				} else {
					fmt.Println("您的体重严重肥胖了")
				}
			} else if ages[i] >= 40 && ages[i] <= 59 {
				//todo
			} else if ages[i] >= 60 {
				//todo
			} else {
				fmt.Println("不判断未成年人的体脂率的区间范围，变化太大，无法评判")
			}
		} else {
			//todo
		}
		/*		var whetherContinue string
				fmt.Print("是否继续录入信息（y/n):")
				fmt.Scanln(&whetherContinue)
				if whetherContinue == "y" || whetherContinue == "Y" {
					continue
				} else {
					break
				}*/
	}
	for i := 0; i < 3; i++ {
		totalFatRate += fatRates[i]
		fmt.Println("姓名：", names[i], "体脂率：", fatRates[i])
	}
	fmt.Println("3人的平均体脂率为：", totalFatRate/3)
}

func getinfoFromInput() (name string, sex string, tall float64, weight float64, age int) {
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Print("请输入性别：（男/女）")
	fmt.Scanln(&sex)
	fmt.Print("请输入身高（单位：米，例如：1.70）：")
	fmt.Scanln(&tall)
	fmt.Print("请输入体重（单位：千克，例如：71.5)：")
	fmt.Scanln(&weight)
	fmt.Print("请输入年龄：")
	fmt.Scanln(&age)
	return
}

func calcBMI(tall float64, weight float64) float64 {
	return weight / (tall * tall)

}
