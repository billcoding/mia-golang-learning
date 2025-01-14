/*
练习1.1：变量重新赋值运行

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
	fmt.Println("小明的学生信息")
	fmt.Println("---")
	var name string = "小明"
	var gender string = "男"
	var age int = 12
	var height int = 135
	var weight int = 35
	var homeaddress string = "湖北武汉"
	fmt.Printf("姓名：%s\n", name)
	fmt.Printf("性别：%s\n", gender)
	fmt.Printf("年龄：%d岁\n", age)
	fmt.Printf("身高：%dcm\n", height)
	fmt.Printf("体重：%dkg\n", weight)
	fmt.Printf("家庭地址：%s\n", homeaddress)

	fmt.Println()
	fmt.Println("小红的学生信息")
	fmt.Println("---")
	name = "小红"
	gender = "女"
	age = 11
	height = 130
	weight = 31
	homeaddress = "湖北仙桃"
	fmt.Printf("姓名：%s\n", name)
	fmt.Printf("性别：%s\n", gender)
	fmt.Printf("年龄：%d岁\n", age)
	fmt.Printf("身高：%dcm\n", height)
	fmt.Printf("体重：%dkg\n", weight)
	fmt.Printf("家庭住址：%s\n", homeaddress)
}
