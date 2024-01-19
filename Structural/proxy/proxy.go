package main

import (
	"fmt"
	"reflect"
)

type Goods struct {
	Kind string
	Fact bool
}

type Shopping interface {
	Buy(goods *Goods)
}

type KoreaShopping struct{}

func (ks *KoreaShopping) Buy(goods *Goods) {
	fmt.Println(goods.Kind)
}

type ChinaShopping struct{}

func (cs *ChinaShopping) Buy(goods *Goods) {
	fmt.Println(goods.Kind)
}

// 我们就在这里抢别人家代码用
type OverseaProxy struct {
	shopping Shopping
}

func NewProxy(shopping Shopping) Shopping {
	return &OverseaProxy{shopping}
}

func (op *OverseaProxy) Buy(goods *Goods) {
	op.Wwww()
	if goods.Fact == true {
		op.shopping.Buy(goods)
	}
}

func (op *OverseaProxy) Wwww() {
	fmt.Println(reflect.TypeOf(op))
}

func main() {
	g1 := Goods{"g1", true}
	g2 := Goods{"g2", false}
	shopping1 := new(KoreaShopping)
	kp := NewProxy(shopping1)
	kp.Buy(&g1)
	shopping2 := new(ChinaShopping)
	cp := NewProxy(shopping2)
	cp.Buy(&g2)
}
