/**
	学习控制结构-switch 结构
 */
package main

import "fmt"

func main() {
	var num1 int = 7

	switch {
	case num1 < 0:
		fmt.Println("Number is negative")
	case num1 > 0 && num1 < 10:
		fmt.Println("Number is between 0 and 10")
	default:
		fmt.Println("Number is 10 or greater")
	}
	// 输出：Number is between 0 and 10
	var num int = 10
	switch num {
		case 11:
			fmt.Println("ok1")
		case 10:
			fmt.Println("ok2 2")
			fallthrough
		case 32:
			fmt.Println("ok3")
		default:
			fmt.Println("没有匹配到")
	}

	var num2 int
	fmt.Println("请输入一个数字")
	fmt.Scanf("%d", &num2)
	switch num2 {
	case 100,99:
		fmt.Println("100,99:")
	case 98:
		fmt.Println("98")
	default:
		fmt.Println("others")
	}
}
/*
	总结：
	1、前花括号 { 必须和 switch 关键字在同一行。
	2、switch的执行的流程是，先执行表达式，得到值，然后和case的表达式进行比较，如果相等， 就匹配到，然后执行对应的 case 的语句块，然后退出 switch 控制。
	3、如果switch的表达式的值没有和任何的case的表达式匹配成功，则执行default的语句块。执行 后退出 switch 的控制.
	4、golang的case后的表达式可以有多个，使用逗号间隔.
	5、golang中的case语句块不需要写break, 因为默认会有,即在默认情况下，当程序执行完 case 语 句块后，就直接退出该 switch 控制结构。
	6、case/switch 后是一个表达式( 即:常量值、变量、一个有返回值的函数等都可以)
	7、case 后的各个表达式的值的数据类型，必须和 switch 的表达式数据类型一致
	8、case后面的表达式如果是常量值(字面量)，则要求不能重复
	9、default 语句不是必须的.
	10、switch 后也可以不带表达式（实际上默认为判断是否为 true），类似 if --else 分支来使用。
	11、switch 后也可以直接声明/定义一个变量，分号结束，不推荐。
	12、switch 穿透-fallthrough ，如果在 case 语句块后增加 fallthrough ,则强制会继续执行下一个 case（不管判断），也 叫 switch 穿透
 */



