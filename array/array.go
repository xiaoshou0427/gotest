package main

import "fmt"

func main() {
	Info := [][3]string{
		[3]string{"小强", "男", "在职"}, //定义二维数组的元素，可以不写[3]string
		{"小李", "男", "在职"},
		{"小苏1", "女", "在职"},
		{"小苏2", "女", "在职"},
		{"小苏3", "女", "在职"},
		{"小苏4", "女", "在职"},
	}
	Info = append(Info, [3]string{"小孟", "男", "在职"})
	fmt.Println("降维输出")
	for d1, d1val := range Info {
		//fmt.Println(d1val)
		for d2, d2val := range d1val {
			fmt.Println(d1, d1val, d2, "d2val2:", d2val)
		}
	}
}
