package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"learn.go/fatrate.refactor/cacl"
	learn_go_tools "learn.go.tools"
)

func main() {
	// 录入
	var (
		name   string
		sex    string
		tall   float64
		weight float64
		age    int
	)

	cmd := cobra.Command{
		Use:   "healthcheck",                            //类似与git push/clone 的push/clone，给这个命令行一个名字
		Short: "体脂计算器，根据提供的信息进行计算,给出建议",                 //短描述
		Long:  "基于BMI的体脂计算器，根据性别，身高，体重，年龄进行体脂计算,给出健康建议", //长（冗余）描述
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("name:", name)
			fmt.Println("sex:", sex)
			fmt.Println("tall:", tall)
			fmt.Println("weight:", weight)
			fmt.Println("age:", age)
			// 计算
			//先算出BMI
			bmi := cacl.CalcBMI(tall, weight)
			fmt.Println("BMI: ", bmi)
			//算出体脂率
			fatRate := cacl.FatRateCalc(bmi, age, sex)
			fmt.Println("FatRate: ", fatRate)
			// 评估结果
			fmt.Println(learn_go_tools.Max(3,5))

		},
	}
	cmd.Flags().StringVar(&name, "name", "", "姓名")
	cmd.Flags().StringVar(&sex, "sex", "", "性别")
	cmd.Flags().Float64Var(&tall, "tall", 0, "身高")
	cmd.Flags().Float64Var(&weight, "weight", 0, "体重")
	cmd.Flags().IntVar(&age, "age", 0, "年龄")

	cmd.Execute()

}
