package main

import (
	"fmt"
	"sync"
	"time"
)

type WebServerV1 struct { //webserver 有配置和读写锁，因为访问量很大的时候需要用读写锁
	config     Config       //嵌套了上面的内容
	configLock sync.RWMutex //多个人都可以读，只有一个人写，如果用普通锁，访问量大，就爆了

}

//实现一个reload功能， 要加锁，写内容，再解锁！这里用的是写锁！
func (ws *WebServerV1) reload() {
	ws.configLock.Lock()
	defer ws.configLock.Unlock()
	ws.config.Content = fmt.Sprintf("%d", time.Now().UnixNano()) //修改配置内容，输出一个纳秒！这里并不直接输出，而是写到内容里！
}

//这里需要周期性的读取配置，监听配置！
func (ws *WebServerV1) ReloadWorker() {
	for { //这里for循环一直做
		time.Sleep(10 * time.Microsecond) //模拟少一点 10毫秒
		ws.reload()                       //每10毫秒去写一下webserver的内容！也是调用对象的reload 成员函数！
	}
}

//需要访问webserver，那么webserver就需要有一些功能
func (ws *WebServerV1) Visit() string {
	ws.configLock.RLock()         //保护，读锁，注意读锁不影响其他的人去读！
	defer ws.configLock.RUnlock() // 最后解锁
	return ws.config.Content      //这里模拟访问的时候返回内容就行了！
}

//上述功能已经完成！
