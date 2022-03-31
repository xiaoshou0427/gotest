package main

import "testing"

func Test_fatRateSuggestion_GetSuggestion(t *testing.T) {
	//利用函数进行初始化，这个动作一定要有，写回到三维数组
	sugg := getFatRateSuggestion()
	//测试的时候要测试人，那就要给个人的信息进来，测试的时候只关注功能，不是所有的信息都需要
	//这里只需要sex，age，fatRate
	//这里用不用指针都行，给这个对象写个值也可以，这里为啥要用slice 因为下面for 循环range 了！
	tests := []Person{
		{
			sex:     "男",
			age:     35,
			fatRate: 0.24, //0.24体型是肥胖
		},
	}
	//&tests[0] 你细品，sugg是个函数，而这个函数返回的是结构体，而GetSuggestion是结构体的成员函数
	//这个成员函数要求形参的类型是Person的指针类型
	//定义变量，判断变量如果不等于肥胖，就报错
	if got := sugg.GetSuggestion(&tests[0]); got != "肥胖" {
		t.Fail()
	}
}

func Test_fatRateSuggestion_GetSuggestion1(t *testing.T) {
	sugg := getFatRateSuggestion()
	type args struct {
		person *Person //结构体内的成员变量指向 Person 这个结构体的指针
	}
	tests := []struct { //这是定义一个变量为一个slice的结构体
		name string //这是成员变量，一个元素的成员变量
		//fields fields 不用这个参数
		args args //这是一个结构体
		want string
	}{
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.0}}, want: "偏瘦"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.01}}, want: "偏瘦"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.02}}, want: "偏瘦"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.03}}, want: "偏瘦"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.04}}, want: "偏瘦"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.05}}, want: "偏瘦"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.06}}, want: "偏瘦"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.07}}, want: "偏瘦"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.08}}, want: "偏瘦"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.09}}, want: "偏瘦"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.10}}, want: "标准"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.11}}, want: "标准"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.12}}, want: "标准"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.13}}, want: "标准"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.14}}, want: "标准"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.15}}, want: "标准"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.16}}, want: "偏重"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.17}}, want: "偏重"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.18}}, want: "偏重"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.19}}, want: "偏重"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.20}}, want: "偏重"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.21}}, want: "肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.22}}, want: "肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.23}}, want: "肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.24}}, want: "肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.25}}, want: "肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.26}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.27}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.28}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.29}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.30}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.31}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.32}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.33}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.34}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.35}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.36}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.37}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.38}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.39}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.40}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.41}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.42}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.43}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.44}}, want: "严重肥胖"},
		{name: "35", args: args{person: &Person{sex: "男", age: 35, fatRate: 0.64}}, want: "严重肥胖"},
		{
			name: "0.24-35-男",
			args: args{ //这个必须这么写
				person: &Person{
					sex:     "男",
					age:     35,
					fatRate: 0.24, //0.24体型是肥胖,这个fatRate是预期结果
				},
			},
			want: "肥胖",
		},
	}
	//这一段测试数据和下面的这段 都没计算体脂率，上面是给了体脂率0.34 --那么就调用GetSuggestion
	// 传入已经定义的0.34 ,转换成第三维的下标：frIdx := int(person.fatRate*100) ,就找到了[][][34]
	//就拿到了原代码中的体脂建议，对比ut提供的体脂建议！

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := sugg //
			if got := s.GetSuggestion(tt.args.person); got != tt.want {
				t.Errorf("GetSuggestion() = %v, want %v", got, tt.want)
			}
		})
	}
}
