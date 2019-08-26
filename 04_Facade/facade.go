package facade

import "fmt"

// AppleWarehouse 苹果仓库
type AppleWarehouse struct {
	num int // 库存量
}

func (a *AppleWarehouse) add(number int) {
	a.num += number
	fmt.Printf("苹果库存添加了：%d个\n", number)
}
func (a *AppleWarehouse) remove(number int) {
	a.num -= number
	fmt.Printf("苹果库存减少了：%d个\n", number)
}

// BananaWarehouse 香蕉仓库
type BananaWarehouse struct {
	num int
}

func (b *BananaWarehouse) add(number int) {
	b.num += number
	fmt.Printf("香蕉库存添加了：%d个\n", number)
}
func (b *BananaWarehouse) remove(number int) {
	b.num -= number
	fmt.Printf("香蕉库存减少了：%d个\n", number)
}

// PearWarehouse 梨子仓库
type PearWarehouse struct {
	num int
}

func (p *PearWarehouse) add(number int) {
	p.num += number
	fmt.Printf("梨子库存添加了：%d个\n", number)
}
func (p *PearWarehouse) remove(number int) {
	p.num -= number
	fmt.Printf("梨子库存减少了：%d个\n", number)
}

// Facade 仓库外观类
type Facade struct {
	apple  *AppleWarehouse
	banana *BananaWarehouse
	pear   *PearWarehouse
}

// NewFacade 实例化外观类
func NewFacade() *Facade {
	var facade Facade
	facade.apple = new(AppleWarehouse)
	facade.banana = new(BananaWarehouse)
	facade.pear = new(PearWarehouse)
	return &facade
}
func (f *Facade) add(number int) {
	fmt.Println("为所有仓库添加库存")
	f.apple.add(number)
	f.banana.add(number)
	f.pear.add(number)
	fmt.Println()
}
func (f *Facade) remove(number int) {
	fmt.Println("所有仓库减少库存")
	f.apple.remove(number)
	f.banana.remove(number)
	f.pear.add(number)
}
