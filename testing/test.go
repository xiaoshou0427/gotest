package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	name := "小强"
	{
		fmt.Println(reflect.TypeOf(name))
		val, err := strconv.Atoi(name)
		fmt.Println(val,err)
		fmt.Println(reflect.TypeOf(val))
	}
}
