package main

import (
	"fmt"
	"testing"
	"time"
)

func TestDefChannel(t *testing.T) {
	var sampleMap map[string]int = map[string]int{} //实例化，空map
	fmt.Println("sampleMap:", sampleMap)
	var intCh chan int //= make(chan int) // 是不是也可以这么写？应该是可以的
	fmt.Println("intCh:", intCh)
	intCh = make(chan int, 1) //把之前定义的变量 做一个初始化，make操作！
	fmt.Println("intCh:", intCh)

	fmt.Println("装入数据")
	intCh <- 3 //装数据
	fmt.Println("取出数据")
	out := <-intCh
	fmt.Println("数据是:", out)
}

func TestChanPutGet(t *testing.T) {
	intCh := make(chan int) //创建一个不带Size的channel（不带buffer）
	workerCount := 10       //10个
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			intCh <- i //起10个goroutine往里面放每个workerCount
		}(i)
	}

	for o := 0; o < workerCount; o++ {
		go func(o int) {
			out := <-intCh //起10个goroutine 取数据
			fmt.Printf("o变量：%d 拿到的值：%d\n", o, out)
		}(o)
	}
	time.Sleep(1 * time.Second) //这里用睡，就没用waitgroup了
}

// 这是一个让out部分等待一段时间再取数据，来观察i的行为
// 结果： 如果没有out，那么in部分会等待，直到有out开始
func TestChanPutGet2_Owait(t *testing.T) {
	intCh := make(chan int) //创建一个不带Size的channel（不带buffer）
	workerCount := 10       //10个
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			fmt.Println(i, "开始工作", time.Now())
			intCh <- i //起10个goroutine往里面放每个workerCount
			fmt.Println(i, "结束工作", time.Now())
		}(i)
	}
	fmt.Println(time.Now())
	time.Sleep(2 * time.Second) //这里用睡
	fmt.Println(time.Now())

	for o := 0; o < workerCount; o++ {
		go func(o int) {
			out := <-intCh //起10个goroutine 取数据
			fmt.Printf("时间：%s o变量：%d 拿到的值：%d\n", time.Now(), o, out)
		}(o)
	}
	time.Sleep(1 * time.Second) //这里用睡，就没用waitgroup了
}

func TestChanPutGet2_Owait_withBuffer(t *testing.T) {
	intCh := make(chan int, 10) //创建一个不带Size的channel（不带buffer）
	workerCount := 10           //10个
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			fmt.Println(i, "开始工作", time.Now())
			intCh <- i //起10个goroutine往里面放每个workerCount
			fmt.Println(i, "结束工作", time.Now())
		}(i)
	}
	fmt.Println(time.Now())
	time.Sleep(2 * time.Second) //这里用睡
	fmt.Println(time.Now())

	for o := 0; o < workerCount; o++ {
		go func(o int) {
			out := <-intCh //起10个goroutine 取数据
			fmt.Printf("时间：%s o变量：%d 拿到的值：%d\n", time.Now(), o, out)
		}(o)
	}
	time.Sleep(1 * time.Second) //这里用睡，就没用waitgroup了
}

func TestChanPutGet2_Owait_withSmallBuffer(t *testing.T) {
	intCh := make(chan int, 2) //创建一个不带Size的channel（不带buffer）
	workerCount := 10          //10个
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			fmt.Println(i, "开始工作", time.Now())
			intCh <- i //起10个goroutine往里面放每个workerCount
			fmt.Println(i, "结束工作", time.Now())
		}(i)
	}
	fmt.Println(time.Now())
	time.Sleep(2 * time.Second) //这里用睡
	fmt.Println(time.Now())

	for o := 0; o < workerCount; o++ {
		go func(o int) {
			out := <-intCh //起10个goroutine 取数据
			fmt.Printf("时间：%s o变量：%d 拿到的值：%d\n", time.Now(), o, out)
		}(o)
	}
	time.Sleep(1 * time.Second) //这里用睡，就没用waitgroup了
}

func TestChanPutGet2_OFirstwithBuffer(t *testing.T) {
	intCh := make(chan int, 10) //创建一个不带Size的channel（不带buffer）
	workerCount := 10           //10个
	for o := 0; o < workerCount; o++ {
		go func(o int) {
			out := <-intCh //起10个goroutine 取数据
			fmt.Printf("时间：%s o变量：%d 拿到的值：%d\n", time.Now(), o, out)
		}(o)
	}
	fmt.Println(time.Now())
	time.Sleep(2 * time.Second) //这里用睡
	fmt.Println(time.Now())

	for i := 0; i < workerCount; i++ {
		go func(i int) {
			fmt.Println(i, "开始工作", time.Now())
			intCh <- i //起10个goroutine往里面放每个workerCount
			fmt.Println(i, "结束工作", time.Now())
		}(i)
	}
	time.Sleep(1 * time.Second) //这里用睡，就没用waitgroup了
}

func TestRangeChannel(t *testing.T) {
	intCh := make(chan int, 10)
	intCh <- 1
	intCh <- 2
	intCh <- 3
	intCh <- 4
	intCh <- 5

	for o := range intCh {
		fmt.Println(o)
	}
}

func TestRangeClosedChannel(t *testing.T) {
	intCh := make(chan int, 10)
	intCh <- 1
	intCh <- 2
	intCh <- 3
	intCh <- 4
	intCh <- 5

	close(intCh) //关闭水管

	{
		o1, ok := <-intCh
		fmt.Println("直接取数：", o1, ok)
	}

	for o := range intCh {
		fmt.Println("range 取数", o)
	}

	o1, ok := <-intCh
	fmt.Println(o1, ok)

}

func TestSelectChannel(t *testing.T) {
	//定义两个channel
	ch1 := make(chan int)
	ch2 := make(chan string)
	fmt.Println("start time:", time.Now())

	//goroutine ，睡一秒，往ch1 写数据
	go func() {
		time.Sleep(1 * time.Second) //等待1秒
		ch1 <- 1                    //装一个数据
	}()

	//goroutine ，睡两秒，往ch2 写数据
	go func() {
		time.Sleep(2 * time.Second) //等待1秒
		ch2 <- "GOGOGO"
	}()

	fmt.Println("select time:", time.Now())

	//谁准备好了，走谁的！
	select {
	case out := <-ch1:
		fmt.Println(time.Now(), "ch1 ready, go", out)
	case out := <-ch2:
		fmt.Println(time.Now(), "ch2 ready, go", out)
	}
	fmt.Println("Done")
}

func TestSelectChannelWithDefault(t *testing.T) {
	//定义两个channel
	ch1 := make(chan int)
	ch2 := make(chan string)
	fmt.Println("start time:", time.Now())

	//goroutine ，睡一秒，往ch1 写数据
	go func() {
		time.Sleep(1 * time.Second) //等待1秒
		ch1 <- 1                    //装一个数据
	}()

	//goroutine ，睡两秒，往ch2 写数据
	go func() {
		time.Sleep(2 * time.Second) //等待1秒
		ch2 <- "GOGOGO"
	}()

	fmt.Println("select time:", time.Now())

	//谁准备好了，走谁的！
	select {
	case out := <-ch1:
		fmt.Println(time.Now(), "ch1 ready, go", out)
	case out := <-ch2:
		fmt.Println(time.Now(), "ch2 ready, go", out)
	default:
		fmt.Println(time.Now(), "所有的channel都不ready，直接走default")
	}
	fmt.Println("Done")
}

func TestSelectChannelWithDefaultAndChannelReady(t *testing.T) {
	//定义两个channel
	ch1 := make(chan int, 1) //给1个size，放数据
	ch2 := make(chan string)
	fmt.Println("start time:", time.Now())

	ch1 <- 1
	//goroutine ，睡一秒，往ch1 写数据
	go func() {
		time.Sleep(1 * time.Second) //等待1秒
		ch1 <- 1                    //装一个数据
	}()

	//goroutine ，睡两秒，往ch2 写数据
	go func() {
		time.Sleep(2 * time.Second) //等待1秒
		ch2 <- "GOGOGO"
	}()

	fmt.Println("select time:", time.Now())

	//谁准备好了，走谁的！
	select {
	case out := <-ch1:
		fmt.Println(time.Now(), "ch1 ready, go", out)
	case out := <-ch2:
		fmt.Println(time.Now(), "ch2 ready, go", out)
	default:
		fmt.Println(time.Now(), "所有的channel都不ready，直接走default")
	}
	fmt.Println("Done")
}

func TestSelectChannelWithDefaultAndCloseChannel(t *testing.T) {
	//定义两个channel
	ch1 := make(chan int)
	ch2 := make(chan string)
	fmt.Println("start time:", time.Now())

	//goroutine ，睡一秒，往ch1 写数据
	//go func() {
	//	time.Sleep(1 * time.Second) //等待1秒
	//	ch1 <- 1                    //装一个数据
	//}()
	fmt.Println("close ch1")
	close(ch1)

	//goroutine ，睡两秒，往ch2 写数据
	go func() {
		time.Sleep(2 * time.Second) //等待1秒
		ch2 <- "GOGOGO"
	}()

	fmt.Println("select time:", time.Now())

	//谁准备好了，走谁的！
	select {
	case out := <-ch1:
		fmt.Println(time.Now(), "ch1 ready, go", out)
	case out := <-ch2:
		fmt.Println(time.Now(), "ch2 ready, go", out)
	default:
		fmt.Println(time.Now(), "所有的channel都不ready，直接走default")
	}
	fmt.Println("Done")
}

func TestMultipleSelect(t *testing.T) {
	ch1 := make(chan int)

	for i := 0; i < 10; i++ {
		go func(i int) {
			select {
			case <-ch1:
				fmt.Println(time.Now(), i)
			}
		}(i)
	}

	fmt.Println("关闭Channel", time.Now())
	close(ch1)

	time.Sleep(1 * time.Second)
}

func TestDualCloseChannel(t *testing.T) {
	c := make(chan struct{})
	close(c)
	close(c)
}

func TestOutonly(t *testing.T) {
	intCh := make(chan int, 10)
	<-intCh
}

func TestSingleGoroutineApp(t *testing.T) {
	intCh := make(chan int)
	intCh <- 1
	<-intCh
}

func TestMultipleChannelReadySelect(t *testing.T) {
	ch1, ch2 := make(chan int), make(chan int)
	close(ch1)
	close(ch2)
	//关闭掉两个channel，意味着，永远都是ready的状态

	ch1Counter, ch2Counter := 0, 0
	for i := 0; i < 10000; i++ {
		select {
		case <-ch1:
			ch1Counter++
		case <-ch2:
			ch2Counter++
		}
	}
	fmt.Println("ch1Counter:", ch1Counter, "ch2Counter:", ch2Counter)
}
