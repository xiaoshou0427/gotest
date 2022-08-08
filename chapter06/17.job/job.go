package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), 0)
	defer cancel()
	successFlag := make(chan bool, 1) //验证是否成功！ 做个channel
	//执行多个goroutine 完成所有功能
	go account(ctx)
	go distributeService(ctx)
	go configure(ctx)
	go verifyService(ctx, successFlag)

	select {
	case <-ctx.Done(): //这个一定是等到最后才ready的！因为defer cancel()
		fmt.Println("超时，没有完成")
	case v := <-successFlag: //取出successFlag 是true还是false
		if v {
			fmt.Println("任务完成，成功结束")
		} else {
			fmt.Println("任务失败，需要重新考虑重试，还是下线服务")
		}
	}
}

//账号管理
func account(ctx context.Context) {
	//做管控
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	doneCh := make(chan string, 2) //两个功能。buffer buffer 用2 就行
	go accountRegister(ctx, doneCh)
	go accountGrantAuth(ctx, doneCh)

	successCount := 0       //计数器
	for v := range doneCh { //循环取出channel的信息
		successCount++ //不断计数
		fmt.Println("job:", v, "done")
		if successCount == 2 { //当为2的时候，关闭channel！
			close(doneCh)
		}
	}
	fmt.Println("账号处理完成")
}

//创建新账号
func accountRegister(ctx context.Context, doneCh chan string) {
	fmt.Println("注册账号")
	defer fmt.Println("注册完成")
	for { //有一定的重试，但是要防范无休止的重试 （这里我们就无限循环了！）
		//。。调用xxx的接口，调用外部的总是不太可控！
		select {
		case <-ctx.Done():
			fmt.Println("context 结束， 不再注册")
			return //函数返回，不再执行函数的内容！
		default: //不会卡在这，而是继续往下走
		}
		doneCh <- "accountRegister"
		fmt.Println("accountRegister成功")
		break // for 循环，如果注册了，就break这个循环了！ 这里没有做判断，直接退出for循环
	}
}

//授权新账号
func accountGrantAuth(ctx context.Context, doneCh chan string) {
	fmt.Println("授权账号")
	defer fmt.Println("授权账号完成")
	for { //。。调用xxx的接口
		select {
		case <-ctx.Done():
			fmt.Println("context 结束， 不再授权")
			return
		default: //不会卡在这，而是继续往下走
		}
		doneCh <- "accountGrantAuth"
		fmt.Println("accountGrantAuth成功")
		break
	}
}

//部署服务
//7分钟内部署完毕，否则超时
func distributeService(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 7*time.Minute)
	defer cancel()
	//这里可以用wait group 来完成计数！等待goroutine做完工作
	// 这是不用channel的方式
	wg := sync.WaitGroup{}
	wg.Add(2)
	go distributeLB(ctx, &wg)
	go distributeInstance(ctx, &wg)
	wg.Wait() //goroutine完成计数就继续往下
	fmt.Println("distributeService Done")

}

//如果不想尊重context 可以func distributeInstance(_ context.Context, w *sync.WaitGroup)
func distributeInstance(ctx context.Context, w *sync.WaitGroup) {
	defer w.Done()
	for {
		select { //尊重上下文，任务在什么时候可以结束
		case <-ctx.Done():
			fmt.Println("上下文结束，要删除已经创建的实例")
			return
		default:
		}
		fmt.Println("部署实例")
		break
	}
}

func distributeLB(ctx context.Context, w *sync.WaitGroup) {
	defer w.Done()
	for {
		select { //尊重上下文，任务在什么时候可以结束
		case <-ctx.Done():
			fmt.Println("上下文结束，要删除已经创建的负载均衡器")
			return
		default:
		}
		fmt.Println("部署负载均衡器")
		break
	}
}

//可以先忽略这个ctx
func configure(_ context.Context) {
	fmt.Println("注入新服务账号")
}

func verifyService(ctx context.Context, flag chan bool) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go verifyFunction(ctx, &wg)
	wg.Wait()
	fmt.Println("验证服务完成")
	flag <- true //最终一定是成功的，所以就尊重这个flag给个true！
}

func verifyFunction(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	//只重试三次
	for i := 0; i < 3; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("上下文结束，不再验证")
			return
		default:
		}
		fmt.Println("开始验证服务")
		time.Sleep(100 * time.Millisecond) //用来替换验证部分的环节，比如：服务调用，服务模拟等等

		if i <= 1 {
			fmt.Println("服务尚未完成，重试中。。。")
			continue //跳入下一次循环，到i=2的时候，不执行这里的内容，跳到break了！
		}
		break
	}
}

//这里是模拟验证结果，不希望太顺利的验证完毕，就可以用if i<=1操作。
