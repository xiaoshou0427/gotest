package main

import "fmt"

//节点信息
type Node struct {
	data  int
	root  *Node //增加上联根节点的信息
	left  *Node
	right *Node
}

//创建树
func buildTree() *Node {
	n1 := &Node{data: 51}
	n2 := &Node{data: 35}
	n3 := &Node{data: 65}
	//树的关联关系
	n1.left = n2
	n2.root = n1
	n1.right = n3
	n3.root = n1
	return n1
}

//插入数据
func insertNode(root *Node, newNode *Node) *Node {
	//如果为空，直接返回新节点
	if root == nil {
		return newNode
	}
	//如果插入的数据是存在的，也就返回root了
	if newNode.data == root.data {
		return root
	}
	if newNode.data < root.data {
		if root.left == nil {
			// 找到位置，插入数据
			root.left = newNode //插入数据
			newNode.root = root //定义上联root信息
		} else {
			insertNode(root.left, newNode) //往左边移动
		}
	} else {
		if root.right == nil {
			// 找到位置，插入数据
			root.right = newNode //插入到右边的数据
			newNode.root = root  //指回root
		} else {
			insertNode(root.right, newNode) //往右边移动
		}
	}
	return root
}

//这里是删除叶子节点的方法！ 只是给你个感官的认识
func deleteNodeLeaf(root *Node, v int) *Node {
	leftRoot := root //把root 定义出来
	//这里是判断叶子节点，所以判断条件是3个，才能符合叶子节点的特性
	if leftRoot.data == v && leftRoot.left == nil && leftRoot.right == nil { //当root的data 等于 要删除的节点时，也就是现在root指向了要删除的节点
		leftRoot = leftRoot.root //这里只是把leftroot重新赋值，把要切断的节点的root定义回去
		//上面是把root这条线断开
		right := root
		//删的时候需要判断一下是删的左边还是右边
		if leftRoot.left == right { //这里可以参考28，来理解，right 是28，leftroot 已经是上游的35了
			//那么35 的左边是等于28 的，那么删除28如下操作：
			//删除左边叶子
			leftRoot.left = nil //清除35的左边
			right.root = nil    //28的上线也清除了
			return leftRoot     //这里返回的不是root！
		} else {
			//删除右边叶子 （以43为例或者65为例）
			leftRoot.right = nil
			right.root = nil
			return leftRoot //这里并不是return root，因为root 此时已经是要删除的节点了
		} //整个删除是有技巧的
	}
	return root
}

//删除节点递归过程

func deleteNode(root *Node, v int) *Node {
	if v < root.data { //找到删除的节点小于根节点，看看谁更小
		deleteNode(root.left, v) //递归去找根节点的左边，要删除的还小于左边的节点，那就继续递归
	} else if v > root.data { //看看谁更大
		deleteNode(root.right, v) //拿右边的节点继续递归查询
	} else { //不是大于 不是小于，那就是等于！！
		//现在root指向要删除的节点，开始找后继
		leftNextGen := findNextGenFromLeft(root.left)    //左边找到最大的后继
		rightNextGen := findNextGenFromRight(root.right) //右边找到最小的后继
		//这里最简单的情况是删除65，那么65的左边和右边都为空
		if leftNextGen == nil && rightNextGen == nil {
			// 现在要删除的是叶子节点，即最底部节点！！两边断掉就行了
			top := root.root      //定义一个top ： 65的上面是51
			down := root          //定义一个down ： 65
			if top.left == down { //51的左边是35，35如果等于down 也就是65 就执行下面的
				//表示是左子树
				top.left = nil
				down.root = nil
				return nil //root已经删除
			} else {
				//表示是右子树
				top.right = nil //砍断右边
				down.root = nil //砍断上联 （这里叫上联，不叫左边）
				return nil
			}
		} else if leftNextGen != nil { //如果左边的最大的后继不为空
			root.data = leftNextGen.data              //把它提到root节点，例如43
			deleteNode(leftNextGen, leftNextGen.data) //继续递归，把原来位置上的节点删除了？
		} else if rightNextGen != nil {
			root.data = rightNextGen.data
			deleteNode(rightNextGen, rightNextGen.data) //整个过程不断的迭代
		}
	}
	return root
}

//找左边的后继，以删除51为例，从35往下找？
//左边找到最大的，最大的一定在节点的右侧
func findNextGenFromLeft(root *Node) *Node {
	if root == nil {
		return nil
	}
	tmpNode := root
	for {
		if tmpNode.right != nil { //节点的右侧一定是最大的
			tmpNode = tmpNode.right //直到找到左边的后继
		} else {
			break
		}
	}
	return tmpNode
}

//找右边的后继，以删除51为例，右边后继找最小的
func findNextGenFromRight(root *Node) *Node {
	if root == nil {
		return nil
	}
	tmpNode := root
	for {
		if tmpNode.left != nil { //找到右边最小的，因为右半边最小的一定是在这个节点的左边出现！！
			tmpNode = tmpNode.left //直到找到后继
		} else {
			break
		}
	}
	return tmpNode
}

func main() {
	root := buildTree()
	insertNode(root, &Node{data: 43})
	insertNode(root, &Node{data: 28})
	fmt.Println("Done")
	fmt.Println("删除Node")
	deleteNode(root, 43)
	deleteNode(root, 28)
	fmt.Println("删除结束")
}
