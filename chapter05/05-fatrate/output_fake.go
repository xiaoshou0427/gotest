package main

//实现output的接口

type fakeOutput struct {
	p Person
	s string
}

//验证预期，成员函数的入参 赋值给结构体？ 这不用指针也能写进去？
func (o *fakeOutput) Output(p Person, s string)  {
	o.p = p
	o.s = s
}
