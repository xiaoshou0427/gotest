package cacl

//目录是cacl

func CalcBMI(tall float64, weight float64) (bmi float64) {
	//需要考虑非法输入
	if tall <= 0 {
		panic("身高不能为0或者负数")
	} else if weight <= 0 { //验证体重的合法性
		panic("体重不能为0或者负数")
	} else {
		return weight / (tall * tall) //预期产出 给到bmi变量
	}
	//bmi = weight / (tall * tall) 可以分两行写
	//return  //这种写法也没有问题，就是多了一行
	//这里直接用 return weight / (tall * tall) ---代码更加简洁
}
