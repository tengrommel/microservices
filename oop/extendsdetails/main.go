package main

import "fmt"

type A struct {
	Name string
	age int
}

func (a *A)SayOk()  {
	fmt.Println("A SayOk", a.Name)
}

func (a *A)hello()  {
	fmt.Println("A hello", a.Name)
}

type B struct {
	A
	Name string
}

func (b *B)SayOk()  {
	fmt.Println("B SayOk", b.Name)
}

func main() {
	//var b B
	//b.A.Name = "tom"
	//b.A.SayOk()
	//b.A.hello()
	// 上面的写法可以简化
	//b.Name = "smith"
	//b.age = 20
	//b.hello()
	var b B
	b.Name = "jack"
	b.age = 100
	b.SayOk()
	b.hello()
}
