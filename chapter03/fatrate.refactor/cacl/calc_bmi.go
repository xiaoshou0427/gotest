package cacl

import (
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
)
//目录是cacl

func CalcBMI(tall float64, weight float64) (bmi float64, err error) {
	if tall <= 0  {
		return 0, fmt.Errorf("身高不能为0或者负数")
	} else if weight <= 10 || weight >=200 {
		return 0, fmt.Errorf("输入的体重\"%f\"超出有效范围", weight)
	}
	bmi,_ = gobmi.BMI(weight,tall)
	return
}
