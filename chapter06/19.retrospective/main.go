package main

import (
	"fmt"
	"time"
)

type Downloader struct {
	fileNameCh chan string
	// some kind of download worker
	finishedCh chan string
}

//Achieve Download function
func (d *Downloader) Download() {
	//循环读取channel的文件
	for fileName := range d.fileNameCh {
		fmt.Println("开始下载文件：", fileName)
		time.Sleep(1 * time.Second)
		fmt.Println("开始处理文件： ", fileName)
		time.Sleep(10 * time.Microsecond)
		fmt.Println("保存文件：", fileName)
		d.finishedCh <- fileName //下载的文件保存到channel
	} //这里面是写入到finishedCH，而fileNameCh没有任何写入动作，在main里面去写入的
}

func main() {
	//实例化
	fileNameCh := make(chan string, 50)
	finishedCh := make(chan string)

	workerCounter := 50 //启动10个goroutine
	for i := 0; i < workerCounter; i++ {
		go func() {
			downloader := &Downloader{ //实例化结构体？
				fileNameCh: fileNameCh,
				finishedCh: finishedCh,
			}
			downloader.Download() //执行下载 处理等动作
		}()
	} //上面这部分循环是准备好了worker但是还没有发送信号去下载

	//定义文件数量
	fileNumber := 100
	//创建一个slice 长度为0，容量为100
	fileNames := make([]string, 0, 100)
	for i := 0; i < fileNumber; i++ {
		fileNames = append(fileNames, fmt.Sprintf("file_%d.txt", i))
		//追加100个文件到fileNames
	}

	//扔文件进去，循环读取内容放到channel里面
	for _, fileItem := range fileNames {
		fileNameCh <- fileItem //把文件放进去，循环写入到fileNameCh里面
	}
	close(fileNameCh) //关闭channel，在goroutine里面执行Download的时候range执行完毕就退出了
	//此时文件已经写入到另外一个finishedCh了
	//等待处理完毕，读取放入到finishedCh的文件名
	finishedFileCount := 0
	for finishedFile := range finishedCh {
		fmt.Println("文件：", finishedFile, "处理完毕")
		finishedFileCount++                  //第一个是1
		if finishedFileCount == fileNumber { //判断从1开始到100，那就是100个循环，也就是处理完毕100个文件
			close(finishedCh) //一定要关闭channel
		}
	}
	fmt.Println("所有文件处理完毕，结束")
}
