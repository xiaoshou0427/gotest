package main

import (
	"fmt"
	"log"
)

type fatRateService struct {
	s *fatRateSuggestion //结构体嵌套哦，整个人的信息都在这里,建议也在这里
}

func (frsvc *fatRateService) GiveSuggestionToPerson(person *Person) string {
	if err := person.calcBmi(); err != nil {
		log.Printf("无法给%s计算体脂:%v", person.name, err)
		return "错误"
	}
	//这个时候拿到某个人的信息，在上面err的变量已经调用了person的成员函数计算bmi
	//下面计算出体脂率，并不需要给到变量，回写到person里面
	person.FatRate()
	return frsvc.s.GetSuggestion(person)
	//这个return 就有意思了！是个闭包！？
	//这就是个闭包，别人在调用这个函数的时候 变量 := 函数 ，拿到的是return的结果
	//那么变量其实等于 frsvc.s.GetSuggestion(person)
	//当你给这个变量一个形参的时候：
	// 变量(入参x) = frsvc.s.GetSuggestion(入参x)） 获取体型结果
	//下次 变量(入参y) 就得到新的体型结果！
}

//给一堆人做建议，入参是个（slice Person）类型的指针那返回值就是一个map！key是Person类型的指针，value是个字符串
func (frsvc *fatRateService) GiveSuggestionToPersons(persons ...*Person) map[*Person]string {
	out := map[*Person]string{} //往里面装
	for _, item := range persons {
		out[item] = frsvc.GiveSuggestionToPerson(item) //这里不是Persons!! 这个item 是个结构体哦！
		fmt.Println("[item]:",item)
		fmt.Println("out[item]:",out[item])
	}
	return out
}
