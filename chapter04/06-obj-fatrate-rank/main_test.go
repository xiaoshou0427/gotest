package main

import "testing"

func TestCast1Part1(t *testing.T) {
	r:=&FatRateRank{}
	r.inputRecord("王强",0.38) //录入第一次数据，体脂率是个小数
	r.inputRecord("王强",0.32) //录入第二次数据，此时王强的最佳体脂为32
	//上述是录入信息，那你要对比信息，所以要get到王强的排名
	{
		rankOfWQ, fatRateOfWQ := r.getRank("王强") //获得王强的排名,同时获取体脂率，下面是做判断
		if rankOfWQ != 1 {                       //如果王强的排名不是第一名，就算失败,希望在不符合预期的时候尽快结束代码--fail fast
			t.Fatalf("预期王强第一，但是得到的是： %d", rankOfWQ) //利用Fatalf 失败了，就停止
		}
		if fatRateOfWQ != 0.32 { //如果获得的体脂率与期望录入的值不相等，那么停止测试 抛出错误
			t.Fatalf("预期王强的体脂率为0.32，但是得到的是：%f", fatRateOfWQ)
		}
	}
}
func TestCase1(t *testing.T) {
	r:=&FatRateRank{}
	r.inputRecord("王强",0.38) //录入第一次数据，体脂率是个小数
	r.inputRecord("王强",0.32) //录入第二次数据，此时王强的最佳体脂为32
	//上述是录入信息，那你要对比信息，所以要get到王强的排名
	{
		rankOfWQ, fatRateOfWQ := r.getRank("王强") //获得王强的排名,同时获取体脂率，下面是做判断
		if rankOfWQ != 1 {                       //如果王强的排名不是第一名，就算失败,希望在不符合预期的时候尽快结束代码--fail fast
			t.Fatalf("预期王强第一，但是得到的是： %d", rankOfWQ) //利用Fatalf 失败了，就停止
		}
		if fatRateOfWQ != 0.32 { //如果获得的体脂率与期望录入的值不相等，那么停止测试 抛出错误
			t.Fatalf("预期王强的体脂率为0.32，但是得到的是：%f", fatRateOfWQ)
		}
	}
	//录入李静的信息
	r.inputRecord("李静",0.28)  //此时李静已经排名第一了，王强第二
	{ //当录入李静之后，要对比王强的，因为此时王强的排名应该是第二名，且体脂率是0.32
		rankOfWQ, fatRateOfWQ := r.getRank("王强") //获得王强的排名,同时获取体脂率，下面是做判断
		if rankOfWQ != 2 {                       //如果王强的排名不是第一名，就算失败
			t.Fatalf("预期王强第二，但是得到的是： %d", rankOfWQ) //利用Fatalf 失败了，就停止
		}
		if fatRateOfWQ != 0.32 { //如果获得的体脂率与期望录入的值不相等，那么停止测试 抛出错误
			t.Fatalf("预期王强的体脂率为0.32，但是得到的是：%f", fatRateOfWQ)
		}
	}
	{ //此时也要判断李静的状态,getRank 是不会带上WQ/LJ 的！！
		rankOfLJ, fatRateOfLJ := r.getRank("李静") //获得李静的排名,同时获取体脂率，下面是做判断
		if rankOfLJ != 1 {                     //如果李静的排名不是第一名，就算失败
			t.Fatalf("预期李静第一，但是得到的是： %d", rankOfLJ) //利用Fatalf 失败了，就停止
		}
		if fatRateOfLJ != 0.28 { //如果获得的体脂率与期望录入的值不相等，那么停止测试 抛出错误
			t.Fatalf("预期李静的体脂率为0.28，但是得到的是：%f", fatRateOfLJ)
		}
	}

}

func TestCase2(t *testing.T) {
	r:=&FatRateRank{}
	r.inputRecord("王强", 0.38)
	r.inputRecord("张伟", 0.38)
	r.inputRecord("李静", 0.28)
	{ //李静的信息
		rankOfLJ, fatRateOfLJ := r.getRank("李静")
		if rankOfLJ != 1 {
			t.Fatalf("预期李静第一，但是得到的是： %d", rankOfLJ) //利用Fatalf 失败了，就停止
		}
		if fatRateOfLJ != 0.28 { //如果获得的体脂率与期望录入的值不相等，那么停止测试 抛出错误
			t.Fatalf("预期李静的体脂率为0.28，但是得到的是：%f", fatRateOfLJ)
		}
	}
	{
		rankOfWQ, fatRateOfWQ := r.getRank("王强") //获得王强的排名,同时获取体脂率，下面是做判断
		if rankOfWQ != 2 {                     //如果王强的排名不是第一名，就算失败,希望在不符合预期的时候尽快结束代码--fail fast
			t.Fatalf("预期王强第二，但是得到的是： %d", rankOfWQ) //利用Fatalf 失败了，就停止
		}
		if fatRateOfWQ != 0.38 { //如果获得的体脂率与期望录入的值不相等，那么停止测试 抛出错误
			t.Fatalf("预期王强的体脂率为0.38，但是得到的是：%f", fatRateOfWQ)
		}
	}
	{
		rankOfZW, fatRateOfZW := r.getRank("张伟") //获得王强的排名,同时获取体脂率，下面是做判断
		if rankOfZW != 2 {                     //如果王强的排名不是第一名，就算失败,希望在不符合预期的时候尽快结束代码--fail fast
			t.Fatalf("预期张伟第二，但是得到的是： %d", rankOfZW) //利用Fatalf 失败了，就停止
		}
		if fatRateOfZW != 0.38 { //如果获得的体脂率与期望录入的值不相等，那么停止测试 抛出错误
			t.Fatalf("预期张伟的体脂率为0.38，但是得到的是：%f", fatRateOfZW)
		}
	}

}

func TestCase3(t *testing.T) {
	r:=&FatRateRank{}
	r.inputRecord("王强", 0.38)
	r.inputRecord("李静", 0.28)
	r.inputRecord("张伟") //需求就是张伟不录入体脂率！
	{ //李静的信息
		rankOfLJ, fatRateOfLJ := r.getRank("李静")
		if rankOfLJ != 1 {
			t.Fatalf("预期李静第一，但是得到的是： %d", rankOfLJ) //利用Fatalf 失败了，就停止
		}
		if fatRateOfLJ != 0.28 { //如果获得的体脂率与期望录入的值不相等，那么停止测试 抛出错误
			t.Fatalf("预期李静的体脂率为0.28，但是得到的是：%f", fatRateOfLJ)
		}
	}
	{
		rankOfWQ, fatRateOfWQ := r.getRank("王强") //获得王强的排名,同时获取体脂率，下面是做判断
		if rankOfWQ != 2 {                     //如果王强的排名不是第一名，就算失败,希望在不符合预期的时候尽快结束代码--fail fast
			t.Fatalf("预期王强第二，但是得到的是： %d", rankOfWQ) //利用Fatalf 失败了，就停止
		}
		if fatRateOfWQ != 0.38 { //如果获得的体脂率与期望录入的值不相等，那么停止测试 抛出错误
			t.Fatalf("预期王强的体脂率为0.38，但是得到的是：%f", fatRateOfWQ)
		}
	}
	{
		rankOfZW, _ := r.getRank("张伟") //获得王强的排名,不需要关注体脂率，下面是做判断
		if rankOfZW != 3 {                     //如果王强的排名不是第一名，就算失败,希望在不符合预期的时候尽快结束代码--fail fast
			t.Fatalf("预期张伟第三，但是得到的是： %d", rankOfZW) //利用Fatalf 失败了，就停止
		}
	}

}