package main

import "fmt"

type Phone interface {
	Show()
}

type Decorator struct {
	phone Phone
}

func (d *Decorator) Show() {}

type Huawei struct{}

func (hw *Huawei) Show() {
	fmt.Println("Huawei")
}

type Xiaomi struct{}

func (xm *Xiaomi) Show() {
	fmt.Println("Xiaomi")
}

type Decorator1 struct {
	Decorator
}

func (d1 *Decorator1) Show() {
	d1.phone.Show()
	fmt.Println("d1")
}

func NewDecorator1(phone Phone) Phone {
	return &Decorator1{Decorator{phone}}
}

type Decorator2 struct {
	Decorator
}

func (d2 *Decorator2) Show() {
	d2.phone.Show()
	fmt.Println("d2")
}

func NewDecorator2(phone Phone) Phone {
	return &Decorator2{Decorator{phone}}
}

func main() {
	huawei := new(Huawei)
	huawei.Show()
	fmt.Println("...........")
	huawei1 := NewDecorator1(huawei)
	huawei1.Show()
	fmt.Println(".............................")
	huawei2 := NewDecorator2(huawei1)
	huawei2.Show()
}

//和代理模式区别，更加动态灵活，难排错，小对象多性能差
