package main

import "fmt"

type A struct {
	Name string
	age int
}

type B struct {
	Name string
	Score float64
}

type C struct {
	A
	B
	Name string
}

type D struct {
	a A
}

type Goods struct {
	Name string
	Price float64
}

type Brand struct {
	Name string
	Address string
}

type TV struct {
	*Goods
	*Brand
}

func main() {
	var c C
	c.Name = "tom"
	var d D
	d.a.Name = "jack"

	tv := TV{&Goods{"电视机001", 5000.99}, &Brand{"海尔", "山东青岛"},}
	tv2 := TV{
		&Goods{
			Name: "电视机002",
			Price:5000.99,
		},
		&Brand{
			Name:"夏普",
			Address:"东京",
		},
	}
	fmt.Println("TV", tv)
	fmt.Println("TV2", tv2)
}
