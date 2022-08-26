package main

import (
	"fmt"
	"sync"
)

type Queue struct {
	sync.Mutex //实现线程安全
	data       []interface{}
}

//对数据进行变更，采用指针的方式
func (q *Queue) Push(data interface{}) {
	q.Lock()                      //加锁
	defer q.Unlock()              //最后解锁
	q.data = append(q.data, data) //往data 里面push 数据，追加到切片的最后面
}

func (q *Queue) Pop() (interface{}, bool) {
	q.Lock()
	defer q.Unlock()
	if len(q.data) > 0 {
		o := q.data[0]
		q.data = q.data[1:] //切一下，把最初的数据从data中切出
		return o, true
	}
	return nil, false //不大于0 返回空
}

func main() {
	q := &Queue{} //为啥加指针？ 往里面写东西？改变数据？没有这个& 一样可以执行！
	q.Push(111)
	q.Push(222)
	q.Push(333)
	q.Push(nil)

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

}
