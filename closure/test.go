package main

import "fmt"

//var base float64 = 1000.0

func getImprovementFunc() func(percentage float64) {
	base := 1000.0
	fmt.Printf("base: %f\n", base)
	return func(percentage float64) {
		base = base * (1 + percentage)
		fmt.Println("New: ", base)
	}
}

func closureMain() {
	imp := getImprovementFunc()
	imp(0.1) //每年百分之10的赠长
	imp(0.2) //每年百分之10的赠长
	imp()    //每年百分之10的赠长
	imp()    //每年百分之10的赠长

}
