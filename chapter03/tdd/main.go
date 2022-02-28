package main

import "sort"

//定义变量 个人的信息：名字-体脂率
var (
	personFatRate = map[string]float64{}
)

//完成录入信息，调用上面的map，通过录入人名和体脂率，写入到人员信息的map中
func inputRecord(name string, fatRate float64) {
	personFatRate[name] = fatRate //通过录入信息，写入map，只要调用这个函数，每次都写入到这个map
}

//继续来满足案例中王强和李静的需求，先从王强开始，再到李静，此时需要一个排名，那么排名需要准备什么呢？
//需要一个数组，从1，2，3，4 中去找到排名，你需要对排名进行排序
func getRank(name string) (rank int, fatRate float64) {
	fatRate2PersonMap := map[float64]string{}         //做一个临时的map，将key变成体脂率，value变成人名，方便做排序
	rankArr := make([]float64, 0, len(personFatRate)) //定义一个数组，长度为0，容量为map的容量，给了一个容量的预期，性能会好很多，先准备好这么多空间
	for nameItem, frItem := range personFatRate {     // 循环录入信息
		fatRate2PersonMap[frItem] = nameItem //进行反转，将personFatRate的key和value进行反转,放到新的map里
		rankArr = append(rankArr, frItem)    //获取体脂率，把体脂率塞到数组中，下一步就是排序
	}
	//引入sort 这个包，将数组直接排序了，不会有什么返回，直接将rankArr 变成从小到大排序了
	sort.Float64s(rankArr)
	for i, frItem := range rankArr { //这里for后面的变量都是局部变量，与上一个for中的frItem不是同一个，所以IDE是蓝绿色
		_name := fatRate2PersonMap[frItem] //临时变量_name:获取体脂率对应的人名，这是排序后遍历数组，通过体脂率key---去找到---对应的人名value，赋值后再做判断
		if _name == name {                 //如果排序后的名字和录入的名字（形参）一样）返回一个排名，注意此时已经排序，上面map写入多少次，都是排好序的
			rank = i + 1 //数组的下标 i 默认从0开始，所以+1 ，这个是当前排名，当录入王强后，又录入李静后，录入李静比对名字，符合就是i+1 为1，返回体脂率，如果不符合呢？增加其他判断
			fatRate = frItem //给fatRate赋值
			return //此时就可以return了，其实这里不写return 也是可以的，直接用下面的return
		}
	}
	return //别往里这个return
}
