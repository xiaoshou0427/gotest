package main

import (
	"fmt"
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	m := map[int]int{}      //定义一个map
	for i := 0; i < 100; i++ { //来100个worker，一起运行
		go func() {
			for {
				v := m[i]  //获得每个key的value
				m[i] = v + 1 //每个value 都加1，即map的每个value都加1
				fmt.Println("i=", m[i])
			}
		}()
	}
	time.Sleep(10*time.Second) //100 worker没做完，你退出来，不好吧！
}
