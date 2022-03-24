package main

import "fmt"

func main()  {
	security:=Assets{assets: []Asset{
		&GlassDoor{},
		&WoodDoor{},
	}}
	fmt.Println("开始上班")
	security.DoStartWork() //开始查岗
	fmt.Println("8小时到，准备下班")
	security.DoStopWork()
	fmt.Println("Done")
}

