package main

import "fmt"

// 条件表达式，这个表达式是可以进行判断，最终的结果是不是真就是假。
// 顾名思义就是代表一个条件（实际上可以作为判断的标准的）的表达式

// 用于判断的符号
// 大于 >
// 小于 <
// 大于等于 >=
// 小于等于 <=
// 等于 == （我猜想 因为=是代表赋值，所以等于号是==）

// %d代表一个int类型的占位（坑）
// %s代表一个string类型的占位（坑）
// %t代表一个bool类型的占位（坑）
// %v代表一个any类型的占位（坑）
func main() {
	var b1 bool = true // 声明1个名称为b1类型为bool值为true的变量
	fmt.Printf("b1=%t\n", b1)
	fmt.Printf("b1=%v\n", b1)

	// 条件表达式
	// 比如说，小明的身高是160cm，小红的身高是150cm。那么我们判断一下小明的身高大于小红的身高吗？
	var heightXM int = 160
	var heightXH int = 150
	fmt.Printf("2024年，小明的身高是%dcm\n", heightXM)
	fmt.Printf("2024年，小红的身高是%dcm\n", heightXH)

	// 回忆一下数学中如何判断一个数是否大于另外一个数的。
	// 120 > 100
	// fmt.Printf("%d\n", 100+200)
	// fmt.Printf("小明的身高大于小红的身高是%t的\n", (heightXM > heightXH)) // 这里实际上是一个布尔值 true和false
	var b2 bool = (heightXM > heightXH)        // (heightXM > heightXH) 这就是一个条件表达式
	fmt.Printf("2024年，小明的身高大于小红的身高是%t的\n", b2) // 这里实际上是一个布尔值 true和false

	// 一年后，小明的身高变成了165cm，小红的身高变成了170cm
	heightXM = 165
	heightXH = 170
	// 首先打印，2025年，他们2个的身高 分别是多少
	fmt.Printf("2025年，小明的身高是%dcm\n", heightXM)
	fmt.Printf("2025年，小红的身高是%dcm\n", heightXH)
	// 然后，判断小明的身高是否大于小红的身高。
	b2 = (heightXM > heightXH)
	fmt.Printf("2025年，小明的身高大于小红的身高是%t的\n", b2)
}
