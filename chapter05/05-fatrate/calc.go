package main

import (
	gobmi "github.com/armstrongli/go-bmi"
	"log"
)

type Calc struct {
}

// 这里注意： 如果结构体没有成员，下面不需要(c Calc),而直接用(Calc)
// 表示我们是定义在这个对象上的方法！！
func (Calc) BMI(person *Person) error { //这个函数是BMI ,传进来的是什么？ 是人纳！是person，用来计算BMI的
	//这里整个计算结果会放到person 里面，而不是说我们计算完毕后return 出去，我们return 出去可以是error
	//读和写都是在person这个对象上的
	bmi, err := gobmi.BMI(person.weight, person.tall) //偷懒去调用本地写的gobmi
	if err != nil {
		log.Println("error when calculate bmi:", err)
		return err //如果err不为空就返回err
		//如果有出错的地方，尽量贴近错误的地方把err打出来，给个log出来
	}
	person.bmi = bmi //把bmi写回到 person.bmi里,为什么要写进去呢，因为下面的fatrate 要用这个值
	return nil       //如果没有出错的话，就返回nil
}

func (Calc) FatRate(person *Person) error {
	person.fatRate = gobmi.FatRateCalc(person.bmi, person.age, person.sex) //偷懒调用函数
	//上面gobmi.FatRateCalc 只有一个返回值，直接用person.fatRate把它装起来，这里赋值即可
	return nil //与BMI 相同的逻辑
} //完成计算器部分
//这里需要修改对象的成员内容，就需要用到指针：https://zhuanlan.zhihu.com/p/46673861