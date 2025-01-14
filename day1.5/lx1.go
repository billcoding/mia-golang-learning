/*
*
练习1：

学校需要登记小明和小红的学生消息，请按照如下格式打印小明和小红的学生信息。
（要求能使用变量的地方一定要使用变量）

小明的学生信息
---
姓名：小明
性别：男
年龄：12岁
身高：135cm
体重：35kg
家庭地址：湖北武汉

小红的学生信息
---
姓名：小红
性别：女
年龄：11岁
身高：130cm
体重：31kg
家庭地址：湖北仙桃
*/
package main

import "fmt"

func main() {
	fmt.Println("小明的学生信息 ")
	fmt.Println("---")
	{
		var name string = "小明"
		var gender string = "男"
		var age int = 12
		var height int = 135
		var weight int = 35
		var homeAddress string = "湖北武汉"
		fmt.Printf("名字：%s\n", name)
		fmt.Printf("性别：%s\n", gender)
		fmt.Printf("年龄：%d岁\n", age)
		fmt.Printf("身高：%dcm\n", height)
		fmt.Printf("体重：%dkg\n", weight)
		fmt.Printf("家庭住址：%s\n", homeAddress)
	}
	fmt.Println()

	fmt.Println("小红的学生信息 ")
	fmt.Println("---")
	{
		var name = "小红" // 这是声明变量的语法 需要替换为 给这个变量赋值
		var gender string = "女"
		var age int = 11
		var height int = 130
		var weight int = 31
		var homeAddress string = "湖北仙桃"
		fmt.Printf("名字：%s\n", name)
		fmt.Printf("性别：%s\n", gender)
		fmt.Printf("年龄：%d岁\n", age)
		fmt.Printf("身高：%dcm\n", height)
		fmt.Printf("体重：%dkg\n", weight)
		fmt.Printf("家庭住址：%s\n", homeAddress)
	}
}

/*

变量之所以称为变量，就是代表一个变量可以被多次赋值。

声明变量的格式：
var 变量名 变量类型 = 变量值

var num int = 10 // 声明一个名称为num类型为int值为10的变量

- 变量叫什么名字 num
- 变量是什么类型 int
- 变量的值是什么（赋值的过程） ***  num = 10

var clothesColor string = "红色" // 今天我穿红色的衣服，声明一个名称为clothesColor类型为string值为红色的变量。

fmt.Printf("今天我穿%s的衣服\n",clothesColor) // 打印：今天我穿红色的衣服

clothesColor = "黄色"  // 明天我穿黄色的衣服

fmt.Printf("明天我穿%s的衣服\n",clothesColor) // 打印：明天我穿黄色的衣服
*/

/*
变量的作用域
一个变量作用的范围，区间或区域。超过了一定的范围，这个变量就不能使用了。

package main

import "fmt"

func main(){
	var aaa int = 10
	fmt.Println(aaa)

	var bbb int = 20
	fmt.Println(bbb)
}

*/
