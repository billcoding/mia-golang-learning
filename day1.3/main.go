// 入口包，主包
package main

// main 主要的
// package 包
// func 函数
// function 函数
// undeclared 未定义、未声明
// arguments 参数
// args 参数
// return 返回
// values 值
// call 调用

// fmt.Print 打印后不换行
// fmt.Printf 格式化打印后不换行
// fmt.Println 打印后换行

// 语法
// 警告：是黄色的波浪线，可以运行，但是可能不符合预期。
// 错误：是红色的波浪线，不可以运行，运行直接会编译报错。

import "fmt"

func main() {
	// fmt -> format 格式化
	// print line 打印并且换行
	fmt.Print("你好吗")   // 打印不换行
	fmt.Println("我爱你") // 打印并且换行
	fmt.Print("你呢\n")  // 打印换行

	// %d 代表1个数字
	// %s 代表1个字符串
	// %v 兼容一切，什么都可以用%v
	var a int = 10
	var b int = 20
	fmt.Println("a=", a, "b=", b)
	fmt.Printf("名字是%v，年龄是%v岁，身高是%vcm，兴趣是%v\n", "张三", 28, 175, "跳舞+唱歌")
}
