package main

import "fmt"

/*
现在有一个需求：

学生信息如下：
学号     名字     性别
---------------------
86      王一毛      男
71      李秋水      女
92      赵逸诚      男
55      蔡点灯      男
68      蔡宝倍      女
35      王毛球      女

将6个学生按照按如下要求进行分组做游戏，最后将每组的人员打印出来。

分组要求：
PP组：学号大于80、性别男
BB组：学号大于70并且小于80、性别女
CC组：另外的人进CC组

打印格式如下：
PP组有：姓名1、姓名2
BB组有：姓名1、姓名2
CC组有：姓名1、姓名2

*/
func main() {
	var (
		groupPP, groupBB, groupCC    string = "PP", "BB", "CC"
		memberPP, memberBB, memberCC string = "", "", ""
		numb1                        int    = 86
		name1, gender1               string = "王一毛", "男"
		numb2                        int    = 71
		name2, gender2               string = "李秋水", "女"
		numb3                        int    = 92
		name3, gender3               string = "赵逸诚", "男"
		numb4                        int    = 55
		name4, gender4               string = "蔡点灯", "男"
		numb5                        int    = 68
		name5, gender5               string = "蔡宝倍", "女"
		numb6                        int    = 35
		name6, gender6               string = "王毛球", "女"
	)

	if numb1 > 80 && gender1 == "男" {
		memberPP = name1
	} else if numb1 > 70 && numb1 <= 80 && gender1 == "女" {
		memberBB = name1
	} else {
		memberCC = name1
	}

	// fmt.Printf("%d,%s,%s",numb1,name1,gender1)

	if numb2 > 80 && gender2 == "男" {
		memberPP = name2
	} else if numb2 > 70 && numb2 <= 80 && gender2 == "女" {
		memberBB = name2
	} else {
		memberCC = name2
	}

	if numb3 > 80 && gender3 == "男" {
		memberPP += "、" + name3
	} else if numb3 > 70 && numb3 <= 80 && gender3 == "女" {
		memberBB += "、" + name3
	} else {
		memberCC += "、" + name3
	}

	if numb4 > 80 && gender4 == "男" {
		memberPP = name4
	} else if numb4 > 70 && numb4 <= 80 && gender4 == "女" {
		memberBB = name4
	} else {
		memberCC = name4
	}

	if numb5 > 80 && gender5 == "男" {
		memberPP += "、" + name4
	} else if numb5 > 70 && numb5 <= 80 && gender5 == "女" {
		memberBB += "、" + name5
	} else {
		memberCC += "、" + name5
	}

	if numb6 > 80 && gender6 == "男" {
		memberPP += "、" + name6
	} else if numb6 > 70 && numb6 <= 80 && gender6 == "女" {
		memberBB += "、" + name6
	} else {
		memberCC += "、" + name6
	}

	fmt.Printf("%s组有：%s\n", groupPP, memberPP)
	fmt.Printf("%s组有：%s\n", groupBB, memberBB)
	fmt.Printf("%s组有：%s\n", groupCC, memberCC)
}
