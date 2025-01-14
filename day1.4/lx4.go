package main
import "fmt"
func main (){
	var class1 string = "Chinese"
	var class2 string = "Math"
	var class3 string = "English"
	fmt.Println("课程1是:",class1,"课程2是:",class2,"课程3是:",class3)
	fmt.Println("课程1是:",class1,"\n","课程2是:",class2,"\n","课程3是:",class3,"\n")
	fmt.Printf("课程1是%v,\n课程2是%v,\n课程3是%v,\n","Chinese","Math","English")
}