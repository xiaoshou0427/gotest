package main

import (
	"fmt"
	"math/rand"
)

func main() {

	vg :=&voteGame{
		students: []*student{ //嵌套的指针，需要实例化嵌套的指针
			&student{name: fmt.Sprintf("%s","A")}, //Sprintf 这里终端不显示，是是把这个值赋给name
			&student{name: fmt.Sprintf("%s","B")},
			&student{name: fmt.Sprintf("%s","C")},
			&student{name: fmt.Sprintf("%s","D")},
			&student{name: fmt.Sprintf("%s","E")},
			&student{name: fmt.Sprintf("%s","F")},
			&student{name: fmt.Sprintf("%s","G")},
			&student{name: fmt.Sprintf("%s","H")},
			&student{name: fmt.Sprintf("%s","I")},
			&student{name: fmt.Sprintf("%s","G")},
			&student{name: fmt.Sprintf("%s","K")},
			&student{name: fmt.Sprintf("%s","L")},
			&student{name: fmt.Sprintf("%s","M")},
			&student{name: fmt.Sprintf("%s","N")},
			&student{name: fmt.Sprintf("%s","O")},
			&student{name: fmt.Sprintf("%s","P")},
			&student{name: fmt.Sprintf("%s","Q")},
		},
	}
	leader := vg.goRun()
	fmt.Println(leader.name)
	leader.Distribute()
}

type voteGame struct {
	students []*student //一个列表是指针结构体student类型，多个学生的得票,嵌套指针，在使用的时候要实例化
}
type Leader student //类型重定义，更容易理解

func (l *Leader) Distribute(){
	fmt.Println("发作业了！班长深沉的说到---",l.name)
}

func (g *voteGame) goRun() *Leader {  //成员函数，goRun，实参为指针类型的Leader,即指针类型的student
	for _, item := range g.students { //循环所有学生
		randInt := rand.Int() //定义一个变量，随机整数
		if randInt%2 ==  0 { //随机数对2取模=0，即随机数为偶数，则投票给下标为randInt%len(g.students)的学生
			//这里len(g.students),比如上面A-E是 长度是5，随机数对5取模，只会是0-4！数学问题，防止 out of range！
			item.voteA(g.students[randInt%len(g.students)]) //随机投票给agree
		} else {
			item.voteA(g.students[randInt%len(g.students)]) //奇数就投disagree
			//这样就随机分配给不同的人了！随机偶数就投agree，投给谁也是随机的！
		}

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
		return (*Leader)(g.students[maxScoreIndex]) //返回一个最大值
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
