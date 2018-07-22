package utils

import "fmt"

type FamilyAccount struct {
	// 声明必须的字段，
	key string
	// 声明一个变量，控制是否退出for
	loop bool
	// 定义账户的余额 []
	balance float64
	// 每次收支的金额
	money float64
	// 每次收支的说明
	note string
	// 定义一个变量，记录是否有收支的行为
	flag bool
	// 当有收支时，只需要对details进行拼接处理即可
	details string
}

// 编写工厂模式的构造方法
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key: "",
		loop:true,
		balance: 10000.0,
		money:0.0,
		note:"",
		flag:false,
		details: "收支\t账户金额	\t收支金额\t说  明",
	}
}

// 将显示明细写成一个方法
func (f *FamilyAccount)showDetails()  {
	fmt.Println("---------当前收支明细记录----------")
	if f.flag == true{
		fmt.Println(f.details)
	}else {
		fmt.Println("当前没有任何一笔收支")
	}
}

// 将登记收入写成一个方法，和*FamilyAccount绑定
func (f *FamilyAccount)income()  {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&f.money)
	f.balance += f.money // 修改账户余额
	fmt.Println("本次收入说明：")
	fmt.Scanln(&f.note)
	f.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", f.balance, f.money, f.note)
	f.flag = true
}

func (f *FamilyAccount)pay()  {
	fmt.Println("本次支出金额：")
	fmt.Scanln(&f.money)
	// 这里需要必要的判断
	if f.money >f.balance {
		fmt.Println("支出的金额不足")
		//break
	}
	f.balance -= f.money // 修改账户余额
	fmt.Println("本次支出说明：")
	fmt.Scanln(&f.note)
	f.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", f.balance, f.money, f.note)
	f.flag = true
}

func (f *FamilyAccount)exit()  {
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
		f.loop = false
	}
}

// 给该结构体绑定相应的方法
func (f *FamilyAccount)MainMenu()  {
	for  {
		fmt.Println("\n------家庭收支记账软件-----------")
		fmt.Println("---------1 收支明细------------")
		fmt.Println("---------2 登记收入------------")
		fmt.Println("---------3 登记支出------------")
		fmt.Println("---------4 退出软件------------")
		fmt.Println("请选择（1-4）：")
		fmt.Scanln(&f.key)
		switch f.key {
		case "1":
			f.showDetails()
		case "2":
			fmt.Println("ok")
			f.income()
		case "3":
			f.pay()
		case "4":
			f.exit()
		default:
			fmt.Println("请输入正确的选项...")
		}
		if !f.loop{
			break
		}
	}
}