package main

import (
	"fmt"
	"microservices/oop/customerManager/service"
	"microservices/oop/customerManager/model"
)

type customerView struct {
	// 定义必要字段
	key string // 接收用户输入...
	loop bool // 表示是否循环的显示
	// 增加一个字段customerService
	customerService *service.CustomerService
}

func (c *customerView)list()  {
	// 首先，获取到当前所有的客户信息（在切片中）
	customers := c.customerService.List()
	// 显示
	fmt.Println("----------------客户列表------------------")
	fmt.Println("编号\t名字\t性别\t年龄\t电话\t邮箱")
	for i:=0;i<len(customers);i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("--------------客户列表完成-----------------")
}

func (c *customerView)add()  {
	fmt.Println("----------添加客户------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer(name, gender, age,phone, email)
	if c.customerService.Add(customer){
		fmt.Println("-----------添加完成--------------")
	}else {
		fmt.Println("-----------添加失败--------------")
	}
}

// 显示主菜单
func (c *customerView)mainMenu() {
	for  {
		fmt.Println("-------------客户信息管理软件--------------")
		fmt.Println("				1 添加客户					")
		fmt.Println("				2 修改客户					")
		fmt.Println("				3 删除客户					")
		fmt.Println("				4 客户列表					")
		fmt.Println("				5 退  出					    ")
		fmt.Println("请选择(1-5):")
		fmt.Scanln(&c.key)
		switch c.key {
		case "1":
			fmt.Println("添 加 客 户")
		case "2":
			fmt.Println("修 改 客 户")
		case "3":
			fmt.Println("删 除 客 户")
		case "4":
			c.list()
		case "5":
			c.loop = false
		default:
			fmt.Println("你的输入有误，请重新输入...")
		}
		if !c.loop{
			break
		}
	}
	fmt.Println("你退出了客户管理系统")
}

func main() {
	// 在主函数中，创建一个customerView，并运行显示主菜单...
	customerView := customerView{
		key: "",
		loop: true,
	}
	// 显示主菜单...
	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()
}
