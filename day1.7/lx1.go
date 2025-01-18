/*
现有3名同学参与运动会，姓名分别是 李斯特，朱毅炎，胡一帆，
他们的年龄分别是：8岁，9岁，7岁，
他们的体重是：35，40，42，
他们的身高是：135，150，145，
跑步成绩排名是：第三名，第二名，第一名
一分钟跳绳完成数量：75，100，60，
*/
package main

import "fmt"

func main() {
	{

		var agelst int = 8
		fmt.Printf("李斯特的年龄是：%d岁\n", agelst)
		var agezyy int = 9
		fmt.Printf("朱毅炎的年龄是：%d岁\n", agezyy)
		var agehyf int = 7
		fmt.Printf("胡一帆的年龄是：%d岁\n", agehyf)

		var weightlst int = 35
		fmt.Printf("李斯特的体重是：%dkg\n", weightlst)
		var weightzyy int = 40
		fmt.Printf("朱毅炎的体重是：%dkg\n", weightzyy)
		var weighthyf int = 42
		fmt.Printf("胡一帆的体重是：%d岁kg\n", weighthyf)

		var rankinglst string = "第三名"
		fmt.Printf("李斯特跑步的排名是：%v\n", rankinglst)
		var rankingzyy string = "第二名"
		fmt.Printf("朱毅炎跑步的排名是：%v\n", rankingzyy)
		var rankinghyf string = "第一名"
		fmt.Printf("胡一帆跑步的排名是：%v\n", rankinghyf)

		var heightlst int = 8
		fmt.Printf("李斯特的身高是：%dcm\n", heightlst)
		var heightzyy int = 9
		fmt.Printf("朱毅炎的身高是：%dcm\n", heightzyy)
		var heighthyf int = 7
		fmt.Printf("胡一帆的身高是：%dcm\n", heighthyf)

		var numb1 int = 75
		fmt.Printf("李斯特一分钟跳绳的数量是：%d个\n", numb1)
		var numb2 int = 100
		fmt.Printf("朱毅炎一分钟跳绳的数量是：%d个\n", numb2)
		var numb3 int = 60
		fmt.Printf("胡一帆一分钟跳绳的数量是：%d个\n", numb3)

		var a bool = numb1 < numb2
		fmt.Printf("李斯特一分钟跳绳的数量小于朱毅炎一分钟跳绳数量是%t\n", a)
		var b bool = heightlst > heighthyf
		fmt.Printf("李斯特的身高大于胡一帆的身高是%t\n", b)
		var b1 bool = weightzyy > weighthyf
		fmt.Printf("朱毅炎的体重大于胡一帆的体重是%t\n", b1)
		var b2 bool = weightzyy > weighthyf || heightlst > heighthyf
		fmt.Printf("李斯特的体重大于胡一帆的体重 或者 李斯特的身高大于胡一帆的身高是%t\n", b2)
		var b3 bool = numb1 == numb2 && rankinglst < rankinghyf
		fmt.Printf("李斯特一分钟跳绳的数量等于朱毅炎一分钟跳绳的数量 并且 李斯特的排名小于朱毅炎的排名是%t\n ", b3)

	}

	{
		var a string = "miss"
		var b string = "you"
		fmt.Printf("%s+%s=%s\n", a, b, a+b)
		a = "mia"
		b = "bill"
		fmt.Printf("%s+%s=%s\n", a, b, a+b)
		var a1 string
		var b1 string = "mia"
		fmt.Printf("%s+%s=%s\n", a1, b1, a1+b1)

	}

	{
		var a int = 15
		var b int = 20
		fmt.Printf("%d+%d=%d\n", a, b, a+b)
		a = 10
		b = 15
		fmt.Printf("%d+%d=%d\n", a, b, a+b)

		var c bool = a > b
		fmt.Printf("a>b是%t的", c)

	}
	fmt.Println()
	{

		var placeA string = "ABCD"
		fmt.Println("placeA=", placeA)
		fmt.Printf("placeA=%s\n", placeA)

	}

}
