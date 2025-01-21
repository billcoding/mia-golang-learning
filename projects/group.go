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

		no2            = 71
		name2, gender2 = "李秋水", "女"

		no3            = 92
		name3, gender3 = "赵逸城", "男"

		no4            = 55
		name4, gender4 = "蔡点灯", "男"

		no5            = 68
		name5, gender5 = "蔡宝倍", "女"

		no6            = 35
		name6, gender6 = "王毛球", "女"
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

	// 定位到这块地方的代码有问题
	// 我们刚才说了，赵逸城是被分到PP组了

	if (no6 > 80) && (gender6 == "男") {
		// 满足PP组，放到PP组里的学生里去
		// 将该学生的名字加到PP组的学生里去
		if personPP != "" {
			personPP += "、"
		}
		personPP += name6
	} else if ((no6 > 70) && (no6 <= 80)) && (gender6 == "女") {
		// 满足BB组的，放到BB组里的学生里去
		// 将该学生的名字加到PP组的学生里去
		if personBB != "" {
			personBB += "、"
		}
		personBB += name6
	} else {
		if personCC != "" {
			personCC += "、"
		}
		personCC += name6 // 将fullName6存入到personCC中
	}

	if (no1 > 80) && (gender1 == "男") {
		// 满足PP组，放到PP组里的学生里去
		// 将该学生的名字加到PP组的学生里去
		if personPP != "" {
			personPP += "、"
		}
		personPP += name1
	} else if ((no1 > 70) && (no1 <= 80)) && (gender1 == "女") {
		// 满足BB组的，放到BB组里的学生里去
		// 将该学生的名字加到PP组的学生里去
		if personBB != "" {
			personBB += "、"
		}
		personBB += name1
	} else {
		if personCC != "" {
			personCC += "、"
		}
		personCC += name1 // 将fullName1存入到personCC中
	}

	if (no3 > 80) && (gender3 == "男") {
		// 满足PP组，放到PP组里的学生里去
		// 将该学生的名字加到PP组的学生里去
		// 问题是不是在这里？？？？
		// 解释下 personPP = name3 的意思：
		// 是不是将name3赋值给personPP，（personPP赋值为name3)，这个是一个重新赋值的过程。
		// 原来的值会被替换掉
		// personPP目前是王一毛，我们预期的结果是：王一毛、赵逸城
		// 目前是 赵逸城
		// 我们观察代码 发现，person
		// 我们是不是可以这样做？
		// 在分组前（将人员加入组之前），判断下组里是否有成员？
		// 如果有成员，先追加一个、再将学生加入到组内
		// 如果没有成员（否则），直接将学生加入到组内
		if personPP != "" {
			personPP += "、"
		}
		personPP += name3
	} else if ((no3 > 70) && (no3 <= 80)) && (gender3 == "女") {
		// 满足BB组的，放到BB组里的学生里去
		// 将该学生的名字加到PP组的学生里去
		if personBB != "" {
			personBB += "、"
		}
		personBB += name3
	} else {
		if personCC != "" {
			personCC += "、"
		}
		personCC += name3 // 将fullName3存入到personCC中
	}

	if (no2 > 80) && (gender2 == "男") {
		// 满足PP组，放到PP组里的学生里去
		// 将该学生的名字加到PP组的学生里去
		if personPP != "" {
			personPP += "、"
		}
		personPP += name2
	} else if ((no2 > 70) && (no2 <= 80)) && (gender2 == "女") {
		// 满足BB组的，放到BB组里的学生里去
		// 将该学生的名字加到PP组的学生里去
		if personBB != "" {
			personBB += "、"
		}
		personBB += name2
	} else {
		if personCC != "" {
			personCC += "、"
		}
		personCC += name2 // 将fullName2存入到personCC中
	}

	if (no4 > 80) && (gender4 == "男") {
		// 满足PP组，放到PP组里的学生里去
		// 将该学生的名字加到PP组的学生里去
		if personPP != "" {
			personPP += "、"
		}
		personPP += name4
	} else if ((no4 > 70) && (no4 <= 80)) && (gender4 == "女") {
		// 满足BB组的，放到BB组里的学生里去
		// 将该学生的名字加到PP组的学生里去
		if personBB != "" {
			personBB += "、"
		}
		personBB += name4
	} else {
		if personCC != "" {
			personCC += "、"
		}
		personCC += name4 // 将fullName4存入到personCC中
	}

	if (no5 > 80) && (gender5 == "男") {
		// 满足PP组，放到PP组里的学生里去
		// 将该学生的名字加到PP组的学生里去
		if personPP != "" {
			personPP += "、"
		}
		personPP += name5
	} else if ((no5 > 70) && (no5 <= 80)) && (gender5 == "女") {
		// 满足BB组的，放到BB组里的学生里去
		// 将该学生的名字加到PP组的学生里去
		if personBB != "" {
			personBB += "、"
		}
		personBB += name5
	} else {
		if personCC != "" {
			personCC += "、"
		}
		personCC += name5 // 将fullName5存入到personCC中
	}

	// 问题描述：王一毛是第一个学生，但是没有出现，我们分析他的学生信息和各组的条件，发现他满足PP组。但是他并没有出现在PP组的名单里。
	// PP组目前只出现了赵逸城，说明王一毛被赵逸城顶替了。
	// 那我们应该检测那一块地方的代码呢？
	// 肯定是 给赵逸城分组的时候 出了问题。
	fmt.Printf("%s有：%s\n", groupPP, personPP)
	fmt.Printf("%s有：%s\n", groupBB, personBB)
	fmt.Printf("%s有：%s\n", groupCC, personCC)
}
