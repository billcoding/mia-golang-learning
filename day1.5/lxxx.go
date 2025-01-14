package main

import "fmt"

func main() {
	// var clothesColor string = "红色" // 今天我穿红色的衣服，声明一个名称为clothesColor类型为string值为红色的变量。
	// 	var clothesColor string // 先给我声明1个变量叫clothesColor 放着
	// 	clothesColor="红色"

	// 	fmt.Printf("今天我穿%s的衣服\n", clothesColor) // 打印：今天我穿红色的衣服

	// 	clothesColor = "黄色" // 明天我穿黄色的衣服

	// 	fmt.Printf("明天我穿%s的衣服\n", clothesColor) // 打印：明天我穿黄色的衣服
	/*

	 */

	// fmt.Println("小明的学生信息")
	// fmt.Println("---")
	var name string = "小明"
	var gender string = "男"
	var age int = 12
	var height int = 135
	var weight int = 35
	var homeaddress string = "湖北武汉"
	// fmt.Printf("姓名：%s\n",name)
	// fmt.Printf("性别：%s\n",gender)
	// fmt.Printf("年龄：%d岁\n",age)
	// fmt.Printf("身高：%dcm\n",height)
	// fmt.Printf("体重：%dkg\n",weight)
	// fmt.Printf("家庭地址：%s\n",homeaddress)
	//  }

	//fmt.Println("姓名：",name,"\n","性别：",gender,"\n","年龄：",age,"岁","\n","身高：",height,"cm","\n","体重：",weight,"kg","\n","家庭地址：",homeaddress,"\n")
	fmt.Printf("小明的学生信息\n---\n姓名：%s\n性别：%s\n年龄：%d岁身高：%dcm\n体重：%dkg家庭住址：%s\n\n", name, gender, age, height, weight, homeaddress)
	//  fmt.Println()
	//  fmt.Println("小红的学生信息")
	//  fmt.Println("---")
	name = "小红" // 这是声明变量的语法 需要替换为 给这个变量赋值
	gender = "女"
	age = 11
	height = 130
	weight = 31
	homeaddress = "湖北仙桃"
	// fmt.Println("姓名：",name,"\n","性别：",gender,"\n","年龄：",age,"岁","\n","身高：",height,"cm","\n","体重：",weight,"kg","\n","家庭地址：",homeaddress,"\n")
	fmt.Printf("小红的学生信息\n---\n姓名：%s\n性别：%s\n年龄：%d岁\n身高：%dcm\n体重：%dkg\n家庭住址:%s\n", name, gender, age, height, weight, homeaddress)
}
