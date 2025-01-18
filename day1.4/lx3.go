package main

import "fmt"

func main() {
	// 说了不建议使用中文作为变量的名称的，你还在使用
	var numb1 int = 50
	var numb2 int = 15
	fmt.Print(numb1, numb2)
	//fmt.Println(numb1 + numb2)
	//fmt.Printf("numb1是%d,numb2是%d",50,10)
	fmt.Println()
	fmt.Println(numb1, numb2)
	fmt.Printf("%d+%d=%d\n", numb1, numb2, numb1+numb2)

	var numb3 int = 25
	var c bool = (numb1 > numb2) || (numb3 == numb2)
	fmt.Printf("numb1大于numb2 或者 numb3==numb2 是%t的\n", c)
	var d bool = numb2 > numb3 && numb3 < numb1
	fmt.Printf("numb2大于numb3 和  numb3<numb1 是%t的\n", d)

}
