package main

import (
	"fmt"
	"sync"
	"time"
)

//思考一下，这样累加，多个进程，你这个数不一定对啊
func main() {
	fmt.Println(time.Now())
	for i := 0; i < 1; i++ { //跑十遍下面的函数
		//countDict()
		//countForgetUnlock(
		//countDictLockPrice()
		countDictGoroutineSafeRW()
	}
	fmt.Println(time.Now())

}

func countDict() {
	fmt.Println("开始数")
	var totalCount int64 = 0       //整体字数
	totalCountLock := sync.Mutex{} //锁,起的名字有讲究，用变量名+Lock，就知道这个锁跟哪些是有紧密的关联关系
	wg := sync.WaitGroup{}
	wg.Add(3)                //有5000件事情要做，或者说有5000个人做事情
	for p := 0; p < 3; p++ { //5000页循环
		go func(p int) { //启动goroutine，循环一次，开一个goroutine，循环5000次，开5000个goroutine
			defer wg.Done() //表示一个人或者这件事做完了，defer最后再做
			//defer totalCountLock.Unlock() //最后再解锁！
			//fmt.Println("正在统计", p, "页")
			totalCountLock.Lock()   //锁起来
			totalCount += 100       //累加整体的字数,固定每一页100个字
			totalCountLock.Unlock() //释放锁
			//if p == 0 { //  当p为0，即其中一个goroutine 睡10秒
			time.Sleep(1 * time.Second)
			//}
		}(p)
	} //这里没有sleep，会跑的非常快
	wg.Wait() //等所有人做完了，才执行下一行
	//我们是有预期的，每一页100个字，那么5000页就是50w字
	fmt.Println("预计有", 100*5000, "字")
	fmt.Println("总共有 ", totalCount, "字")
}

func countDictGoroutineSafeRW() {
	fmt.Println("开始数")
	var totalCount int64 = 0         //整体字数
	totalCountLock := sync.RWMutex{} //读写锁,起的名字有讲究，用变量名+Lock，就知道这个锁跟哪些是有紧密的关联关系
	wg := sync.WaitGroup{}
	workerCount := 5
	wg.Add(workerCount) //有workerCount件事情要做，或者说有workerCount个人做事情
	doneCh:=make(chan struct{}) //定义一个channel，这里不定义，读锁和写锁就没有关联了？
	for p := 0; p < workerCount; p++ {
		go func(p int) {
			fmt.Println(p,"读锁开始时间：",time.Now())
			totalCountLock.RLock() //用read lock
			fmt.Println(p,"拿到读锁的时间：",time.Now())
			time.Sleep(1*time.Second)
			totalCountLock.RUnlock() //用read unlock
		}(p)
	}

	for p := 0; p < workerCount; p++ { //workerCount页循环
		go func(p int) { //启动goroutine，循环一次，开一个goroutine，循环workerCount次，开workerCount个goroutine
			defer wg.Done() //表示一个人或者这件事做完了，defer最后再做
			//写锁总是只有一个人能拿到！！！总是会看到某个worker开始，某个worker结束，下一个worker开始，下一个worker结束
			fmt.Println(p, "写锁开始时间", time.Now())
			totalCountLock.Lock() //锁起来
			totalCount += 100     //累加整体的字数,固定每一页100个字
			fmt.Println(p, "拿到写锁的时间", time.Now())
			totalCountLock.Unlock() //释放锁
			//if p == 0 { //  当p为0，即其中一个goroutine 睡10秒
			time.Sleep(1 * time.Second)
			//}
		}(p)
	} //这里没有sleep，会跑的非常快
	wg.Wait() //等所有人做完了，才执行下一行
	close(doneCh) //关闭channel
	//我们是有预期的，每一页100个字，那么5000页就是50w字
	fmt.Println("预计有", 100*5000, "字")
	fmt.Println("总共有 ", totalCount, "字")
}

func countForgetUnlock() {
	fmt.Println("开始数")
	var totalCount int64 = 0       //整体字数
	totalCountLock := sync.Mutex{} //锁,起的名字有讲究，用变量名+Lock，就知道这个锁跟哪些是有紧密的关联关系
	wg := sync.WaitGroup{}
	wg.Add(5)                //有5件事情要做，或者说有5000个人做事情
	for p := 0; p < 5; p++ { //5页循环
		go func() { //启动goroutine，循环一次，开一个goroutine，循环5000次，开5000个goroutine
			defer wg.Done() //表示一个人或者这件事做完了，defer最后再做
			//fmt.Println("正在统计", p, "页")
			totalCountLock.Lock() //锁起来
			totalCount += 100     //累加整体的字数,固定每一页100个字
			//totalCountLock.Unlock() //释放锁
		}()
	} //这里没有sleep，会跑的非常快
	wg.Wait() //等所有人做完了，才执行下一行
	//我们是有预期的，每一页100个字，那么5000页就是50w字
	fmt.Println("预计有", 100*5000, "字")
	fmt.Println("总共有 ", totalCount, "字")
}

func countDictLockPrice() {
	fmt.Println("开始数")
	var totalCount int64 = 0       //整体字数
	totalCountLock := sync.Mutex{} //锁,起的名字有讲究，用变量名+Lock，就知道这个锁跟哪些是有紧密的关联关系
	wg := sync.WaitGroup{}
	wg.Add(5)                //有5件事情要做，或者说有5000个人做事情
	for p := 0; p < 5; p++ { //5页循环
		go func(pageNum int) { //启动goroutine，循环一次，开一个goroutine，循环5000次，开5000个goroutine
			defer wg.Done() //表示一个人或者这件事做完了，defer最后再做
			//fmt.Println("正在统计", p, "页")
			totalCountLock.Lock() //锁起来
			totalCount += 100     //累加整体的字数,固定每一页100个字
			if pageNum == 3 {
				time.Sleep(3 * time.Second) //当页面为3的时候，等3秒再unlock
			}
			totalCountLock.Unlock() //释放锁
		}(p) //这里p call pageNum ，可以理解为pageNum = p ？
	} //这里没有sleep，会跑的非常快
	wg.Wait() //等所有人做完了，才执行下一行
	//我们是有预期的，每一页100个字，那么5000页就是50w字
	fmt.Println("预计有", 100*5000, "字")
	fmt.Println("总共有 ", totalCount, "字")
}
