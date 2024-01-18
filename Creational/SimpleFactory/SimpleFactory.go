package main

type Fruit interface {
	Show()
}

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

type Factory struct {
}

func (fac *Factory) CreateFruit(kind string) Fruit {
	var fruit Fruit
	if kind == "apple" {
		fruit = new(Apple)
	} else if kind == "banana" {
		fruit = new(Banana)
	}
	return fruit
}

// 实现了对象创建和使用的分离，过度依赖工厂类，增在了类的个数，用简单的参数代表了类的名称，修改工厂类，违背开闭原则
func main() {
	fac := Factory{}
	apple := fac.CreateFruit("apple")
	apple.Show()
	banana := fac.CreateFruit("banana")
	banana.Show()
}
