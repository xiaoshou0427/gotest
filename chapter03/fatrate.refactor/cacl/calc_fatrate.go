package cacl

import gobmi "github.com/armstrongli/go-bmi"

func FatRateCalc(bmi float64, age int, sex string) (fatRate float64) {
	fatRate = gobmi.FatRateCalc(bmi,age,sex)
	return
}
