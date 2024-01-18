package main

type Fruit interface {
	Show()
}

type AbstractFactory interface {
	CreateFruit() Fruit
}

// -------------------------------------------------
type Apple struct {
	Fruit
}

func (a *Apple) Show() {
	println("apple")
}

type Banana struct {
	Fruit
}

func (b *Banana) Show() {
	println("banana")
}

// -------------------------------------------------
type AppleFactory struct {
	AbstractFactory
}

type BananaFactory struct {
	AbstractFactory
}

func (fac *AppleFactory) CreateFruit() Fruit {
	var fruit Fruit
	fruit = new(Apple)
	return fruit
}

func (fac *BananaFactory) CreateFruit() Fruit {
	var fruit Fruit
	fruit = new(Banana)
	return fruit
}

// 比起simplefactory增加了符合开闭原则
func main() {
	apple := new(AppleFactory).CreateFruit()
	apple.Show()
	banana := new(BananaFactory).CreateFruit()
	banana.Show()
}
