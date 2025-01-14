// 一个Go的结构，就是怎么个写法

// 首先声明主（main）包
package main

// 打印hello，world
// 1. 首先导入fmt包
import "fmt"

// 再声明主（main）函数
func main() {
	//2.调用fmt包里的Println函数
	fmt.Println("hello,world")
}

//运行
//go run lx.go
