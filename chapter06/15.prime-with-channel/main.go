package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	startTime := time.Now() //开始时间
	//result := []int{}                    //统计结果
	result := make(chan int, 200000) //设置一个接数据的channel

	workerNumber := 16                //起16个
	baseNumCh := make(chan int, 1024) //给一个size 是1024的buffer

	wg := sync.WaitGroup{}
	wg.Add(workerNumber) //16个goroutine

	//启动worker （先把worker启动好）
	for i := 0; i < workerNumber; i++ {
		go func() {
			defer wg.Done()
			for oNum := range baseNumCh { //一个goroutine，循环取channel里面的数据！ 读取数据哦！
				if isPrime(oNum) { //调用函数，如果为真，执行下面的操作
					result <- oNum //算出是素数的，存入result 的 channel中
				}
			}
		}()
	}

	//再往里面装数据！（先启动了worker），buffer 是1024，当buffer装满，就装不进去数据了！
	for num := 2; num <= 200000; num++ { //不停的循环，将2-200000 数字写入到baseNumCh 这个channel中
		baseNumCh <- num //写入到baseNumCh ,这里循环写入，没有buffer 就写入一个，取一个，才再写入下一个，再取走下一个。
		//这里buffer 1024，那么就先放1024个进到buffer，然后每个goroutine 从buffer循环取数据，取走的，就会从baseNumCh中移除
		//这里注意channel 里面的数据取走了，就没有了！所以每个goroutine 是直接从buffer里面取数据的！channel 的buffer 写入1024个数据，那么多个goroutine
		//就会消费这些数据，并计算结果，写入另外一个channel。
		//go routine 没循环完 baseNumCh 就不会结束，而for num 是主程序的循环，也在不断的执行，那么go routine不断的从baseNumCh获取数据，直到消费完成。
		//主进程for 循环完毕后，关闭channel，此时还是可以读取channel的哦，goroutine读取完baseNumCH，计算完毕后，写入result，此时wg.Done了，才继续向下进行。
	}
	//循环完毕，关闭channel！ 这个是一定要做的！因为你goroutine里面range这个channel了（看笔记）！不然报错（fatal error: all goroutines are asleep - deadlock!）！
	close(baseNumCh) //关闭channel
	wg.Wait()

	finishTime := time.Now()                       //for循环退出，当前时间
	fmt.Println("从2到200000，共有素数：", len(result))    //一共有多少个素数
	fmt.Println("共耗时:", finishTime.Sub(startTime)) //算出耗时
}

func isPrime(num int) (isPrime bool) { //返回true还是false
	for i := 2; i < num; i++ {
		if num%i == 0 { //如果能被i整除，就不是素数！
			isPrime = false //返回false 不记录这个num！
			return
		}
	}
	isPrime = true
	return
}
