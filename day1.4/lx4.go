package main

import "fmt"

func main() {
	var class1 string = "Chinese"
	var class2 string = "Math"
	var class3 string = "English"
	fmt.Println("课程1:", class1, "课程2:", class2, "课程3:", class3)
	fmt.Println("课程1:", class1, "\n", "课程2:", class2, "\n", "课程3:", class3, "\n")
	fmt.Printf("课程1是%v\n课程2是%v\n课程3是%v\n\n", "Chinese", "Math", "English")
	fmt.Printf("课程1是%s\n课程2是%s\n课程3是%s\n", "Chinese", "Math", "English")
}
