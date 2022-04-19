package main

import (
	"fmt"
	"sync"
	"time"
)

//首先要有一个仓库(共享仓库)
type Store struct {
	DataCount int        //仓库里面存放的东西
	Max       int        //库存最大值
	lock      sync.Mutex //生产者和消费者都是往里面写的，并非多读少写的锁，用同步锁即可
}

//生产者
type Producer struct {
}

//生产者的成员函数,生产者没有其他成员变量，可以不用写名字
//这里传入的是Store的指针，因为我们所有的操作都是基于Store做的事情，如果不用指针，对象的数据不会被更改
func (Producer) Produce(s *Store) {
	s.lock.Lock()         //拿锁
	defer s.lock.Unlock() //最后再解锁
	//查看库存，如果库存满了,就直接return掉，就不生产；没有满就生产，+1
	if s.DataCount == s.Max {
		fmt.Println("生产者看到库存已经满了，不生产")
		return
	}
	fmt.Println("开始生产 +1")
	s.DataCount++
}

type Consumer struct {
}

//也是用指针，不然Store对象的数据不会被更改
func (Consumer) Consume(s *Store) {
	s.lock.Lock()
	defer s.lock.Unlock()
	//查看库存，如果库存为哦，就直接return掉，不再消费；如果没有满，就-1
	if s.DataCount == 0 {
		fmt.Println("消费者看到没有库存了，不消费")
		return
	}
	fmt.Println("消费者消费 -1")
	s.DataCount--
}

//定义mian函数 串起来
func main() {
	s := &Store{
		Max: 10,
		//这里注意lock会自动初始化，因为引用了一个sync.Mutex对象，
		//当初始化store的时候，里面的对象也被初始化了，
		//如果是一个引用类型的，如果是map，*sync.Mutex，就需要专门为它做初始化！！！
	}
	//需要多个producer和consumer，各来50个
	pCount,cCount := 50,50
	for i :=0;i<pCount;i++{
		go func() { //goroutine
			for {
				time.Sleep(500*time.Microsecond)
				Producer{}.Produce(s) //去调用对象下的成员函数,这里是生产正生产
			}
		}()
	}
	for i :=0;i<cCount;i++{
		go func() { //goroutine
			for{
				time.Sleep(500*time.Microsecond)
				Consumer{}.Consume(s) //去调用对象下的成员函数，这里是消费者消费
			}
		}()
	}
	time.Sleep(1*time.Second)
	fmt.Println("库存余量：",s.DataCount)
}
