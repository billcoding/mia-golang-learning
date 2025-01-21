/*
现在有一个需求：

学生信息如下：
学号     名字     性别
---------------------
86      王一毛      男
71      李秋水      女
92      赵逸城      男
55      蔡点灯      男
68      蔡宝倍      女
35      王毛球      女

将6个学生按照按如下要求进行分组做游戏，最后将每组的人员打印出来。

分组要求：
PP组：学号大于80、性别男（30，女）
BB组：学号大于70并且小于80、性别女（50~80，男）
CC组：另外的人进CC组

打印格式如下：
PP组有：姓名1、姓名2
BB组有：姓名1、姓名2
CC组有：姓名1、姓名2
*/
package main

import "fmt"

func main() {
	var (
		numb1, numb2, numb3, numb4, numb5, numb6 int    = 86, 71, 92, 55, 68, 35
		name1, gender1                           string = "王一毛", "男"
		name2, gender2                           string = "李秋水", "女"
		name3, gender3                           string = "赵逸城", "男"
		name4, gender4                           string = "蔡点灯", "男"
		name5, gender5                           string = "蔡宝倍", "女"
		name6, gender6                           string = "王毛球", "女"
		groupPP, groupBB, groupCC                string = "PP组", "BB组", "CC组"
		memberPP, memberBB, memberCC             string = "", "", ""
	)
	if numb1 > 30 && gender1 == "女" {
		memberPP = name1
	} else if numb1 > 50 && numb1 < 80 && gender1 == "男" {
		memberBB = name1
	} else {
		memberCC = name1
	}

	if numb2 > 30 && gender2 == "女" {
		memberPP = name2
	} else if numb2 > 50 && numb2 < 80 && gender2 == "男" {
		memberBB = name2
	} else {
		memberCC = name2
	}

	if numb3 > 30 && gender3 == "女" {
		memberPP += "," + name3
	} else if numb3 > 50 && numb3 < 80 && gender3 == "男" {
		memberBB += "," + name3
	} else {
		memberCC += "," + name3
	}

	if numb4 > 30 && gender4 == "女" {
		memberPP = name4
	} else if numb4 > 50 && numb4 < 80 && gender4 == "男" {
		memberBB = name4
	} else {
		memberCC = name4
	}

	if numb5 > 30 && gender5 == "女" {
		memberPP += "," + name5
	} else if numb5 > 50 && numb5 < 80 && gender5 == "男" {
		memberBB += "," + name5
	} else {
		memberCC += "," + name5
	}

	if numb6 > 30 && gender6 == "女" {
		memberPP += "," + name6
	} else if numb6 > 50 && numb6 < 80 && gender6 == "男" {
		memberBB += "," + name6
	} else {
		memberCC += "," + name6
	}

	fmt.Printf("%s组：%s\n", groupPP, memberPP)
	fmt.Printf("%s组：%s\n", groupBB, memberBB)
	fmt.Printf("%s组：%s\n", groupCC, memberCC)
}
