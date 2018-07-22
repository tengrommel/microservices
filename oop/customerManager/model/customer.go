package model

import "fmt"

// 声明一个Customer结构体，表示一个客户信息
type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}


// 使用编写一个工厂模式，返回一个Customer的实例
func NewCustomer(name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		Name:name,
		Gender:gender,
		Age:age,
		Phone:phone,
		Email:email,
	}
}

// 返回用户的信息
func (c Customer)GetInfo() string {
	info :=fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v",
		c.Id, c.Name, c.Gender, c.Age, c.Phone, c.Email)
	return info
}
