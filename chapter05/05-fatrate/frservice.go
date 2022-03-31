package main

import (
	"log"
)

type fatRateService struct {
	input InputService       //对象里面成员变量的类型是个接口
	s     *fatRateSuggestion //结构体嵌套哦，整个人的信息都在这里,建议也在这里
	output OutputService //
}

func (frsvc *fatRateService) GiveSuggestionToPerson(person *Person)  {
	if err := person.calcBmi(); err != nil {
		log.Printf("无法给%s计算体脂:%v", person.name, err)
		return //没有预期产出，直接给个return
	}
	//这个时候拿到某个人的信息，在上面err的变量已经调用了person的成员函数计算bmi
	//下面计算出体脂率，并不需要给到变量，回写到person里面
	person.FatRate()
	frsvc.output.Output(*person,frsvc.s.GetSuggestion(person)) //输出到output这个接口了
	//调用的时候需要实例化！
	//return frsvc.s.GetSuggestion(person) //这里就不用return 这个结果出来了
}


