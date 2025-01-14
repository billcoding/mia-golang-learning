/*
我现在有个需求：打印姓名，年龄，身高，体重
需求里有四个不同的东西需要分开装起来 假设这四个装东西的容器为a,b,c,d
1.声明四个变量 a,b,c ,d,类型分别为字符串string,整型int,整型int,整型int
2.类型分好后，分别给他们赋值，，将姓名赋值给变量a,年龄赋值给b,身高赋值给c,体重赋值给d,
3.导入fmt包
4.调用Println函数打印并换行

问题：
1. 目前：4个变量a b c d，那如果我现在有100个变量呢，难道a,b,c,d,e,f,g...... 试问一下，变量h代表什么？
问题：变量名不规范。解决：规范变量名，如a改成name ,b改成age,c改成height，d改成weight。
2. 单纯打印“Mia 33 158 62”，其实我们不知道 Mia代表什么 33代表什么 158代表什么 62代表什么。
问题：打印内容不明确。解决：就是在打印前增加这个变量的说明。比如：fmt.Println(name) => fmt.Println("名字：", name)
*/
package main

import "fmt"

func main() {
	var name string = "mia"
	var age int = 33
	var height int = 158
	var weight int = 62
	fmt.Println("姓名是：", name, "年龄是：", age, "岁", "身高是：", height, "cm", "体重是：", weight, "kg")
	//fmt.Println("姓名是：", name, "\n", "年龄是：", age, "岁", "\n", "身高是：", height, "cm", "\n", "体重是：", weight, "kg", "\n")
	// 试试调用fmt.Printf函数打印以上内容
	//fmt.Printf("姓名是%s\n,年龄是%d岁\n,身高是%dcm\n,体重是%dkg\n","Mia",33,158,62)
	fmt.Printf("姓名是%s\n年龄是%d岁\n身高是%dcm\n体重是%dkg\n",name,age,height,weight)

}
