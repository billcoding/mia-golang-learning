 /*先进市2024年转来5名新生，需要把这5名学生分进 超满意小学，超幸福小学，超稳定小学，最后将所分人员打印出来
5名学生信心如下
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
姓名     性别     年龄   成绩
李明      男        8     70
王莉丝    女        11    96
杜一茗    男        10    89
戴斯      女        7     70
雷鸣      男        13    60

分组情况：
超满意小学：年龄大于7 性别女 成绩大于60
超幸福小学：年龄在7到13之间，性别男 成绩在60到90之间
超稳定小学：其他同学进超稳定小学

打印条件：
超满意小学有：姓名1、姓名2
超幸福小学有：姓名1、姓名2
超稳定小学有：姓名1、姓名2
 */

 package main
 import "fmt"
 func main(){

   var(
      schoolcmy,schoolcxf,schoolcwd  = "超满意小学","超幸福小学","超稳定小学"
      membercmy,membercxf,membercwd  = "","",""

       name1,gender1 string = "李明","男"
       age1,score1 int  =  8, 70

       name2,gender2 string = "王莉丝","女"
       age2,score2 int  =  11,96

       name3,gender3 string = "杜一茗","男"
       age3,score3 int  =  10,89

        name4,gender4 string = "戴斯","女"
        age4,score4 int  =  7,70

        name5,gender5 string = "雷鸣","男"
        age5,score5 int  =  13,60
       )

      if(age1>7&&gender1=="女"&&score1>60){
        membercmy = name1

      }else if(age1>7&&age1<=13&&gender1=="女"&&score1>60&&score1<90){

        membercxf = name1
      }else{
         membercwd = name1
      }


      if(age2>12&&gender2=="男"&&score2>70){
        membercmy += ","+name2

      }else if(age2>7&&age2<=13&&gender2=="女"&&score2>60&&score2<90){

        membercxf += ","+name2
      }else{
         membercwd += ","+name2
      }

      if(age2>12&&gender2=="男"&&score2>70){
        membercmy = name3

      }else if(age3>7&&age3<=13&&gender3=="女"&&score3>60&&score3<90){

        membercxf += ","+name3
      }else{
         membercwd += ","+name3
      }

      if(age2>12&&gender2=="男"&&score2>70){
        membercmy += ","+name4

      }else if(age4>7&&age4<=13&&gender4=="女"&&score4>60&&score4<90){

        membercxf += ","+name4
      }else{
         membercwd += ","+name4
      }


      if(age2>12&&gender2=="男"&&score2>70){
        membercmy += ","+name5

      }else if(age5>7&&age5<=13&&gender5=="女"&&score5>60&&score5<90){

        membercxf += ","+name5
      }else{
         membercwd += ","+name5
      }

      fmt.Printf("%s有：%s\n", schoolcmy, membercmy)
      fmt.Printf("%s有：%s\n", schoolcxf, membercxf)
      fmt.Printf("%s有：%s\n", schoolcwd, membercwd)




     }



