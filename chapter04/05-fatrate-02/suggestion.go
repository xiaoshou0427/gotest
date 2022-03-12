package main

//某个东西需要按照特定的规则才能工作的话，需要有自己的一个初始化，需要符合我们的需求
//三维数组[][][] ===> [男][0][0，1，2，3，4],[男][1][0,1,2,3,4],2{0,1,2,3,4}},女{0,1,2} }

func getFatRateSuggestion() *fatRateSuggestion {
	return &fatRateSuggestion{
		suggArr: [][][]int{
			//第一个元素表示：男 [] 固定为0
			{
				//3个年龄段 [男][]
				{ //年龄18 - 39 [男][0] 这个年龄段固定为0， 如果从这一层取[][] 取的是下面的整体数组
					//这里0代表偏瘦，剩下的代表其他体型，一共45个逗号，每一个代表这个年龄段，不同体脂对应的体型（年龄相同，体脂不同，体型有重叠）
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
				},
				{ //年龄40 - 59 [男][1] 这个年龄段固定为1

				},
				{ //年龄60+ [男][2] 这个年龄段固定为2

				},
			},
			//第二个元素表示：女 [] 固定为1
			{
				//3个年龄段 [女][]
				{ //年龄18 - 39

				},
				{ //年龄40 - 59

				},
				{ //年龄60+

				},
			},
		},
	}
}
//这里就是把期望的数据写进结构体！！！！ 往结构体的第三维写了数据！

//这个值不能随便来放，需要专门的函数来放，看上面
type fatRateSuggestion struct {
	suggArr [][][]int // 这里性别，年龄，体脂率，都用int，体脂率*100即可？
}

func (s *fatRateSuggestion) GetSuggestion(person *Person) string {
	sexIdx := s.getIndexOfSex(person.sex) //获得sex的索引，这里拿到结构体的成员函数getIndexOfSex,性别从person拿，这里person 是结构体，不是指针
	ageIdx := s.getIndexOfAge(person.age) //获得age的索引
	//获取建议的index，这里要思考一下，你要获取的是最终的0，1，2，3，4 才能转换成体型信息
	suggIdx := s.suggArr[sexIdx][ageIdx][int(person.fatRate*100)] //这里是整个的suggestion id
	//前面两个取来 男/女--》 0/1，取了年龄 --》 0，1，2
	//最后一个维度从person的体脂率*100取整数，比如计算的体脂率为0.05xxx ---> 0.05xxx *100 取整数 = 5
	//那么index 为5 ， 比如前两个是[0][0]，最后一个是5 即：s.suggArr[0][0][5] = 0
	//带入给翻译器：
	return s.translateResult(suggIdx)
	// 直接给出了结果
}
//这里重点是什么呢？ 是第三维度的计算，体脂率计算是从0.00 --- > 0.45xxxx  体脂率0% - %45 这个中间有46个体脂率
//取整之后 是0 到45 ---> 对应第三维的index 值！将体脂率的结果取整，转换成了数组的下标！！！
//根据这个下标获取：回写到对象的第三维数据---> 再根据结构体的成员函数获得体型的信息，并返回（根据体脂率获得体型信息！）
//这是一个数学问题！！！！！！！！ 秒啊！

//成员函数，可以往fatRateSuggestion 里面写入，用指针

//第一维数组的index，性别维度
func (s *fatRateSuggestion) getIndexOfSex(sex string) int {
	if sex == "男" {
		return 0
	}
	return 1
}

//男性和女性的年龄取值都一样，所以用相同的一个成员函数就可以了
// 第二维数组的index，年龄维度
func (s *fatRateSuggestion) getIndexOfAge(age int) int {
	switch {
	case age >= 18 && age <= 39:
		return 0 //第二个index
	case age >= 40 && age <= 59:
		return 1
	case age >= 60:
		return 2
	default:
		return -1 //给一个-1 ，如果不在范围内的
	}
}

//第三维数组的index，体脂率，不用if/else 太长了！
//根据体脂率来确认你的体型，那么体脂率和体型是映射关系！这里可以用map！
//做一个翻译的功能，但凡索引为0 1 2 3 4 就返回不同的体型信息

func (s *fatRateSuggestion) translateResult(idx int) string {
	switch idx {
	case 0:
		return "偏瘦"
	case 1:
		return "标准"
	case 2:
		return "偏重"
	case 3:
		return "肥胖"
	case 4:
		return "严重肥胖"
	default:
		return "未知" //这里加default 或者外面加个return 也可以！视频是外面加的return
	}
}
