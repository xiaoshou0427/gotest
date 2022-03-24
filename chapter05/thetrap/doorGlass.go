package main

import "fmt"

//匿名函数，类型为Door，加一个标志
var _ Door = &GlassDoor{}
//语义：我们要把一个具体的GlassDoor这个对象的指针，赋值给一个Door类型的变量，赋值到的地方是nowhere
// 不把它放在某个实际的变量上，这样就表示我要强制实现这个接口的所有方法！

type GlassDoor struct {
}

func (d *GlassDoor) Unlock() {
	fmt.Println("GlassDoor Unlock")
}

func (d *GlassDoor) Lock() {
	fmt.Println("GlassDoor Lock")
}

func (*GlassDoor) Open() {
	fmt.Println("GlassDoor Open")
}
func (*GlassDoor) Close() {
	fmt.Println("GlassDoor Close")
}
