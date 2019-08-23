package bridge

import "fmt"

// AbstractDef 抽象定义
type AbstractDef struct {
	impl IImplementer // 抽象定义通过聚合接口IImplementer,同具体实现连接
}

func (ad *AbstractDef) operation() {
	fmt.Println("抽象定义 operation")
	ad.impl.operationImpl()
}

// SpecAbstractDef 细分的抽象定义
type SpecAbstractDef struct {
	AbstractDef
}

// IImplementer AbstractDef聚合的接口
type IImplementer interface {
	operationImpl()
}

// ConcreteImplementerA 具体实现A
type ConcreteImplementerA struct {
}

// operationImpl 是IImplementer中operationImpl()方法的具体实现
func (c *ConcreteImplementerA) operationImpl() {
	fmt.Println("具体实现 A")
}

// ConcreteImplementerB 具体实现B
type ConcreteImplementerB struct{}

// operationImpl 是IImplementer中operationImpl()方法的具体实现
func (c *ConcreteImplementerB) operationImpl() {
	fmt.Println("具体实现 B")
}
