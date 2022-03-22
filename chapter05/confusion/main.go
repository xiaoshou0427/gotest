package main

func main() {
	dd := &downloadFromDisk{
		secret: &mobileTokenDynamic{mobileNumber: "13000000111"},
		filePath: "接口变成.ppt",
	}
	dd.DownloadFile()
}

//定义一个接口
type DynamicSecret interface {
	GetSecret() string
}

//做一个具体的实现(接口的实现用下面的对象),发送短信验证，如果它实现了上面接口里面的方法，就可以完成上述案例
type mobileTokenDynamic struct {
	mobileNumber string
}

func (d *mobileTokenDynamic) GetSecret() string {
	return "something"
}
