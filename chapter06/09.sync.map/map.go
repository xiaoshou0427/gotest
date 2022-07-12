package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := sync.Map{}
	for i := 0; i < 1; i++ {
		go func(i int) { //循环开100个goroutine，每个routine的i成为一个key
			m.Store(i, 1) //这里给第一个key一个value为1 存储起来，避免空指针
			for {         //增加for循环，目的是每个goroutine不断的给v加1！
				v, _ := m.Load(i) //这里的i是key，出参给的是value和bool
				//这里的v 已经load了，那么v就是等于1
				m.Store(i, v.(int)+1)                        //这里要强转，因为v 是 interface类型，这里是存储key的value
				fmt.Printf("key= %d,value=%d\n", i, v.(int)) //打印出value来！不用写成m[i]这种东西，v是个接口类型
			}
		}(i) //这个i是外面的i，go func(i int)的i 是调用的这个i， i == 外部的i
	}
	time.Sleep(1 * time.Second) //主routine等10秒
}
