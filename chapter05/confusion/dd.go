package main

type downloadFromDisk struct {
	secret   DynamicSecret //换成一个接口类型
	filePath string
}

//下载文件
func (dd *downloadFromDisk) DownloadFile() {
	if err := dd.loginChcek(); err != nil { //判断是否登陆验证成功
		//todo 重新登陆
	}
	//如果登陆成功，就去阿里下载文件
	dd.DownloadFromAliYun(dd.filePath)
}

//登陆验证,与阿里云的服务进行交互，check账号密码！
func (dd *downloadFromDisk) loginChcek() error {
	dd.checkSecret(dd.secret.GetSecret())
	return nil
}

//从远端下载
func (dd *downloadFromDisk) DownloadFromAliYun(file string) {
	//todo like wget/curl the URL
}

//拿到密码，去远端检查密码
func (dd *downloadFromDisk) checkSecret(secret string) {
	//todo 调用阿里云的验证接口去验证密码是否有效
}
