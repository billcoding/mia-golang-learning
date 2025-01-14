package main

import "fmt"

func main() {
	// 说了不建议使用中文作为变量的名称的，你还在使用
	var numb1 int = 50
	var numb2 int = 15
	fmt.Println(numb1, numb2)
	//fmt.Println(numb1 + numb2)
	//fmt.Printf("numb1是%d,numb2是%d",50,10)
	fmt.Print(numb1,numb2)
}
