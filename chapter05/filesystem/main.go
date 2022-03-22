package main

import "fmt"

func main() {
	var data string
	{
		var equipment IOInterface = &Soft{} //这个结构体给到这个接口
		data = equipment.Read()             //data 读取接口即可，不管你是啥，我只调用这个接口的功能
		fmt.Println(data)
	}
	{
		var equipment IOInterface = &Mag{} //这个结构体给到这个接口
		data = equipment.Read()            //data 读取接口即可，不管你是啥，我只调用这个接口的功能
		fmt.Println(data)
	} //这里你会发现只要给不同的结构体，也就是所谓的变量，就会获取不同的结果！就可以从不同的设备上读取数据
	{
		var equipment IOInterface = &Paper{} //这个结构体给到这个接口
		data = equipment.Read()              //data 读取接口即可，不管你是啥，我只调用这个接口的功能
		fmt.Println(data)
	}
	{
		var equipment IOInterface = &Sata{} //这个结构体给到这个接口
		data = equipment.Read()              //data 读取接口即可，不管你是啥，我只调用这个接口的功能
		fmt.Println(data)
	}
}

type IOInterface interface { //定义一 个接口！
	Read() (data string) //你只要符合这个接口，上面就能读取到！
}

//定义一个结构体
type Soft struct {
}

func (Soft) Read() string { //成员函数，这里没有给Soft命名，是因为结构体没有其他成员变量
	return "从1.4寸软盘进行读取----啦啦啦啦啦" //你只要符合我接口的定义，就能插入！！
}

type Mag struct {
}

func (Mag) Read() string {
	return "从磁带进行读取----滋滋滋滋滋滋滋滋滋滋"

}

type Paper struct {
}

func (Paper) Read() string {
	return "从纸带进行读取----000000001111111"

}

type Sata struct {
}

func (Sata) Read() string{
	return "从Stag盘读取-----呲溜呲溜呲溜"
}
