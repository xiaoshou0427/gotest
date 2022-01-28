package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	defer func() {
		finishTime := time.Now()
		fmt.Println(startTime, finishTime)
		fmt.Println(finishTime.Sub(startTime))
	}()
	arr2 := testPanic()

	arr3 := make([]string, 3, 4)
	copy(arr3, arr2)
	println(arr3[0])

}

func testPanic() []string {
	arr2 := make([]string, 0, 4)
	fmt.Printf("arr2---》len: %d, cap: %d\n", len(arr2), cap(arr2))
	arr2[0] = "1"
	fmt.Println("arr2-1", arr2[0])
	fmt.Println("arr2-2", arr2[1])
	fmt.Println("arr2-3", arr2[2])
	return arr2
}
