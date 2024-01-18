package main

import "fmt"

type AbstractApple interface {
	ShowApple()
}

type AbstractBanana interface {
	ShowBanana()
}

type AbstractFactory interface {
	CreateApple() AbstractApple
	CreateBanana() AbstractBanana
}

// ----------------------------------------------------
type ChinaApple struct {
}

func (ca *ChinaApple) ShowApple() {
	fmt.Println("china apple")
}

type ChineBanana struct {
}

func (cb *ChineBanana) ShowBanana() {
	fmt.Println("china banana")
}

// -----------------------------------------------------
type ChinaFactory struct {
}

func (cf *ChinaFactory) CreateApple() AbstractApple {
	apple := new(ChinaApple)
	return apple
}

func (cf *ChinaFactory) CreateBanana() AbstractBanana {
	banana := new(ChineBanana)
	return banana
}

func main() {
	cfac := new(ChinaFactory)
	cfac.CreateApple().ShowApple()
	cfac.CreateBanana().ShowBanana()
}
