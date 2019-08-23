package adapter

import "fmt"

// targetInterface 目标接口,能满足客户需求
type targetInterface interface {
	operationA()
}

// someInterface 这是目标类targetClass实现的接口
type someInterface interface {
	operationB()
}

// targetClass 目标对象,需要被匹配的类
type targetClass struct{}

// operationB 目标对象的方法
func (*targetClass) operationB() {
	fmt.Println("调用到 targetClass的方法 operationB 成功!!")
}

// Adapter 适配器
type Adapter struct {
	targetClass // 继承类 targetClass ,类适配器写法
	// targetClass targetClass // 聚合类 targetClass, 对象适配器写法
}

// operationA 实现targetInterface方法
func (a *Adapter) operationA() {
	fmt.Println("调用接口 operationA")
	a.operationB() // 类适配器写法
	// a.targetClass.operationB() // 对象适配器写法
}
