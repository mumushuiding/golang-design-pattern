package facade

import "testing"

func TestTest(t *testing.T) {
	facade := NewFacade()
	// 添加库存
	facade.add(10)
	// 减少库存
	facade.remove(5)
}
