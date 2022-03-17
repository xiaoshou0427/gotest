package main

import "fmt"

func main() {

	vg :=&voteGame{
		students: []*student{ //嵌套的指针，需要实例化嵌套的指针
			&student{name: fmt.Sprintf("%d",1)}, //Sprintf 这里终端不显示，是是把这个值赋给name
			&student{name: fmt.Sprintf("%d",2)},
			&student{name: fmt.Sprintf("%d",3)},
			&student{name: fmt.Sprintf("%d",4)},
			&student{name: fmt.Sprintf("%d",5)},
		},
	}
	leader := vg.goRun()
	fmt.Println(leader)
}

type voteGame struct {
	students []*student //一个列表是指针结构体student类型，多个学生的得票,嵌套指针，在使用的时候要实例化
}
type Leader = student //类型重命名，更容易理解

func (g *voteGame) goRun() *Leader {  //成员函数，goRun，实参为指针类型的Leader,即指针类型的student
	for _, item := range g.students { //循环所有学生
		item.voteA(g.students[0]) //总是给第一个投票，//todo 可以用随机数代替
	}
	//找到票数最高的那个人！
	maxScore := -1
	maxScoreIndex := -1
	for i, item := range g.students {
		if maxScore < item.agree { //也永远是整数
			maxScore = item.agree //如果maxScore小于最高分，就将maxScore等于最高分
			maxScoreIndex = i     //此时下标是最大值的下标
		}
	}
	if maxScoreIndex >= 0 { //下标一定是从0开始的，防止意外，如果没有学生，那么index就是默认值-1
		return g.students[maxScoreIndex] //返回一个最大值
	}
	return nil //剩下就给个空，指针类型可以用nil 作为返回结果
}

type student struct {
	name 	 string
	agree    int //一个人有赞成和反对票
	disagree int
}

//赞成票
func (std *student) voteA(target *student) {
	target.agree++
}

//反对票
func (std *student) voteD(target *student) {
	target.disagree++
}
