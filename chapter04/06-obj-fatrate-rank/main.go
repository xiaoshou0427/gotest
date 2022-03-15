package main

import (
	"math"
	"sort"
)

type RankItem struct {
	Name    string
	FatRate float64
}

type FatRateRank struct {
	items []RankItem //这里没用指针，是不希望被篡改数据
}

//这里改用指针，可以省去很多资源
func (r *FatRateRank) inputRecord(name string, fatRate ...float64) { //修改为不定长参数，不定长取出最小值作为最终值的方法如下：
	minFatRate := math.MaxFloat64  //常量 math.MaxFloat64 表示 float64 能取到的最大数值，大约是 1.8e308
	for _, item := range fatRate { //将切片进行循环，其实每次都只有一个值
		if minFatRate > item { //有多个item，有很多值，这些值是item，第一次跟最大值比较，如果小于最大值，就执行下面的
			minFatRate = item //就把item 最大的值赋值给minFatRate，直到循环比较出最小值给到minFatRate !!!! 所以这是最小值 min
		}
	}
	found := false                 //给一个false，往下执行r.items 为空，下面for循环不执行，跳出进入下面if ！found
	for i, item := range r.items { //如果没有值就是空的items，上面定义的是空的列表，下面才往里面写
		if item.Name == name { //相同名字，判断最小的体脂率，写到item里面
			if item.FatRate >= minFatRate { //从上面拿到最小的体脂率，赋值回去
				item.FatRate = minFatRate //注意这里item 是个临时变量！这么赋值回去并不会写入items
			} //比旧变量的值小，就改写旧的变量！为更小的值
			r.items[i] = item //写回到结构体！
			found = true
			break
		}
	}
	if !found { //第一次就跳出这里了，在空列表中进行了追加！
		r.items = append(r.items, RankItem{
			Name:    name,
			FatRate: minFatRate,
		})
	}
}

func (r *FatRateRank) getRank(name string) (rank int, fatRate float64) {
	//排序 slice的排序，自己看解释
	sort.Slice(r.items, func(i, j int) bool {
		return r.items[i].FatRate < r.items[j].FatRate
	})
	frs := map[float64]struct{}{} //map类型的结构体
	//根据体脂率排序后，名字
	for _, item := range r.items { //拿到每个人信息，人名和fatrate，这是个列表
		frs[item.FatRate] = struct{}{} //赋值操作，体脂可能相同，人不一定相同，给一个空的结构体到这个key（体脂率）
		//！！！！精华！！！！这里很微妙，map的key相同，就覆盖了！ 所以比较的时候，就不会再比较相同体脂的，只会跟不同体脂的比较！
		if item.Name == name {
			fatRate = item.FatRate
			break
		}
	}

	rankArr := make([]float64, 0, len(frs)) //定义一个数组，长度为0，容量为frs 的map的容量，给了一个容量的预期，性能会好很多，先准备好这么多空间
	for k := range frs {                    //循环获得key 也就是体脂率! 我只要key！只要体脂率进行排序
		rankArr = append(rankArr, k)
	}
	sort.Float64s(rankArr)
	//再去寻找排行
	for i, frItem := range rankArr { //这里for后面的变量都是局部变量，所以IDE是蓝绿色，这里取出已经排好序的体脂
		if frItem == fatRate { //判断排序后的体脂是否等于录入人员信息的体脂率！如果相同，代表rank 排名
			rank = i + 1 //获得rank
			break
		}
	}

	return
}
