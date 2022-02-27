package cacl

import "testing"

func TestCalcBMI(t *testing.T) {
	inputHeight, inputWeight := 1.7, 70.0
	expectedOutput := 24.221453287197235
	t.Logf("开始计算,输入：height：%f，weight：%f", inputHeight, inputWeight)
	actualOutput, err := CalcBMI(inputHeight, inputWeight)
	t.Logf("实际得到：%f，error：%v", actualOutput, err) //error 是什么我实际不知道，给个%v
	if err != nil {
		t.Fatalf("expecting no err, but get: %v", err) //如果有err 就直接退出
	}
	if expectedOutput != actualOutput {
		t.Errorf("expecting %f, but got %v", expectedOutput, actualOutput) //可以用那个errorf 来表示，就不需要f.log + f.errorf
	}
}
