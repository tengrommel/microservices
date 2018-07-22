package main

import "fmt"

func main() {
	// 显示这个主菜单
	key := ""
	// 声明一个变量，控制是否退出for
	loop := true
	// 定义账户的余额 []
	balance := 10000.0
	// 每次收支的金额
	money := 0.0
	// 每次收支的说明
	note := ""
	// 定义一个变量，记录是否有收支的行为
	flag := false
	// 收支的一个详情使用字符串来记录
	// 当有收支时，只需要对detail进行处理
	details := "收支\t账户金额	\t收支金额\t说  明"
	// 显示这个主菜单
	for  {
		fmt.Println("\n------家庭收支记账软件-----------")
		fmt.Println("---------1 收支明细------------")
		fmt.Println("---------2 登记收入------------")
		fmt.Println("---------3 登记支出------------")
		fmt.Println("---------4 退出软件------------")
		fmt.Println("请选择（1-4）：")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("---------当前收支明细记录----------")
			if flag == true{
				fmt.Println(details)
			}else {
				fmt.Println("当前没有任何一笔收支")
			}
		case "2":
			fmt.Println("本次收入金额：")
			fmt.Scanln(&money)
			balance += money // 修改账户余额
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
			flag = true
		case "3":
			fmt.Println("本次支出金额：")
			fmt.Scanln(&money)
			// 这里需要必要的判断
			if money >balance {
				fmt.Println("支出的金额不足")
				break
			}
			balance -= money // 修改账户余额
			fmt.Println("本次支出说明：")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, note)
			flag = true
		case "4":
			fmt.Println("你确定要退出吗？y/n")
			choice := ""
			for {
				fmt.Scanln(&choice)
				if choice == "y" || choice == "n"{
					break
				}
				fmt.Println("你的输入有误，请重新输入y/n")
			}
			if choice == "y"{
				loop = false
			}
		default:
			fmt.Println("请输入正确的选项...")
		}
		if !loop{
			break
		}
	}
	fmt.Println("你退出家庭记账软件的使用...")
}
