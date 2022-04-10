package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	fmt.Println("开始数")
	var totalCount int64 = 0     //整体字数
	for p := 0; p <= 50; p++ { //5000页循环
		fmt.Println("正在统计", p, "页")
		time.Sleep(1 * time.Second) //先用每数一页用1秒测试
		//time.Sleep(10*time.Second) //每数一页用10秒
		r, _ := rand.Int(rand.Reader, big.NewInt(800)) //未知数读取，假设一页最多有800字
		fmt.Println("有", r.Int64(),"字")
		totalCount += r.Int64() //累加整体的字数
	}
	fmt.Println("总共有 ", totalCount, "字")
}
