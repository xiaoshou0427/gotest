package main

import (
	"fmt"
)

//type CompareFunc func(left,right interface{}) bool //interface 用这个对比？

type LinkNode struct {
	data     int       //简化过程，用int 不用interface
	next     *LinkNode //这个类型还能这么填？ 自己填自己，下面在实例化的时候，next 可以是下一个实例
	previous *LinkNode //双链表增加节点（左一个，next 右一个）
}

//前面的next 指向后一个的previous
//第一个previous 是nil，最后一个next 是nil

//返回root节点
func buildDLink() *LinkNode {
	n1 := &LinkNode{data: 1}
	n2 := &LinkNode{data: 5}
	n3 := &LinkNode{data: 10}
	//如何链接这些节点，如下
	n1.next = n2     //单向的 从左到右
	n2.previous = n1 //单向的 从右到左

	n2.next = n3
	n3.previous = n2

	return n1 //返回n1
}

//实现插入节点
func insertNode(root *LinkNode, newNode *LinkNode) *LinkNode {
	//循环找到新节点要插入的位置
	tmpNode := root
	//这里如果root为空，也就是链表为空，直接返回新节点
	if root == nil {
		return newNode
	}
	//新节点入头部的情况：
	if root.data >= newNode.data {
		newNode.next = tmpNode //链起来就行了，跟尾部类似
		tmpNode.previous = newNode
		return newNode
	}
	for {
		//入尾部的情况：
		if tmpNode.next == nil {
			//已经到头了，追加节点即可，当然是双向链路的
			tmpNode.next = newNode     //最后一个节点的右边指向新节点
			newNode.previous = tmpNode //新节点的左边指向原来最后一个节点
			return root                //每次追加完，返回root就行了
		} else { //不是尾部的情况
			//这里root的next的data 大于 等于新插入数据的data
			if tmpNode.next.data >= newNode.data {
				// 找到位置，在此插入新节点 todo
				newNode.previous = tmpNode  //新节点的左边指向上一个节点
				newNode.next = tmpNode.next //新节点的右边指向下一个节点
				//此时新节点指向了前一个和后一个，原来节点上的线路还没拆除，也没有指向新的节点
				//此时拆除旧关联，关联新节点
				tmpNode.next = newNode //前一个节点的next指向新节点
				//后一个节点的previous 指向新节点! 这里不太好找，你怎么能找到这里呢？
				//只能通过新节点的下线来找，这里新节点的next是下一个节点（newNode.next = tmpNode.next ）这里已经定义了
				//那么它的previous 就可以关联上一个节点，也就是新节点！可以看着图去想一想
				newNode.next.previous = newNode
				//这里不能用tmpNode.next.previous 因为这里tmpNode.next = newNode 已经把旧的链路已经断开了，只能从新的里面关联起来
				//这里一定要注意，最好画个草图会方便很多！
				return root
			}
		}
		//那如果新的数字比前几个都大呢？ 比如7 大于1 和5，上面那个就不中用了
		tmpNode = tmpNode.next
		//如果新节点是7，1的next.data是5，那么5 >= 7 吗？
		//不大于，所有tmpNode 要往后挪一次，此时tmpNode为 tmpNode.next 也就是等于5,再继续循环
	}
}

//删除节点操作,要有返回值，可能也会删除头部
func deleteNode(root *LinkNode, v int) *LinkNode {
	if root == nil {
		return nil
	}
	//删除头部（删除的数据在第一个节点）
	if root.data == v { //当删除的节点出现在第一个节点的时候
		//这里还有一种情况没有判断，那就是只有一个节点的情况，视频让自己做
		//todo 需要解决只有一个节点的情况，我这边判断了一下，如果next为空，就返回空，清除完毕
		if root.next == nil {
			return nil
		}
		//此处用左手这个变量来承接要删除的第一个节点，然后把root移动到第二个节点，这个妙啊！
		leftHand := root //定义一个左手，把root移走，左手是原来的第一个节点，也就是要删除的节点
		root = root.next //重新定义root，把root移走到下一个节点！
		//断开第一个节点
		leftHand.next = nil //原本要删除的节点，也就是第一个节点
		root.previous = nil //root已经是第二个节点了
		//返回root 其实返回的是新的链接
		return root
	}

	//处理删除节点在中间的情况
	tmpNode := root
	for {
		//走到头
		if tmpNode.next == nil { //next为空了,要么没找到，要么就是最后一个了
			return root //没找到要删除的数据，直接返回root
		} else { //判断左边的节点的next的data 是我们需要删除的数据时，做相应的操作
			if tmpNode.next.data == v {
				//找到节点，开始删除，删除完成后返回原root
				// todo
				rightHand := tmpNode.next     //这里如果找到了，那么开始删除，右手等于要删除的数据，找出来
				tmpNode.next = rightHand.next //这就把要删除的节点左侧节点链到删除节点右侧节点

				//如果删除最后一个节点，就需要判断，不然会报错
				//此时的tmpNode.next已经赋值为要删除节点的next 这里如果是最后一个节点就为nil了！
				if tmpNode.next == nil { //这里不用链接右边节点到左边节点，因为是最后一个节点
					//rightHand.next = nil //这条可以不加的，因为最后一个节点就是nil！
					rightHand.previous = nil //删除最后一个节点的向前的链接
				} else {
					//要删除右边节点链到删除节点的线路，右边节点链到左边节点去！跳过要删除的节点
					rightHand.next.previous = tmpNode //就这么干

					//清理需要删除节点的剩余链接，方便GC垃圾回收即可
					rightHand.next = nil
					rightHand.previous = nil
				}
				return root
			}
		}
		//上述没找见，继续往后找
		tmpNode = tmpNode.next

	}
}

//读取链表数据
func rangeLink(root *LinkNode) {
	fmt.Println("从头到尾")
	tmpNode := root

	for {
		fmt.Println(tmpNode.data) //循环读出内容
		if tmpNode.next == nil {  //这一部分如果放在fmt.Println(tmpNode.data) 之前，那么11就打印不出来，我是debug 出来的
			break //到了尾巴就退出了
		}
		tmpNode = tmpNode.next //往后移
	}
	//此时上面的for循环已经移到最后面了
	fmt.Println("从尾到头")
	for {
		fmt.Println(tmpNode.data)
		if tmpNode.previous == nil {
			break
		}
		tmpNode = tmpNode.previous //反过来
	}

}

func main() {
	root := buildDLink() //初始化一个链表
	root = insertNode(root, &LinkNode{data: 3})
	root = insertNode(root, &LinkNode{data: 7})
	root = insertNode(root, &LinkNode{data: 11})
	root = insertNode(root, &LinkNode{data: 0})
	root = insertNode(root, &LinkNode{data: -1})
	rangeLink(root)
	fmt.Println("删除节点")
	root = deleteNode(root, -1)
	root = deleteNode(root, 3)
	root = deleteNode(root, 11)
	rangeLink(root)

}
