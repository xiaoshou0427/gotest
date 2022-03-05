package main

import (
	"fmt"
)

func main() {
	a, b := 1, 2
	add(&a, &b)
	fmt.Println("a:", a, "&a:", &a)
	c := &a // c的类型是 *int， c 指向a的盒子，*c 就可以拿到a里面的东西，3
	fmt.Println("c：", c, "*c:", *c)
	d := &c // d的类型是**int，d指向c的盒子，d本身是指针，它存放的东西也是指针。
	fmt.Println("d：", d, "*d ", *d, "**d:", **d)

	m := map[string]string{}
	mp1 := &m // mp1的类型就是 *map[string]string
	fmt.Println(mp1)
	put(m)
	fmt.Println("打印mp1：", *mp1)

	f1 := add  // f1 = func(int,int) f1 是一个函数
	f1(&a, &b) //必须是指针类型的a 和 b，不能是其他数字
	fmt.Println("f1, add =", a)
	f1p1 := &f1     //f1p1 = *func(int,int)
	(*f1p1)(&a, &b) //这里有个golang优先级的概念，先计算，再给个结果，再给个星！所以用括号扩起来，来先拿到函数的指针
	fmt.Println("f1p1, add = ", a)


	var nothing *int
	//*nothing = 3
	fmt.Println(nothing)

	var nothingMap map[string]string = map[string]string{}
	nothingMap["aaa"] = "aaa"
	fmt.Println(nothingMap)

	{
		var nothingSlice []int
		nothingSlice= append(nothingSlice,100)
		fmt.Println("nothingSlice:",nothingSlice)
	}
}

func put(m map[string]string) {
	m["a"] = "AAA"
}
func add(a, b *int) {
	*a = *a + *b
}
