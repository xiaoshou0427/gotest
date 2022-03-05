package main

import (
	"fmt"
)

func main() {
	var left, right int = 1, 2
	//var op string = "+" //这个时候 加号没什么用

	c := Calculator{
		left:  left,
		right: right,
		//op: op,
	}
	fmt.Printf("&c @ main = %p\n", &c)
	fmt.Println(c.Add()) //调了对象的成员函数
	fmt.Println("c.result=", c.result)

	newC := NewCalculter{}
	newC.left = 100
	newC.right = 200
	fmt.Println(newC.Add())

	mc := MyCommand{
		commandOptions: map[string]string{},
	} //定义变量为结构体（对象）
	mc.commandOptions["aaa"] = "AAA" //变量取结构体的成员变量，赋值成员变量
	fmt.Println(mc.ToCmdStr()) //打印成员变量的输出
}

type MyCommand struct {
	commandOptions map[string]string
	//定义成员变量和成员变量的类型
}

func (my MyCommand) ToCmdStr() string { //这里成员函数
	out := ""
	for k, v := range my.commandOptions { //成员函数调用了成员变量
		out = out + fmt.Sprintf("--%s=%s", k, v) //sprintf 返回字符串
	}
	return out
}
