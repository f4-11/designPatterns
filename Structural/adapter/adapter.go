package main

import "fmt"

// 通过adapter使用一个方法，其实是使用另一个类的方法
type V5 interface {
	USE5V()
}

type Phone struct {
	v V5
}

func NewPhone(v V5) *Phone {
	return &Phone{v}
}

func (p *Phone) Charge() {
	fmt.Println("phone is charging")
	p.v.USE5V()

}
func main() {

}
