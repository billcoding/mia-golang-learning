package main

import "fmt"

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
PP组：学号大于80、性别男
BB组：学号大于70并且小于80、性别女
CC组：另外的人进CC组

打印格式如下：
PP组有：姓名1、姓名2
BB组有：姓名1、姓名2
CC组有：姓名1、姓名2

*/

func main() {
	// 分析：
	// 这里面涉及有变量吗？
	// 这里面涉及有判断语句？
	// 首先把 3个组 变量声明出来
	var (
		groupPP, groupBB, groupCC    = "PP组", "BB组", "CC组" // 这里3个变量是存放组的名称
		personPP, personBB, personCC = "", "", ""          // 这里3个变量是存放组的学生
	)

	// 学生信息的变量
	var (
		no1            = 86
		name1, gender1 = "王一毛", "男"

		// no2            = 71
		// name2, gender2 = "李秋水", "女"

		// no3            = 92
		// name3, gender3 = "赵逸城", "男"
		//
		// no4            = 55
		// name4, gender4 = "蔡点灯", "男"
		//
		// no5            = 68
		// name5, gender5 = "蔡宝倍", "女"
		//
		// no6            = 35
		// name6, gender6 = "王毛球", "女"
	)

	// 分组
	// 根据要求进行判断分组
	// 分组要求：
	// PP组：学号大于80、性别男
	// BB组：学号大于70并且小于80、性别女
	// CC组：另外的人进CC组
	// 一共有6个学生，那么给这6个人分组呢？
	// 实际上是1个个的分的，就是先分第1个人，再分第2个人，以此类推。

	// 先分 王一毛，判断他该分到哪个组
	// 开始判断
	// 王一毛 如果不满足PP组的要求 怎么办？
	// 再判断他 满不满足 BB组的要求
	// 如果以上都不满足 就放到CC组

	if (no1 > 80) && (gender1 == "男") {
		// 满足PP组，放到PP组里的学生里去
		// 将该学生的名字加到PP组的学生里去
		personPP = name1
	} else if ((no1 > 70) && (no1 <= 80)) && (gender1 == "女") {
		// 满足BB组的，放到BB组里的学生里去
		// 将该学生的名字加到PP组的学生里去
		personBB = name1
	} else {
		personCC = name1 // 将fullName1存入到personCC中
	}

	fmt.Printf("%s有：%s\n", groupPP, personPP)
	fmt.Printf("%s有：%s\n", groupBB, personBB)
	fmt.Printf("%s有：%s\n", groupCC, personCC)
}
