package service

import "microservices/oop/customerManager/model"

// 该CustomerService,完成对Customer的操作，包括增删改查
type CustomerService struct {
	customers []model.Customer
	// 声明一个字段，表示当前切片含有多少个客户
	customerNum int
}

// 编写一个方法，可以返回 *CustomerService
func NewCustomerService() *CustomerService {
	// 为了能够看到有客户在切片中，我们初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer("张三","男", 20,"112", "zs@souhu.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

// 返回客户切片
func (c *CustomerService)List() []model.Customer {
	return c.customers
}

func (c *CustomerService)Add(customer model.Customer) bool {
	c.customerNum++
	customer.Id = c.customerNum
	c.customers = append(c.customers, customer)
	return true
}