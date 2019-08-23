package bridge

import "testing"

func TestTest(t *testing.T) {
	var abstractDef AbstractDef
	abstractDef.impl = new(ConcreteImplementerA)

	// 具体实现对于客户是隐藏的
	abstractDef.operation()

	// 在运行时可以随时替换具体实现
	abstractDef.impl = new(ConcreteImplementerB)
	abstractDef.operation()

	var specDef SpecAbstractDef
	specDef.impl = new(ConcreteImplementerB)
	specDef.operation()
}
