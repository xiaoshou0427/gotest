package main

import "testing"

func TestFrServiceSuggestion(t *testing.T) {
	realOutput := &fakeOutput{} // 存储的值放在这里，用指针，成员函数是写回到成员变量
	frSvc := &fatRateService{
		s: getFatRateSuggestion(),
		input: &fakeInput{},
		output: realOutput, //实例化这个接口
	}
	p:= frSvc.input.GetInput()
	//来个预期
	expOutput := &fakeOutput{
		p: p,
		s: "偏重",
	}
	frSvc.GiveSuggestionToPerson(&p)

	if expOutput.s != realOutput.s {
		t.Fatalf("预期：%s,实际输出：%s", expOutput.s, realOutput.s) //输出文字，退出
	}
}
