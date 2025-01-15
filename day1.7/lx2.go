package main

import "fmt"

// 条件表达式的运算
// 并且 &&
// 或者 ||
// 小A的体重比小B的体重重 或者 小C的体重比小A的体重轻
// 小A的体重比小B的体重重 并且 小C的体重比小A的体重重
// 2个条件表达式运算后仍是1个表达式

// b1 && b2 && b3 ...，b1和b2都为真时，则最后的表达式为真。
// b1 || b2 || b3 ...，b1和b2其中有1个为真时，则最后的表达式为真。

func main() {
	// 现在有一个班级，有3个学生，分别是：小刚、小华和小菊。
	// 他们的年龄分别是12、11、15岁
	// 他们的身高分别是150、140、155cm
	// 他们的体重分别是35、32、36kg
	// 按需求先声明变量，然后做如下判断并打印。
	var ageXG int = 12
	var ageXH int = 11
	var ageXJ int = 15
	fmt.Printf("小刚的年龄：%d岁\n", ageXG)
	fmt.Printf("小华的年龄：%d岁\n", ageXH)
	fmt.Printf("小菊的年龄：%d岁\n", ageXJ)

	var heightXG int = 150
	var heightXH int = 140
	var heightXJ int = 155
	fmt.Printf("小刚的身高：%dcm\n", heightXG)
	fmt.Printf("小华的身高：%dcm\n", heightXH)
	fmt.Printf("小菊的身高：%dcm\n", heightXJ)

	var weightXG int = 35
	var weightXH int = 32
	var weightXJ int = 36
	fmt.Printf("小刚的体重：%dkg\n", weightXG)
	fmt.Printf("小华的体重：%dkg\n", weightXH)
	fmt.Printf("小菊的体重：%dkg\n", weightXJ)

	// 1. 小刚的年龄是否大于小华的年龄
	var b1 bool = ageXG > ageXH
	fmt.Printf("小刚的年龄是否大于小华的年龄:%t\n", b1)

	// 2. 小华的体重是否小于小菊的体重
	var b2 bool = weightXH < weightXJ
	fmt.Printf("小华的体重是否小于小菊的体重:%t\n", b2)

	// 3. 小菊的身高是否等于小华的身高
	var b3 bool = heightXJ == heightXH
	fmt.Printf("小菊的身高是否等于小华的身高:%t\n", b3)

	// 4. 小刚的年龄大于小华的年龄 并且 小菊的身高小于小刚的身高
	var b4 bool = (ageXG > ageXH) && (heightXJ < heightXG)
	fmt.Printf("小刚的年龄大于小华的年龄 并且 小菊的身高小于小刚的身高%t\n", b4)

	// 5. 小华的体重小于小菊的体重 或者 小刚的年龄大于等于小华的年龄
	var b5 bool = (weightXH < weightXJ) || (ageXG >= ageXH)
	fmt.Printf("小华的体重小于小菊的体重 或者 小刚的年龄大于等于小华的年龄%t\n", b5)

	// 6. 小刚的体重小于等于小华的体重 并且 小刚的年龄小于等于小华的年龄 并且 小菊的身高小于小刚的身高
	var b6 bool = (weightXG <= weightXH) && (ageXG <= ageXH) && (heightXJ < heightXG)
	fmt.Printf("小刚的体重小于等于小华的体重 并且 小刚的年龄小于等于小华的年龄 并且 小菊的身高小于小刚的身高%t\n", b6)
}
