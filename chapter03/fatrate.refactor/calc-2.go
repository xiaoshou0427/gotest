package main

import (
	"fmt"
	"learn.go/fatrate.refactor/cacl"
)

func main() {
	for {
		mainFatRateBody()
		// 其他
		if condition := whetherContinue(); !condition {
			break
		}
	}
}

func mainFatRateBody() {
	//录入个人信息
	name, sex, tall, weight, age := getInfo()
	//计算体脂
	fatRate := fatRateCalc(tall, weight, age, sex, name)
	//fatRate := fatRateCalc(bmi, age, sexWeight, name)
	//获取体脂信息
	getFatRateResult(sex, age, fatRate)
}

func getFatRateResult(sex string, age int, fatRate float64) {
	if sex == "男" {
		getHealthinessSuggestions(age, fatRate, getHealthinessSuggestionForMale)
		//这里调用函数,age 和fatRate 是形参，getHealthinessSuggestionForMale 是个函数赋值给getSuggestion
		//调用的是getHealthinessSuggestionForMale(age,fatRate) 函数
		// 这里面age 和fatRate 是上面一个函数传入进来的，带入这里，再给到getSuggestion(age, fatRate)，也就是给到getHealthinessSuggestionForMale(age,fatRate)
	} else {
		getHealthinessSuggestions(age, fatRate, getHealthinessSuggestionForFemale)
		//这里相当于我调用别人的函数（别人写好的工具），将我的参数传入到工具中获得结果！所以前面两个参数是为了使用后面这个工具而设置的！
		//对于使用者来说，传入参数，给到后面的工具即可！
		//fixme
	}
}

func getHealthinessSuggestions(age int, fatRate float64, getSuggestion func(age int, fatRate float64)) {
	getSuggestion(age, fatRate) //你的函数，回调这个工具！
} // 这里定义了一个函数，age 和fatRate是需要上一层传入的！定义一个参数叫getSuggestion ，这个变量是一个函数 可以理解为 getSuggestion := func(xx,xx)
//函数形参输入后，返回形参的结果。 这里这个工具也是回调别的函数

func getHealthinessSuggestionForMale(age int, fatRate float64) {
	if age >= 18 && age <= 39 {
		if fatRate <= 0.1 { // 不同区间得到不同的结果
			fmt.Println("您的体重有些偏瘦")
		} else if fatRate > 0.1 && fatRate <= 0.16 {
			fmt.Println("您的体重很标准")
		} else if fatRate > 0.16 && fatRate <= 0.21 {
			fmt.Println("您的体重有点偏重")
		} else if fatRate > 0.21 && fatRate <= 0.26 {
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
}

func getHealthinessSuggestionForFemale(age int, fatRate float64) {
	if age >= 18 && age <= 39 {
		if fatRate <= 0.1 { // 不同区间得到不同的结果
			fmt.Println("您的体重有些偏瘦")
		} else if fatRate > 0.1 && fatRate <= 0.16 {
			fmt.Println("您的体重很标准")
		} else if fatRate > 0.16 && fatRate <= 0.21 {
			fmt.Println("您的体重有点偏重")
		} else if fatRate > 0.21 && fatRate <= 0.26 {
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
}

func fatRateCalc(tall float64, weight float64, age int, sex string, name string) (fatRate float64) {
	bmi := cacl.CalcBMI(tall, weight)         //对于我来说bmi 只需要计算出结果，来进行体脂率计算即可，所以不用传入参数，这里调用cacl包
	fatRate = cacl.FatRateCalc(bmi, age, sex) //在这里调用包cacl包
	fmt.Printf("您好，%v，体脂率为: %v\n", name, fatRate)
	return fatRate
}

func getInfo() (name string, sex string, tall float64, weight float64, age int) {
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Print("请输入性别（男/女）：")
	fmt.Scanln(&sex)
	fmt.Print("请输入身高（单位：米，例如：1.70）：")
	fmt.Scanln(&tall)
	fmt.Print("请输入体重（单位：千克，例如：71.5)：")
	fmt.Scanln(&weight)
	fmt.Print("请输入年龄：")
	fmt.Scanln(&age)
	return name, sex, tall, weight, age
}

func whetherContinue() bool {
	var whetherContinue string
	fmt.Print("是否继续录入信息（y/n):")
	fmt.Scanln(&whetherContinue)
	if whetherContinue == "y" || whetherContinue == "Y" {
		return true //这里一定是这么写的，视频是只判断了不等于y 就返回false
	} else {
		return false
	}
}
