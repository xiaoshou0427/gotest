package main

//有一个清单

type Assets struct {
	assets []Asset //是一堆的资产
}

//保安根据清单,上班的时候开门，下班的时候关门，start和stop操作！前提如果是门，才做这个动作

func (a *Assets) DoStartWork() {
	for _, item := range a.assets {
		if d, ok := item.(Door); ok { //判断接口类型，判断item的类型是不是Door，是不是个门，这里Door是个类型
			d.Unlock()
			d.Open() //如果它是个门，开启这个门，d为Door类型
		}
	}
}

func (a *Assets) DoStopWork() {
	for _, item := range a.assets {
		if d, ok := item.(Door); ok { //判断类型，判断item是不是Door这个类型，是不是个门
			d.Close() //如果它是个门，关闭这个门
			d.Lock()
		}
	}
}

//这里关于断言，我是这么理解的：就是item能不能实现Door这个接口，item是个特殊范围（更小范围），而Door是个通用范围
// item 是不是可以实现Door这个接口！显然现在是可以实现的，下面增加锁的例子就是不能实现的！
//这里是判断ok是否为true，还可以直接强转，就是不判断ok，直接变量:= xx.(aaa)??是这么理解吗？
