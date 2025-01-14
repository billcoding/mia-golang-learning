package main

import "fmt"

func main() {
	// 变量引用采用的是就近原则，前提是 可以访问 到有 不同作用域 的 相同名称的变量（多个）
	var name string = "我菜"
	fmt.Println(name)

	{
		var name string = "张三" //声明1个name变量
		fmt.Println(name)
		// 因为这里的变量name的作用域是 第7行 - 第9行
	}

	fmt.Println(name) // 这里不存在就近原则。 打印结果是什么

	{
		name = "你菜"            //重新赋值了 赋值为你菜
		var name string = "李四" //又声明1个name变量
		// age这个名称 很奇怪 因为它的值实际上不是代表年龄 而是姓名
		fmt.Println(name)
	}

	{
		name = "爷菜" //重新赋值了 赋值为爷菜
		// 变量引用采用的是就近原则，前提是 可以访问 到有 不同作用域 的 相同名称的变量（多个）
		var name string = "她菜"
		fmt.Println(name) //
		name = "你们都菜"
		fmt.Println(name) //你们都菜
	}

	fmt.Println(name) //爷菜
}
