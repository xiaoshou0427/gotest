package main

import "fmt"

func main() {
	msg := constructHelloMessage()
	fmt.Println(msg)
}

func constructHelloMessage() (test string) {
	test = "oscar"
	return
}
