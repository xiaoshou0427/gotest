package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//withCancel()
	//withTimeout()
	//withValue()
	withDeadline()

}

func withDeadline() {
	now := time.Now()
	newtime := now.Add(1 * time.Second)                     //加一秒
	ctx, _ := context.WithDeadline(context.TODO(), newtime) //1秒后自动结束
	go tv(ctx)
	go mobile(ctx)
	go game(ctx)

	time.Sleep(2 * time.Second) //睡两秒保证goroutine正常运行
}

func tv(ctx context.Context) {
	for { //这里做了个for循环,一直在看电视
		select {
		case <-ctx.Done():
			fmt.Println("关电视")
			return
		default: //继续往下走
		}
		fmt.Println("看电视")
		time.Sleep(300 * time.Microsecond) //每300毫秒输出看电视（一直在看电视）
	}
}

func mobile(ctx context.Context) {
	for { //这里做了个for循环,一直在看手机
		select {
		case <-ctx.Done():
			fmt.Println("关手机")
			return
		default: //继续往下走
		}
		fmt.Println("看手机")
		time.Sleep(300 * time.Microsecond) //每300毫秒输出看手机（一直在看手机）
	}
}

func game(ctx context.Context) {
	for { //这里做了个for循环
		select {
		case <-ctx.Done():
			fmt.Println("关游戏")
			return
		default: //继续往下走
		}
		fmt.Println("玩游戏")
		time.Sleep(300 * time.Microsecond)
	}
}

func withValue() {
	//这里省略了ctx := context.TODO()，简写
	ctx := context.WithValue(context.TODO(), "1", "钱包")
	go func(ctx context.Context) {
		time.Sleep(1 * time.Second)
		fmt.Println("withValue：1", ctx.Value("1"))
		fmt.Println("withValue：2", ctx.Value("2"))
		fmt.Println("withValue：3", ctx.Value("3"))
		fmt.Println("withValue：4", ctx.Value("4"))
	}(ctx)
	goToPapa(ctx)
	time.Sleep(2 * time.Second)
}

func goToPapa(ctx context.Context) {
	ctx = context.WithValue(ctx, "2", "充电宝")
	go func(ctx context.Context) {
		time.Sleep(1 * time.Second)
		fmt.Println("goToPapa：1", ctx.Value("1"))
		fmt.Println("goToPapa：2", ctx.Value("2"))
		fmt.Println("goToPapa：3", ctx.Value("3"))
		fmt.Println("goToPapa：4", ctx.Value("4"))
	}(ctx)
	goToMama(ctx)
}

func goToMama(ctx context.Context) {
	ctx = context.WithValue(ctx, "3", "小夹克")
	go func(ctx context.Context) {
		time.Sleep(1 * time.Second)
		fmt.Println("goToMama：1", ctx.Value("1"))
		fmt.Println("goToMama：2", ctx.Value("2"))
		fmt.Println("goToMama：3", ctx.Value("3"))
		fmt.Println("goToMama：4", ctx.Value("4"))
	}(ctx)
	goToGrandma(ctx)
}

func goToGrandma(ctx context.Context) {
	ctx = context.WithValue(ctx, "4", "大苹果")
	go func(ctx context.Context) {
		time.Sleep(1 * time.Second)
		fmt.Println("goToGrandma：1", ctx.Value("1"))
		fmt.Println("goToGrandma：2", ctx.Value("2"))
		fmt.Println("goToGrandma：3", ctx.Value("3"))
		fmt.Println("goToGrandma：4", ctx.Value("4"))
	}(ctx)
	goToParty(ctx)
}

func goToParty(ctx context.Context) {
	fmt.Println("goToParty：1", ctx.Value("1"))
	fmt.Println("goToParty：2", ctx.Value("2"))
	fmt.Println("goToParty：3", ctx.Value("3"))
	fmt.Println("goToParty：4", ctx.Value("4"))
}

func withTimeout() {
	ctx, _ := context.WithTimeout(context.TODO(), 1*time.Second)
	//这里只等待1秒，就timeout，自动cancel哦！所以这里不用写canncel，原因可以把鼠标放到ctx.Done上面
	fmt.Println("开始部署望远镜，发送信号")
	//这里执行goroutine哦！ 跑到另外一个进程上去跑！
	go distributeMainFrame(ctx)
	go distributeMainBody(ctx)
	go distributeCover(ctx)

	select {
	case <-ctx.Done(): //这里是一定会做的，等待1秒后执行这个！
		fmt.Println("任务超时没有完成")
	}

	time.Sleep(20 * time.Second) //这里等待2秒，主程序退出
}

func distributeMainFrame(ctx context.Context) {
	time.Sleep(10 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("任务取消：distrubuteMainFrame")
		return
	default: //不阻塞，继续往下走
	}
	fmt.Println("部署主框架")
}

func distributeMainBody(ctx context.Context) {
	time.Sleep(10 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("任务取消：distributeMainBode")
		return
	default: //不阻塞，继续往下走
	}
	fmt.Println("部署旋转主体")
}

func distributeCover(ctx context.Context) {
	time.Sleep(10 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("任务取消：distributeCover")
		return
	default: //不阻塞，继续往下走
	}
	fmt.Println("部署遮阳板")
}

//这里也是只做了前三个测试

func withCancelinstruction() {
	//看源码，两个的写法是一样的，但是语义是不同的！要了解语义！
	context.TODO()       //这个是还没有定义
	context.Background() //这个是已经定义
}

func withCancel() {
	ctx := context.TODO() //ctx（常用这个名字）定义一个新的context
	ctx, cancel := context.WithCancel(ctx)
	fmt.Println("做蛋挞，要买材料")
	// 启动goroutine,每个成员去买东西都会尊重这个context （ctx）
	go buyFlour(ctx)
	go buyOil(ctx)
	go buyEgg(ctx)
	time.Sleep(500 * time.Microsecond) //防止运行过快
	fmt.Println("没电了，取消购买所有食材")
	cancel() //调用这个cancel function ，上面定义的cancel

	time.Sleep(2 * time.Second)
}

//参数传入context
func buyFlour(ctx context.Context) {
	fmt.Println("去买面 ")
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done(): //todo 介绍一下 struct{}
		fmt.Println("收到消息，不买面了")
		return
	default: //不阻塞，继续往下走
	}
	fmt.Println("Buy flour")
}

func buyOil(ctx context.Context) {
	fmt.Println("去买油 ")
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买油了")
		return
	default: //不阻塞，继续往下走
	}
	fmt.Println("Buy Oil")
}

func buyEgg(ctx1 context.Context) {
	ctx, _ := context.WithCancel(ctx1) //这里定义了新的context，跟上面的context不是同一个，这里也可以用ctx1 来区别上面的ctx
	//defer cancel()  //这里一定会去执行的，就看你的goroutine执行的时间，执行时间长，这个主routine退出，就没有效果了
	fmt.Println("去买蛋 ")
	//time.Sleep(1 * time.Second) //这里隐藏了，不隐藏，会直接跑到defer，goroutine还没执行
	select {
	case <-ctx.Done(): //因为新的cancel 是defer定义的，所以这里不执行
		fmt.Println("收到消息，不买蛋了")
		return
	default: //不阻塞，继续往下走
	}
	fmt.Println("Buy Egg")
	//这里注意了， 去买蛋，买大蛋和小蛋
	go buySEgg(ctx)
	go buyBEgg(ctx)
	//这里原始是没有关联context，需要在这个函数的最上面增加context，类似于嵌套
	time.Sleep(1 * time.Second)

}

//先买这三样

//small egg
func buySEgg(ctx context.Context) {
	fmt.Println("去买小蛋 ")
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买小蛋了")
		return
	default: //不阻塞，继续往下走
	}
	fmt.Println("Buy 小Egg")
}

//big egg
func buyBEgg(ctx context.Context) {
	fmt.Println("去买大蛋 ")
	time.Sleep(1 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("收到消息，不买大蛋了")
		return
	default: //不阻塞，继续往下走
	}
	fmt.Println("Buy 大Egg")
}
