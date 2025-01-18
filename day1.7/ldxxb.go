/*
仙桃满庭春当代城一期65栋楼信息表
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
楼栋名称及编号：building nanme and number桃满庭春当代城一期65栋
楼栋总层数和高度：building floors and height 33层 100m
每一层户数及户型分布：floor distribution  每层2梯四户，左右两边是边户，中间是两中户
建筑年代：era  2013
建筑面积：building area  851328.92㎡
容积率：3.49
绿化率：40.5%
建筑类型：板搭结合
产权描述：普通住宅产权为70年
物业费：1.25元/㎡
区域：干河
物业公司：第一物业服务湖北有限公司仙桃分公司
停车位：地上地下约2800席
开发商：湖北万星置业有限公司
小区地址：building address 湖北省仙桃市黄金大道与学府路交汇处
小区配套设施：篮球场，乒乓球台，滑滑梯，各类健身器材，周边药店，早餐店，小吃店，生活超市，娱乐设施一应俱全
*/
package main

import "fmt"

func main() {
	fmt.Println("仙桃满庭春当代城一期65栋楼信息表")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	var a string = "仙桃满庭春当代城一期65栋"
	fmt.Printf("楼栋名称及编号：%s\n", a)

	var b int = 33
	fmt.Printf("楼层总数：%d层\n", b)

	var c int = 100
	fmt.Printf("楼层高度：%d米\n", c)

	var d string = "每层2梯四户,左右两边是边户,中间是两中户"
	fmt.Printf("每一层户型分布：%s\n", d)

	var e int = 2013
	fmt.Printf("建筑年代:%d\n", e)

	var f float64 = 851328.92
	fmt.Printf("建筑面积：%.2f㎡\n", f)

	var g float64 = 3.49
	fmt.Printf("容积率：%.2f\n", g)

	var a1 float64 = 40.5
	fmt.Printf("绿化率：%.1f％\n", a1)

	var b2 string = "板搭结合"
	fmt.Printf("建筑类型：%s\n", b2)

	var c2 string = "普通住宅产权为70年"
	fmt.Printf("产权描述：%s\n", c2)

	var d2 int = 2
	fmt.Printf("物业费：%d元/㎡\n", d2)

	var e2 string = "干河"
	fmt.Printf("区域：%s\n", e2)

	var f2 string = "第一物业服务湖北有限公司仙桃分公司"
	fmt.Printf("物业公司：%s\n", f2)

	var g2 string = "地上地下约2800席"
	fmt.Printf("停车位：%s席\n", g2)

	var a3 string = "湖北万星置业有限公司"
	fmt.Printf("开发商：%s\n", a3)

	var b3 string = "湖北省仙桃市黄金大道与学府路交汇处"
	fmt.Printf("小区地址：%s\n", b3)

	var c3 string = "篮球场，乒乓球台，滑滑梯，各类健身器材，周边药店，早餐店，小吃店，生活超市，娱乐设施一应俱全"
	fmt.Printf("小区配套设施：%s\n", c3)

}
