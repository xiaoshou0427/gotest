package main

import (
	"fmt"
	"time"
)

type WebServerV2 struct { //webserver 有配置和读写锁，因为访问量很大的时候需要用读写锁
	config *Config //嵌套了上面的内容，把config变成原子的操作，这里把Config改成指针类型
}

//实现一个reload功能， 这里是写入到指针的内容？
func (ws *WebServerV2) reload() {
	ws.config = &Config{
		Content: fmt.Sprintf("%d", time.Now().UnixNano()),
	}
}

//这里需要周期性的读取配置，监听配置！
func (ws *WebServerV2) ReloadWorker() {
	for { //这里for循环一直做
		time.Sleep(10 * time.Microsecond) //模拟少一点 10毫秒
		ws.reload()                       //每10毫秒去写一下webserver的内容！也是调用对象的reload 成员函数！
	}
}

//每隔10秒重新写入Content！ 这个赋值操作是安全的！

//需要访问webserver，那么webserver就需要有一些功能
func (ws *WebServerV2) Visit() string {
	return ws.config.Content //这里读的是之前的数据
}
