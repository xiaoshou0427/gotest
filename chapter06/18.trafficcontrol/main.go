package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 200; i++ { //2个goroutine跑，这里调整为200个并发，channel的buffer改成10 不用1w，那就会出现请求被拒绝的情况
		go sendRequest()
	}
	time.Sleep(5 * time.Second) //让goroutine跑5秒，主程序退出！
}

//不停发送请求,这个发送时间短一些，而处理时间也就是server request长一些，就可以看出服务被拒绝！？
func sendRequest() {
	for { //不断循环执行！
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Int63()%30) * time.Microsecond) //那为啥这里用30？老师写的30
		serverRequest()                                               //随机时间后执行serverRequest
	}
}

//全局变量无法做动态调整（自己考虑）
var jobcontrolCh = make(chan struct{}, 10)

func serverRequest() {
	accept := trafficeControl_Start() //服务请求是否被接受
	if accept {                       //如果接受请求，才会进行转发或者说向外发出流量
		defer trafficeControl_Finish()
		fmt.Println("服务请求")
		rand.Seed(time.Now().UnixNano())
		//为啥这么写？？？
		time.Sleep(time.Duration(rand.Int63()%30) * time.Microsecond) //模拟不同服务器返回时间 ,30毫秒以内完成服务
	} else {
		fmt.Println("服务请求被拒绝")
	}
}

//不停的发送请求，调用某个参数说要做某件事情，defer取执行finish动作
//先start后finish

func trafficeControl_Start() (accept bool) {
	select {
	case jobcontrolCh <- struct{}{}: //不断的加，buffer满了就不加了？？？？
		fmt.Println("接受请求")
		return true //接受请求在这里返回 true
	default: //满了就拒绝了
		fmt.Println("拒绝请求")
		return //否则返回false ， 默认不写 ，为空，布尔值默认是false
	}
}

//请求完成，往外出流量
func trafficeControl_Finish() {
	<-jobcontrolCh //从channel 踢出去 ，不需要变量接受这个数据
}
