/*
影片《误杀3》的基本信息
------
上映时间：2024年12月28日在中国上映Release time
片长：109分钟length of a film
类型：剧情，悬疑，犯罪Film Genre
主演：肖央，佟丽娅，段奕宏，刘雅瑟，王龙正star
票价：39元fare
影片评价：观众评价普遍较高，影片剧情紧凑，反转不断Film review

影片《彷徨之刃》的基本信息
------
上映时间：2024年5月17日在中国上映
片长：107分钟。
类型：剧情，犯罪
主演：王千源，齐溪，阿如那吴双等
票价：20元
影片评价：无
*/
package main

import "fmt"

func main() {
	fmt.Println("影片《误杀3》的基本信息")
	fmt.Println("------")

	{
		var Releasetime string = "2024年12月28日在中国上映"
		var lengthofafilm int = 109
		var FilmGenre string = "剧情，悬疑，犯罪"
		var star string = "肖央，佟丽娅，段奕宏，刘雅瑟，王龙正"
		var fare int = 39
		var Filmreview string = "观众评价普遍较高，影片剧情紧凑，反转不断"
		// fmt.Println("上映时间：",Releasetime ,"\n","片长：",lengthofafilm,"分钟","\n","类型：",FilmGenre,"\n","主演：",star,"\n","票价：",fare,"元","\n","影片评价：",Filmreview,"\n" ,"\n")}
		//fmt.Printf("影片《误杀3》的基本信息\n------\n上映时间：%s\n片长：%d分钟\n类型：%s\n主演：%s\n票价：%d元\n影片评价：%s\n\n",Releasetime ,lengthofafilm ,FilmGenre ,star,fare ,Filmreview )

		fmt.Println("上映时间：", Releasetime)
		fmt.Println("片长：", lengthofafilm)
		fmt.Println("类型：", FilmGenre)
		fmt.Println("主演：", star)
		fmt.Println("票价：", fare)
		fmt.Println("影片评价：", Filmreview, "\n")
	}

	fmt.Println("影片《彷徨之刃》的基本信息")
	fmt.Println("------")

	{
		var Releasetime string = "2024年5月17日在中国上映"
		var lengthofafilm int = 107
		var FilmGenre string = "剧情，犯罪"
		var star string = "王千源，齐溪，阿如那,吴双等"
		var fare int = 20
		var Filmreview string = "无"
		//fmt.Printf("影片《彷徨之刃》的基本信息\n------\n上映时间：%s\n片长：%d分钟\n类型：%s\n主演：%s\n票价：%d元\n影片评价：%s\n",Releasetime,lengthofafilm,FilmGenre ,star,fare,Filmreview )
		// fmt.Println("上映时间：",Releasetime ,"\n","片长：",lengthofafilm,"分钟","\n","类型：",FilmGenre,"\n","主演：",star,"\n","票价：",fare,"元","\n","影片评价：",Filmreview )

		fmt.Println("上映时间：", Releasetime)
		fmt.Println("片长：", lengthofafilm,"分钟")
		fmt.Println("类型：", FilmGenre)
		fmt.Println("主演：", star)
		fmt.Println("票价：", fare,"元")
		fmt.Println("影片评价：", Filmreview)
	}

}
