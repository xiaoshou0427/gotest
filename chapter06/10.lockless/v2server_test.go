package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWebServerV2(t *testing.T) {
	v2 := &WebServerV2{
		config: &Config{Content: "a"},
	} //初始化，这里要初始化一个值，跟版本一不同！！！
	go v2.ReloadWorker() //启动这个webserver 写操作的worker，此时就会不停的，每隔10毫秒，去跑这个写操作（写配置哦）
	//上面是写操作，每10毫秒
	//现在需要起很多个worker去访问这个webserver,起1w个worker来读
	//统计单位时间内做的多还是少，来判断它的性能 //
	//每个goroutine 访问1千次，来看下总消耗时间是多少，来确定它的性能！
	start := time.Now()          //记录启动时间，
	wg := sync.WaitGroup{}       //用waitgroup！计数，不做完，不往下走！
	wg.Add(10000)                //计数1w次
	for i := 0; i < 10000; i++ { //循环1w次，并行1w个goroutine
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ { //每个goroutine 循环执行1000次
				v2.Visit() //调用访问
			}
		}()
	}
	wg.Wait()
	finish := time.Now()
	fmt.Println("totally time:", finish.Sub(start))
}
