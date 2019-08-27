package composite

import (
	"fmt"
)

// INode 节点接口
type INode interface {
	operation()
	addNode(node INode)
	remove(node INode)
	getName() string
	setLevel(level int)
}

// Node 节点,定义了派生类的基本接口和行为
type Node struct {
	name     string
	level    int
	children []INode
}

func (*Node) operation() {}
func (*Node) addNode(node INode) {
	fmt.Println("添加节点")
}
func (*Node) remove(node INode) {
	fmt.Println("删除节点")
}
func (n *Node) getName() string {
	return n.name
}
func (n *Node) setLevel(level int) {
	n.level = level
}

// -------------------------------------

// Compound 组合节点
type Compound struct {
	Node // 继承了Node的对象所有属性和方法
}

func (c *Compound) operation() {
	// 输出空格
	format := "%" + fmt.Sprintf("%d", c.level*2) + "s"
	fmt.Printf(format, "")
	fmt.Printf("+ %s\n", c.getName())
	// 调用所有的子元素
	// 计算新的输出导数
	c.level++
	for _, n := range c.children {
		n.setLevel(c.level)
		n.operation()
	}
	// 输出层数重置
	c.level--
}
func (c *Compound) addNode(node INode) {
	c.children = append(c.children, node)
}

// remove 属于简化写法
func (c *Compound) remove(node INode) {
	out := c.children[:0]
	for _, n := range c.children {
		if n != node {
			out = append(out, n)
		}
	}
	c.children = out
}

// ------------------------

// Leaf 叶节点,没有子节点
type Leaf struct {
	Node
}

func (l *Leaf) operation() {
	// 输出空格
	format := "%" + fmt.Sprintf("%d", l.level*2) + "s"
	fmt.Printf(format, "")
	fmt.Printf("- %s\n", l.getName())
}
