package decorator

import "fmt"

// IAuto 汽车接口
type IAuto interface {
	getPrice() int // 汽车报价
	showDetails()  // 报价明细
}

// Car1 IAuto的实现类
type Car1 struct{}

func (*Car1) getPrice() int {
	return 35000
}
func (*Car1) showDetails() {
	fmt.Print("Car1")
}

// Car2 IAuto的实现类
type Car2 struct{}

func (*Car2) getPrice() int {
	return 50000
}
func (*Car2) showDetails() {
	fmt.Print("Car2")
}

// Decorator 装饰类,它聚集一个IAuto接口
type Decorator struct {
	auto IAuto
}

func (d *Decorator) new(auto IAuto) *Decorator {
	d.auto = auto
	return d
}
func (*Decorator) getPrice() int {
	return 0
}
func (*Decorator) showDetails() {
	fmt.Print("Decorator")
}

// AirConditioner 继承了Decorator
type AirConditioner struct {
	Decorator
}

func (a *AirConditioner) new(auto IAuto) *AirConditioner {
	a.auto = auto
	return a
}
func (a *AirConditioner) showDetails() {
	a.auto.showDetails()
	fmt.Print(", AirConditioner")
}
func (a *AirConditioner) getPrice() int {
	return a.auto.getPrice() + 1500
}

// Navigation 继承了Decorator
type Navigation struct {
	Decorator
}

func (a *Navigation) new(auto IAuto) *Navigation {
	a.auto = auto
	return a
}
func (a *Navigation) showDetails() {
	a.auto.showDetails()
	fmt.Print(", Navigation")
}
func (a *Navigation) getPrice() int {
	return a.auto.getPrice() + 2500
}

// Airbag 继承了Decorator
type Airbag struct {
	Decorator
}

func (a *Airbag) new(auto IAuto) *Airbag {
	a.auto = auto
	return a
}
func (a *Airbag) showDetails() {
	a.auto.showDetails()
	fmt.Print(", Airbag")
}
func (a *Airbag) getPrice() int {
	return a.auto.getPrice() + 1000
}
