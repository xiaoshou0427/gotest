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
	fatRate2PersonMap := map[float64][]string{}  //这里改成嵌套数组，一个key 对应一组值
	rankArr := make([]float64, 0, len(personFatRate)) //定义一个数组，长度为0，容量为map的容量，给了一个容量的预期，性能会好很多，先准备好这么多空间
	for nameItem, frItem := range personFatRate {     // 循环录入信息
		fatRate2PersonMap[frItem] = append(fatRate2PersonMap[frItem],nameItem) //进行反转，这里数组只能用append 不能直接赋值
		rankArr = append(rankArr, frItem)    //获取体脂率，把体脂率塞到数组中，下一步就是排序
	}
	//引入sort 这个包，将数组直接排序了，不会有什么返回，直接将rankArr 变成从小到大排序了
	sort.Float64s(rankArr)
	for i, frItem := range rankArr { //这里for后面的变量都是局部变量，与上一个for中的frItem不是同一个，所以IDE是蓝绿色
		_names := fatRate2PersonMap[frItem] // 这个_names 记录体脂率对应的名字，用数组来解决相同体脂率，不同名字的情况
		for _,_name := range _names { //这里循环数组，把每个名字取出来，对比名字
			if _name == name {                 //
				rank = i + 1 //数组的下标 i 默认从0开始，所以+1 ，这个是当前排名，当录入王强后，又录入李静后，录入李静比对名字，符合就是i+1 为1，返回体脂率，如果不符合呢？增加其他判断
				fatRate = frItem //给fatRate赋值
				return //此时就可以return了，其实这里不写return 也是可以的，直接用下面的return
			}
		}

	}
	return //别忘了这个return
}
