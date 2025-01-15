package main

import "fmt"

// 字符串的运算
// 字符串只有1个运算符，就相加，也可以说拼接。
// 字符串相加后，仍然是1个字符串

func main() {
	var s1 string = "hello"
	fmt.Println(s1)
	var s2 string = "world"
	fmt.Println(s2)

	//将s1和s2相加起来（拼接起来）
	var s3 string = s1 + s2
	fmt.Printf("%s+%s=%s\n", s1, s2, s3)

	// fmt.Print("11", "22", "33")          //112233
	// fmt.Println("wow\nmam,person,\nwaa") // Println这个函数，不管里面是否有内容，都会在最后换1个行
	fmt.Printf("早上%d点钟%s%s和%v%s今天一起上学了\n下午%s点放学时\n他们又一起%s了。",
		8, "\n", "小A", "\n", "小B", "5", "回家")
	// 早上8点钟
	// 小A和
	// 小B今天一起上学了
	// 下午5点放学时
	// 他们又一起回家了。
}
