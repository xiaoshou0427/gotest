package main

import "fmt"

func main() {
	a := []int{111, 222, 333, 444, 555, 666}
	fmt.Println(a)
	fmt.Printf("打印a的内存：%p\n", &a)
	fmt.Printf("打印a[2]的内存：%p\n", &a[2])
	//在222后面插入101010
	//BackUp := append([]int{},a[2:]...) //这里用一个新的slice 放入333，444，555，666，这样就不会指向之前的内存
	BackUp := []int{1, 2, 3, 4, 5} //这里很有意思，如果写[]int{} 会报错，默认是0，必须要写具体的元素个数
	copy(BackUp, a[2:])            //这里是把后面四个元素复制给了BackUp
	fmt.Println(BackUp, a[2:])
	fmt.Printf("打印BackUp的内存：%p\n", &BackUp)
	fmt.Printf("打印BackUp[0]的内存：%p\n", &BackUp[0])
	fmt.Println(BackUp[0])

	a = append(a[:2], 101010) //这个操作是在元素222后面追加101010，原本的值就没有了
	fmt.Println(a)
	fmt.Printf("打印a[2]的内存：%p\n", &a[2])
	fmt.Println(a[2])
	fmt.Printf("打印BackUp[0]的内存：%p\n", &BackUp[0])
	fmt.Println(BackUp[0])

	a = append(a, BackUp...) //一定要写个点点点
	fmt.Println(a)
	fmt.Printf("打印a[2]的内存：%p\n打印a[3]的内存：%p\n", &a[2], &a[3])
	fmt.Println("a[2]:", a[2], "a[3]:", a[3]) //把Backup[0]的值放到了a[3]中

}
