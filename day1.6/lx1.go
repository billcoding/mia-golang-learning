/*
超满意小学2（5）班秋季运动会比赛项目信息
~~~~~~
参赛地点：足球场，篮球场
参赛项目：立定跳远，男子50米，一分钟仰卧起坐，一分钟跳绳，跳远，跑步等
参赛人员：A B C D E
参赛方法：每班选男女生各5人，每人最多能参加2个项目（共10人）
评判标准：最先到达终点的获胜，用时最短，效率最高的获胜
*/
package main

import "fmt"

func main() {

	{

		var place string = "足球场,篮球场"
		var program string = "立定跳远,男子50米,一分钟仰卧起坐,一分钟跳绳,跳远,跑步等"
		var person string = "A B C D E"
		var way string = "每班选男女生各5人,每人最多能参加2个项目(共10人)"
		var standard = "最先到达终点的获胜，用时最短，效率最高的获胜"

		//  fmt.Printf("超满意小学2(5)班秋季运动会比赛项目信息\n~~~~~~\n")
		//  fmt.Printf("参赛地点：%s\n",place )
		//  fmt.Printf("参赛项目：%s\n",program)
		//  fmt.Printf("参赛人员：%s\n",person)
		//  fmt.Printf("参赛方式：%s\n",way)
		// fmt.Printf("参赛标准：%s\n",standard )

		//fmt.Printf("超满意小学2(5)班秋季运动会比赛项目信息\n~~~~~~\n参赛地点：%s\n参赛项目：%s\n参赛人员：%s\n参赛方式：%s\n参赛标准：%s\n\n",place,program,person,way,standard)

		fmt.Println("超满意小学2(5)班秋季运动会比赛项目信息\n", "~~~~~~\n", "参赛地点：", place, "\n", "赛项目：", program, "\n", "参赛人员:", person, "\n", "参赛方式：", way, "\n", "赛标准：", standard)

	}

	{

		var var1 int = 3
		var var2 int = 9
		fmt.Printf("%d+%d=%d\n", var1, var2, var1+var2)
		var var3 bool = true
		fmt.Println("var3=", var3)

	}

	{
		var numb1 int = 10
		var numb2 int = 99
		fmt.Printf("%d*%d=%d\n", numb1, numb2, numb1*numb2)

		var a bool
		fmt.Println("a=", a)
		var d bool = true
		fmt.Println("d=", d)

	}
	{
		var name string = "love"
		fmt.Println("name:", name)
		fmt.Println(name)

		var numb1 int
		fmt.Println("numb1=", numb1)
		fmt.Println(numb1)

		var n bool
		fmt.Println("n=", n)
		fmt.Println(n)

		var numb3 int = 30
		numb3 += 3         //33
		numb3 -= 4         //29
		numb3++            //30
		numb3--            //29
		numb3 = numb3 - -1 //30
		fmt.Println("numb3=", numb3)
		fmt.Println(numb3)

	}

}
