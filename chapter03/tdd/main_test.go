package main

import "testing"

func TestCase1(t *testing.T) {
	inputRecord("王强",0.38) //录入第一次数据，体脂率是个小数
	inputRecord("王强",0.32) //录入第二次数据，此时王强的最佳体脂为32
	//上述是录入信息，那你要对比信息，所以要get到王强的排名
	{
		randOfWQ, fatRateOfWQ := getRand("王强") //获得王强的排名,同时获取体脂率，下面是做判断
		if randOfWQ != 1 {                       //如果王强的排名不是第一名，就算失败,希望在不符合预期的时候尽快结束代码--fail fast
			t.Fatalf("预期王强第一，但是得到的是： %d", randOfWQ) //利用Fatalf 失败了，就停止
		}
		if fatRateOfWQ != 0.32 { //如果获得的体脂率与期望录入的值不相等，那么停止测试 抛出错误
			t.Fatalf("预期王强的体脂率为0.32，但是得到的是：%f", fatRateOfWQ)
		}
	}
	//录入李静的信息
	inputRecord("李静",0.28)  //此时李静已经排名第一了，王强第二
	{ //当录入李静之后，要对比王强的，因为此时王强的排名应该是第二名，且体脂率是0.32
		randOfWQ, fatRateOfWQ := getRand("王强") //获得王强的排名,同时获取体脂率，下面是做判断
		if randOfWQ != 2 {                       //如果王强的排名不是第一名，就算失败
			t.Fatalf("预期王强第一，但是得到的是： %d", randOfWQ) //利用Fatalf 失败了，就停止
		}
		if fatRateOfWQ != 0.32 { //如果获得的体脂率与期望录入的值不相等，那么停止测试 抛出错误
			t.Fatalf("预期王强的体脂率为0.32，但是得到的是：%f", fatRateOfWQ)
		}
	}
	{ //此时也要判断李静的状态,getRand 是不会带上WQ/LJ 的！！
		randOfLJ, fatRateOfLJ := getRand("王强") //获得李静的排名,同时获取体脂率，下面是做判断
		if randOfLJ != 1 {                     //如果李静的排名不是第一名，就算失败
			t.Fatalf("预期王强第一，但是得到的是： %d", randOfLJ) //利用Fatalf 失败了，就停止
		}
		if fatRateOfLJ != 0.28 { //如果获得的体脂率与期望录入的值不相等，那么停止测试 抛出错误
			t.Fatalf("预期王强的体脂率为0.28，但是得到的是：%f", fatRateOfLJ)
		}
	}

}
