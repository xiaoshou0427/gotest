package main

import "fmt"

func main() {
	m3 := map[string]int{"王强": 60, "李静": 83, "苗文": 91}
	fmt.Println(m3)
	fmt.Println(m3["abc"])
	abcScore, ok := m3["abc"] //这里定义两个变量， 第一个是值，第二个是option，返回true/false
	fmt.Println(abcScore, ">>>>>>", ok)
	m3["abc"] = 99
	abcScore, ok = m3["abc"] //这里是赋值操作！
	fmt.Println(abcScore, ">>>>>>", ok)

}
