package main

import "fmt"

// 变量和数据类型
// 声明变量的格式，变量类型可以省略，会自动推导变量的类型。
// var 变量名 变量类型 = 变量值

// 先声明变量（注意变量类型不能省略），后赋值
// var 变量名 变量类型
// 变量名 = 变量值

// 变量的二次赋值

// 变量的默认值
// 整型int变量的默认值是0
// 字符串string变量的默认值是空
// 布尔bool变量的默认值是false（就是否）

//布尔类型，也是一种变量类型。有点特别，只有2个值：true真和false假（或者：是和否）。
// 声明1个假值的布尔变量
// var b1 bool = false
// 声明1个真值的布尔变量
// var b2 bool = true

func main() {
	var var1 int = 10 // 声明1个名称为var1类型为int值为10的变量
	fmt.Println("var1 =", var1)

	var var2 int // 只是声明这个变量，但是没有赋值，但是实际上会有默认值
	// var var2 int = 0 与上面是等价的
	fmt.Println("var2 =", var2)

	var var3 string
	//var var3 string = "" 与上面是等价的
	fmt.Println("var3 =", var3)

	var var4 bool
	fmt.Println("var4 =", var4)

	var var5 bool = true
	fmt.Println("var5 =", var5)
}
