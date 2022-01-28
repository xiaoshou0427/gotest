package cacl

func FatRateCalc(bmi float64, age int, sex string) (fatRate float64) {
	sexWeight := 0  // 这个变量不用在外面定义，也不用形参，因为这个不在外面有任何意义，所以可以在这里定义
	if sex == "男" { //根据传入的参数，来确认 男女用什么值来计算出体脂率
		sexWeight = 1
	} else {
		sexWeight = 0
	}
	//这里没有考虑非法输入，视频是在外面做的
	fatRate = (1.2*bmi + getAgeWeight(age)*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	return
}

//计算当不同年龄，更改计算系数
func getAgeWeight(age int) (ageWeight float64) {
	ageWeight = 0.23
	if age >= 30 && age <= 40 {
		ageWeight = 0.22
	}
	return
}
