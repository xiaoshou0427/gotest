package main

type NewCalculter struct {
	Calculator //与之前结构体的区别在于，这里没有写属性名，而是直接把另外一个对象写了进来
	// *Calculator 如果是嵌套的指针，则在实例化NewCalculator的时候，必须实例化嵌套的实体，否则会报空指针
}