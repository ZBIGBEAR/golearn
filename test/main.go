//当前程序的包名
package main

//导入其他的包
import "fmt"

//同时导入多个包
/*import{
	"fmt",
	"io"
}*/

//给包起别名,为fmt起别名为fmt2
//import fmt2 "fmt"

//省略调用
// 调用的时候只需要Println()，而不需要fmt.Println()
//import . "fmt"

//常量定义
const PI = 3.14

//全局变量的声明和赋值
var name = "alex"

//一般变量的声明
type i int

//结构体的声明
type alex struct{}

//接口的声明
type alexInter interface{}

//main函数作为程序的入口
func main(){
	fmt.Println("hello,alex")

	//定义两个局部变量
	var a = 1.5
	var b = 2
	fmt.Println(a,b)
	test()
}


//在 Go 中，布尔值的类型为 bool，值是 true 或 false，默认为 false。
var isActive bool //全局变量声明

var enabled,disabled = true,false  //忽略类型的声明，直接赋值。go会自动判断数据类型

func test(){
	var available bool //一般声明
	valid := false //简短声明。声明的时候赋值，不需要指定var关键字和数据类型
	available = false //赋值

	fmt.Println(valid,available)
}
/*
总结：
1.通过const关键字来进行常量的定义
2.通过在函数外部使用var关键字来进行全局变量的声明和赋值
3.通过type关键字来进行结构体（struct）和接口（interface）的声明
4.通过func关键字来进行函数的声明
5.go语言中，使用大小写来决定常量、变量、结构体、接口、函数以及自定义数据类型是否可被外部包所调用
小写为不可调用
大写为可调研
*/
