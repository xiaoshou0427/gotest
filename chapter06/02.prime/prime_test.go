package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestRunPrime(t *testing.T) {
	startTime := time.Now()              //开始时间
	result := []int{}                    //统计结果
	for num := 2; num <= 200000; num++ { //开始计算，从二开始统计
		if isPrime(num) {
			result = append(result, num) //算出是素数的，追加到列表
		}
	}
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

func TestRunPrime2(t *testing.T) {
	startTime := time.Now()
	result := []int{}
	go func() {
		fmt.Println("第一个worker开始计算", time.Now())
		result = append(result, collectPrime(2, 100000)...)
		fmt.Println("第一个worker完成计算", time.Now())

	}()
	go func() {
		fmt.Println("第二个worker开始计算", time.Now())
		result = append(result, collectPrime(100001, 200000)...)
		fmt.Println("第二个worker完成计算", time.Now())

	}()
	time.Sleep(15 * time.Second) //睡15秒
	finishTime := time.Now()
	fmt.Println(len(result))
	fmt.Println("共耗时:", finishTime.Sub(startTime))
}

func collectPrime(start int, end int) (result []int) { //上面append 应该是返回一堆，加入到列表，所以这里也是个列表
	for num := start; num <= end; num++ { //开始计算，从二开始统计
		if isPrime(num) {
			result = append(result, num) //算出是素数的，追加到列表
		}
	}
	return //这里不用返回变量result啦！多此一举
}

func TestRunPrime3(t *testing.T) {
	startTime := time.Now()
	result := []int{}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("第一个worker开始计算", time.Now())
		result = append(result, collectPrime(2, 100000)...)
		fmt.Println("第一个worker完成计算", time.Now())
	}()
	go func() {
		defer wg.Done()
		fmt.Println("第二个worker开始计算", time.Now())
		result = append(result, collectPrime(100001, 200000)...)
		fmt.Println("第二个worker完成计算", time.Now())
	}()
	wg.Wait()
	finishTime := time.Now()
	fmt.Println(len(result))
	fmt.Println("共耗时:", finishTime.Sub(startTime))
}

func TestHelloGoroutine(t *testing.T) {
	go fmt.Println("Hello goroutine")
}

func TestHelloGoroutine2(t *testing.T) {
	go fmt.Println("Hello goroutine")
	time.Sleep(1 * time.Second)
}

func TestForLoop(t *testing.T) {
	go func() {
		for i := 0; i < 5; i++ { //goroutine
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}
	}()
	for i := 100; i < 120; i++ { //主routine，循环时间长!
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}

func TestForeverGoroutine(t *testing.T) {
	go func() { //goroutine
		for {
			time.Sleep(1 * time.Second)  //睡1s
			go func() { //循环里面继续goroutine
				fmt.Println("启动新的goroutine@", time.Now())
				time.Sleep(1 * time.Hour) //执行1小时就结束了，不要无限下去
			}()
		 }
	  }()
	for { //主routine 可以理解为主进程
		fmt.Println(runtime.NumGoroutine()) //for 永久跑这个，看有多少个goroutine存在
		time.Sleep(2 * time.Second)  //与上面的独立goroutine区分一下，这里不加你就尴尬

	}
}
