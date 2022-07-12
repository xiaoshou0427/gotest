package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//运动员
type Runner struct {
	Name string
}

//成员函数，这个对象的动作是什么！
//这里入参是变量：wg，类型是waitgroup对象（结构体），这是个指针，指针变量wg 指向这个结构体的内存地址，这样快！？
//startPointWg 给每个人，一个起跑令(这也是一个waitgroup类型的对象)
func (r Runner) Run(startPointWg, wg *sync.WaitGroup) {
	defer wg.Done()     //这个跟下面的wg不是一个意思哦！这个done定义在了这里哦！
	startPointWg.Wait() //我在这等着，我在各个routine里面等着，等着统一的信号！然后大家才一起开始跑
	start := time.Now()
	fmt.Println(r.Name, "开始跑@", start)
	rand.Seed(time.Now().UnixNano())                          //给一个随机时间？
	time.Sleep(time.Duration(rand.Uint64()%10) * time.Second) //给一个随机数,随机数对10取余数，不会大于10哦！只会是0-9
	//time.Duration 做强制转换，换成时间，用rand.int()有风险，可能是负数!
	finish := time.Now()
	fmt.Println(r.Name, "跑到终点，用时：", finish.Sub(start)) //把耗时也打出来！

}

func main() {
	runnerCount := 10     //10个人
	runners := []Runner{} //实例化，空的列表，往这个变量里面写数据

	wg := sync.WaitGroup{}
	wg.Add(runnerCount) //计数器10个

	startPointWg := sync.WaitGroup{}
	startPointWg.Add(1) //起跑的计数（起跑铃）

	//把人名写到了runners里面，每个index对应一个Runner对象
	for i := 0; i < runnerCount; i++ {
		runners = append(runners, Runner{ //Runner 生成一个实例化对象
			Name: fmt.Sprintf("%d", i), //每个人都有自己的编号了！这个sprintf 是不输出到标准输出，直接返回值，看源码
		})
	}
	//每个人开始跑了,循环这些人
	for _, runnerItem := range runners {
		go runnerItem.Run(&startPointWg, &wg) //开始跑，随机数后，到达终点，需要waitgroup参与到其中！
		//这里是开启了goroutine，可以理解为新的进程去做这个函数
	}
	//上面这一部分是起的goroutine在跑，每个go routine都在工作，此时都卡在了startPointWg.Wait()这里等待
	//此时main goroutine 继续执行（这个时间是并行的），执行到了下面的这些print
	fmt.Println("各就各位")
	time.Sleep(2 * time.Second)
	fmt.Println("预备：跑！")
	//main gouroutine 执行到这里，发现有个Done--那么其他goroutine的startPointWg.Wait() 看到变为0了，继续向下执行剩余的内容！
	startPointWg.Done()

	wg.Wait() //等待所有的done
	fmt.Println("赛跑结束")

	//不管怎样，所有人都是在10秒内跑完，用了goroutine！

}
