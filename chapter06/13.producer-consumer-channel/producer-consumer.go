package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//首先要有一个仓库(共享仓库)
type Store struct {
	init  sync.Once //初始化一次channel
	store chan int  //定义一个channel
	//不需要 DataCount int        //仓库里面存放的东西
	Max int //库存最大值
	//不需要 lock      sync.Mutex //生产者和消费者都是往里面写的，并非多读少写的锁，用同步锁即可
}

//做，且只做一次的。专门做一次初始化
func (s *Store) instrument() { //修改store里面的内容，用指针
	s.init.Do(func() {
		s.store = make(chan int, s.Max) //Max 有多少个最大的库存，作为size
	})
}

//生产者
type Producer struct {
}

//
func (Producer) Produce(s *Store) {
	fmt.Println("开始生产 +1")
	s.store <- rand.Int() //随机数往里放
}

type Consumer struct {
}

//
func (Consumer) Consume(s *Store) {
	fmt.Println("消费者消费 -1", <-s.store) //直接输出随机的生产者
}

//定义mian函数 串起来
func main() {
	s := &Store{
		Max: 10, //初始化10个库存
	}
	s.instrument()
	//需要多个producer和consumer，各来50个(go routine)
	pCount, cCount := 50, 50
	for i := 0; i < pCount; i++ {
		go func() { //goroutine
			for {
				time.Sleep(500 * time.Microsecond)
				Producer{}.Produce(s) //去调用对象下的成员函数,这里是生产正生产
			}
		}()
	}
	for i := 0; i < cCount; i++ {
		go func() { //goroutine
			for {
				time.Sleep(500 * time.Microsecond)
				Consumer{}.Consume(s) //去调用对象下的成员函数，这里是消费者消费
			}
		}()
	}
	time.Sleep(1 * time.Second)
}
