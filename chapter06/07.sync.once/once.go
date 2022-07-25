package main

import (
	"fmt"
	"sync"
)

type rank struct {
	standard []string
}

var globalRank = &rank{}

//初始化一个once
var once sync.Once = sync.Once{}

//一个once.Do 解决，没有用锁，即使下面 goroutine 都调用这个函数
//这里once 只跑一次！！！！ 不会多跑一次！尤其是全局变量，如数据库初始化，用once做全局的配置。
//不管怎么跑，都是没有锁的！性能非常好！
//once.Do ( ) 以func 传入！！！
func initGlobalRankStandard(standard []string) {
	once.Do(func() {
		globalRank.standard = standard
	})

}

var facStore = &dbFactoryStore{}

type dbFactoryStore struct {
	store map[string]DBFactory
}

type Conn struct {
}

type DBFactory interface {
	GetConnection() *Conn
}

func initMySqlFac(connStr string) DBFactory {
	return &MySqlDBFactory{}
}

type MySqlDBFactory struct {
	once sync.Once
}

func (MySqlDBFactory) GetConnection() *Conn {
	once.Do(func() {
		initMySqlFac("")
	})
	//todo
	return nil
}

var counter int = 0
var counterOnce sync.Once

type School struct {
	classroomLocation map[string]string //初始化一次就够了，这里只是举例说明，没有引用
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("第x次：", i)
		counterOnce.Do(func() {
			fmt.Println("初始化")
			counter++
		})
	}
	fmt.Println("最终结果：", counter)
	////程序注入一个standard
	//standard := []string{"asia"}
	////使用goroutine 就有多个线程再跑上面的函数，都会去执行锁和解锁，判断，无形中会成为整个系统的瓶颈！
	////而且上面func initGlobalRankStandard写的很长很长
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		initGlobalRankStandard(standard)
	//	}()
	//}
	//connStr := "xxxxx"
}
