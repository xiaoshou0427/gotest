package main

import "testing"

func TestWithDirection(t *testing.T) {
	c := make(chan int, 100)
	inOnly(c)
	outOnly(c)
}

func inOnly(c chan<- int) {
	c <- 1
	//<-c
}

func outOnly(c <-chan int) {
	//c <- 1
	<-c
}
