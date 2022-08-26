package main

import "fmt"

type LinkNode struct {
	data int
	next *LinkNode //这个类型还能这么填？ 自己填自己，下面在实例化的时候，next 可以是下一个实例
}

func main() {
	n1 := &LinkNode{ //这里用的是指针，nx 都是指针变量
		data: 1,
		next: nil,
	}
	n2 := &LinkNode{
		data: 2,
		next: nil,
	}
	n3 := &LinkNode{
		data: 3,
		next: nil,
	}
	n4 := &LinkNode{
		data: 4,
		next: nil,
	}
	n6 := &LinkNode{
		data: 6,
		next: nil,
	}

	n1.next = n2 //这里指向指针了
	n2.next = n3
	n3.next = n4
	n4.next = n6

	//循环， 做个判断，next 不为空，打印data
	{

	}

	{
		fmt.Println("插入5")
		n5 := &LinkNode{
			data: 5,
			next: nil,
		}

		insertNode(n1, n5)
		//先插入，后调用循环查看
		//还可以继续insert
		insertNode(n1, &LinkNode{
			data: 7,
			next: nil,
		})
		rangeLink(n1)
	}

	{
		fmt.Println("删除节点")
		n1 = deleteNode(n1, 3) //有返回值，要接住它！
		n1 = deleteNode(n1, 5)
		n1 = deleteNode(n1, 1)
		rangeLink(n1)

	}

}

func deleteNode(root *LinkNode, data int) *LinkNode {
	tmpNode := root

	if root != nil && root.data == data { //假设你从n1开始，又刚好需要删除n1
		if root.next == nil {
			return nil
		}
		right := root.next //取了n2
		tmpNode.next = nil //将n1 断开，也就是n1的next 为nil
		return right       //返回n2 把n1清除掉
	}

	for {
		if tmpNode.next == nil { //遇到nil退出循环，找到头的位置
			break
		}
		right := tmpNode.next
		if right.data == data {
			//找到要删除的节点，开始删除
			tmpNode.next = right.next // 比如5 ，此时4 指向 6，不再指向5
			right.next = nil          //清空5 ，不再与6关联，GC回收
			return root               //返回root，自己回去想吧，为啥这里返回root！
		}
		//如果不相等，继续往下找
		tmpNode = tmpNode.next

	}
	return root
}

func insertNode(root *LinkNode, newNode *LinkNode) {
	tmpNode := root
	for {
		if tmpNode != nil {
			//fmt.Println(tmpNode.data)
			if newNode.data > tmpNode.data { //对比现在的data，大于这个data往下走
				if tmpNode.next == nil { //这里按照上面的节奏，不会到达这里，6才会next为nil，其他没有
					//已经到结尾，直接追加
					tmpNode.next = newNode //比6大 就走这里
				} else {
					if tmpNode.next.data >= newNode.data { //到4了，6 比 5 大
						//找到合适位置，准备插入数据。
						newNode.next = tmpNode.next //这里n5 指向n6， tmpNode 此时为4，next指向6
						tmpNode.next = newNode      //把4 的next 改为5 就行了
						break
					}
				}
			}
		} else {
			break
		}
		//还要移位,把next赋值给tmpNode
		tmpNode = tmpNode.next
	}
}

func rangeLink(root *LinkNode) {
	tmpNode := root
	for {
		if tmpNode != nil {
			fmt.Println(tmpNode.data)
		} else {
			break
		}
		//还要移位,把next赋值给tmpNode
		tmpNode = tmpNode.next
	}
}

//每次都是从1开始找的！
