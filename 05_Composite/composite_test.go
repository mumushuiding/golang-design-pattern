package composite

import (
	"fmt"
	"testing"
)

func TestTest(t *testing.T) {
	fmt.Printf("-------------------\n\n")
	c := &Compound{}
	c.name = "组合节点第1层"
	c11 := &Compound{}
	c11.name = "组合节点第11层"
	c12 := &Compound{}
	c12.name = "组合节点第12层"
	c.addNode(c11)
	c.addNode(c12)

	leaf111 := &Leaf{}
	leaf111.name = "leaf 111"
	leaf112 := &Leaf{}
	leaf112.name = "leaf 112"
	c11.addNode(leaf111)
	c11.addNode(leaf112)

	leaf121 := &Leaf{}
	leaf121.name = "leaf 121"
	c12.addNode(leaf121)

	c.operation()

	fmt.Printf("-------------------\n\n")
}
