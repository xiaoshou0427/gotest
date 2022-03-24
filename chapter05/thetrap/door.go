package main

//门：有开和关
type Door interface {
	Unlock() //解锁
	Open()
	Close()
	Lock() //加锁
}