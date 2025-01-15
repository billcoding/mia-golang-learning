// 整型的运算
// 数学里有+-*/，Go语言也有这些运算。

package main

import "fmt"

func main() {
	{
		// 加法
		var int1 int = 15
		var int2 int = 11
		// 这里有几个占位（坑），3个占位，从第2个参数开始，是不是也要连续给3个值（填坑）
		// 11+5=16
		fmt.Printf("%d+%d=%d\n", int1, int2, (int1 + int2))
	}

	{
		// 减法
		var int1 int = 20
		var int2 int = 15
		fmt.Printf("%d-%d=%d\n", int1, int2, (int1 - int2))
	}

	{
		//乘法
		var int1 int = 5
		var int2 int = 6
		fmt.Printf("%d*%d=%d\n", int1, int2, (int1 * int2))
	}

	{
		//除法，2个数相除，结果只能是整数，不管小数点后有多少位，全部抹掉
		//除数不能为0，否则报错
		var int1 int = 12
		var int2 int = 2
		fmt.Printf("%d/%d=%d\n", int1, int2, (int1 / int2))
	}

	{
		// 加等于+=
		var int1 int = 10 //声明1个名称为int1类型为int值为10的变量
		int1 = int1 + 5   // 重新给int1赋值为int1+5 原来int1是10，现在是10+5=15
		// int1 = 15
		int1 += 5       // 这里就是加等于运算 => 等价于 int1 = int1 + 5
		int1 = 12       // 12
		int1 += 4       // => int1 = int1 + 4 => int1 = 12 + 4 = 16
		int1 = int1 + 6 // 16 + 6 = 22
		fmt.Println("int1 =", int1)
	}

	{
		// 减等于-=
		var int1 int = 20
		int1 = int1 - 8 //12
		int1 -= -4      //int1 = int1 - (-4) => int1 + 4 => 16
		int1 = 16       //16
		int1 -= 8 - 1   //9
		int1 = int1 - 3 // 6
		fmt.Println("int1 =", int1)
	}

	{
		// 自增 ++ 自己+1
		var var1 int = 10
		var1++    //自增一下
		var1 += 1 // var1 += 1 、var1++ 、var1 = var1 + 1 三种写法都是等价的
		var1 += 2
		var1++
		var1 += 3
		var1++
		fmt.Println("var1 =", var1)
	}

	{
		// 自减 -- 自己-1
		var var2 int = 20
		var2-- //自减一下
		var2 -= 15
		var2 = var2 - 5
		var2--
		fmt.Println("var2", var2)
	}

}
