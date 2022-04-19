package main

import (
	"fmt"
	"sort"
)

// 按钮，有很多个楼层
type Button struct {
	Floor int
}

//电梯有按钮和现在所在的位置
type Elevator struct {
	buttons  Buttons
	position int
}

type Buttons []*Button //类型重定义

//自动生成-成员函数，这个时候传进去的buttons 就可以排序了
func (b Buttons) Len() int {
	return len(b) //返回长度
}

//如果i 小于 j ，排序位置不发生变化（这只是个方法！！）
//下面是返回true/false，ihej的floor 进行对比，i的floor小于j的floor 就true
func (b Buttons) Less(i, j int) bool {
	return b[i].Floor < b[j].Floor //对比floor，b[i]是index i 的：{Floor: x}，b[j]也类似
}

//交换位置
func (b Buttons) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
	//tmpObj := b[i] 这种写法就挫了，用上面这个写法！！！
	//b[i] = b[j]
	//b[j] = tmpObj
}

func main() {
	ev := &Elevator{
		position: 2, //现在的具体位置
		buttons: Buttons{
			{Floor: 3}, //按3楼
			{Floor: 1}, //按1楼
			{Floor: 5},
			{Floor: 2},
			{Floor: 4},
		},
	}
	//要对楼层进行排序
	//sort.Sort(ev.buttons) //自己debug吧
	sort.Sort(sort.Reverse(ev.buttons))

	fmt.Println(ev.buttons) //指针
	fmt.Printf("%+v\n",ev.buttons) //结果出不来
	for _,item:=range ev.buttons{
		fmt.Println(item.Floor) //可以看到排好序的楼层！
	}

}
