package adapter

import (
	"testing"
)

func TestClass(t *testing.T) {
	// 通过目标接口targetInterface的operationA()方法,
	// 最终调用someInterface(targetClass目标类实现的接口)的operationB()方法
	var targetInterface targetInterface
	targetInterface = new(Adapter)
	targetInterface.operationA()
}
