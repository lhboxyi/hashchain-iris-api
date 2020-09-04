package demo

/**
1）default定义在其他case的前面或者后面没有任何影响。
2）如果case1和case2都满足条件，谁排在前面执行谁
3）没有default时，如果所有case都不满足条件，则switch case直接跳出，什么都不执行。即switch case结构本身没有循环执行能力。
4）可以使用for+switch结构来让switch循环执行，但是switch里面的break没有任何意义。可以用break+循环名称来指定break哪种循环结构
5）switch后面可以不带表达式，也可以带表达式(或者变量)，对应的case 表达式也不一样，但是落脚到case最终的结构，只有true或者false两种情况。
 */
import "fmt"

func SwitchCaseDemo1() { //===1===
	switch 1 {
	case 1:
		fmt.Println("===1===")
	case 2:
		fmt.Println("===2===")
	default:
		fmt.Println("===0===")
	}
}

func SwitchCaseDemo2() { //===2===
	switch 1 {
	case 2:
		fmt.Println("===2===")
	case 1:
		fmt.Println("===1===")
	}
}
