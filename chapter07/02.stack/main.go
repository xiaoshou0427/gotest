package main

import "fmt"

type Stack struct {
	data []interface{}
}

//在顶部push一个进去，也就是永远在最前面！
//还能这么定义data为slice类型的接口？ --》 神奇了 []interface{}{data}
func (s *Stack) Push(data interface{}) {
	s.data = append([]interface{}{data}, s.data...) //用一个新的slice，把已有的append到新的slice
}

func (s *Stack) Pop() (interface{}, bool) {
	if len(s.data) > 0 {
		o := s.data[0]
		s.data = s.data[1:] //切一下，把最初的数据从data中切出
		return o, true
	}
	return nil, false //不大于0 返回空
}

func main() {
	s := &Stack{} //为啥加指针？ 往里面写东西？改变数据？没有这个& 一样可以执行！
	s.Push(111)
	s.Push(222)
	s.Push(333)
	s.Push(nil)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

}
