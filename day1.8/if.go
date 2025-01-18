package main

import "fmt"

/*
if-else语句
格式：

if 条件表达式 {

}else {

}

条件表达式最后的结果是一个布尔值，所以要么真要么假。
else { }是可以不写，具体在什么时候不写是要看情况的。
*/

func main() {

	{
		var var1 int = 10

		// 如果变量var1的值大于10，就打印var1的值大于10，否则就打印var1的值小于等于10
		if var1 > 10 {
			fmt.Println("var1的值大于10")
		} else {
			fmt.Println("var1的值小于等于10")
		}
	}

	{
		var name string = "李明"
		var name1 string = "威尔斯"
		if name < name1 {
			fmt.Println("name的值小于name1")

		} else {
			fmt.Printf("name的值大于等于name1\n")
		}
	}

	{
		var a int = 15
		var b int = 30
		if a > b {
			fmt.Println("a的值大于b")
		} else {
			fmt.Println("a的值小于等于b")

		}
	}

	{
		if 30 < 20 {
			fmt.Println("30 < 20")
		} else {
			fmt.Print("30 >= 20\n")
		}
	}

	{
		var temprature int = 25
		if temprature > 25 {
			fmt.Println("天气温暖")
		} else {
			fmt.Println("天气寒冷")
		}
	}

	{
		var numb1 int = 10
		var numb2 int = 20
		var numb3 int = 30
		if numb1 < numb2 {
			fmt.Println("numb1小于numb2")
		}
		if numb2 < numb3 {
			fmt.Println("numb2小于numb3")
		} else {
			fmt.Println("numb1大于等于numb2")
		}
	}

	{
		// var name string = "lili"
		var weighta int = 50
		var weightb int = 55
		var agea int = 55
		var ageb int = 60
		var heighta int = 150
		var heightb int = 160
		if weighta > weightb && agea < ageb && heighta < heightb {
			fmt.Println("a的体重大于b的体重与a的年龄小于b年龄与a的体重小于b的体重的值是真的")
		} else {
			fmt.Println("a的体重大于b的体重与a的年龄小于b年龄与a的体重小于b的体重的值是假的")
		}
	}

	{
		var namexf string = "小芳"
		var namexm string = "小明"
		var agexf int = 20
		var agexm int = 29
		var weightxf int = 55
		var weightxm int = 62
		if namexf < namexm || agexf > agexm || weightxf == weightxm {
			fmt.Println("小芳的姓名小于小明的姓名或小芳的年龄大于小明的年龄或小芳的体重等于小明的体重是真的")
		} else if agexf < agexm || weightxf < weightxm {
			fmt.Println("小芳的年龄大于小明的年龄或小芳的体重等于小明的体重是真的")
		} else {
			fmt.Println("小芳的姓名小于小明的姓名或小芳的年龄大于小明的年龄或小芳的体重等于小明的体重是假的")
		}

	}

	{

		var numb1 float64 = 5.66
		var numb2 float64 = 7.329
		fmt.Printf("%v+%v=%v\n", numb1, numb2, numb1+numb2)
		if numb2 < numb1 {
			fmt.Println("numb2的值小于numbe1的值")
		} else if numb2 == numb1 {
			fmt.Println("numb2的值等于numbe1的值")
		} else {
			fmt.Println("numb2的值大于等于numbe1的值")
		}
	}
	{
		var name string = "mia"
		var name1 string = "bill"
		fmt.Print("name+name1=", name+name1, "\n")
		var numb int = 13        //整数
		var numb1 int = 25       //整数
		var numb2 float64 = 2.55 //小数
		fmt.Printf("%v+%v+%v=%v\n", numb, numb1, numb2, numb+numb1+int(numb2))
	}

	{
		var int1 int = 10
		var float1 float64 = float64(int1)
		fmt.Printf("%f\n", float1)

		var float2 float64 = 10.999
		var int2 int = int(float2)
		fmt.Printf("%d\n", int2)
	}
	{
		var numb float64 = 3.58
		var numb1 float64 = 5.0
		fmt.Printf("%v+%v=%v", numb, numb1, numb+numb1)

	}
}

// 类型转换：

// 整数->小数
// var int1 int = 10
// var float1 float64 = float64(int1)

// 小数->整数，注意小数点后面的会丢失
// var float1 float64 = 10.1
// var int1 int = int(float1)
// 10.1 -> 10
// 10.9 -> 10

// 整数和小数，不能直接一起计算 ，要么都转换为整数，要么都转换为小数。
