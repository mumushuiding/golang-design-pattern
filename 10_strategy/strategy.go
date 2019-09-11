package strategy

import "fmt"

// IStrategy 策略接口
type IStrategy interface {
	algorithm()
}

// Context 环境
type Context struct {
	istrategy IStrategy
}

func (c *Context) operation() {
	c.istrategy.algorithm()
}

// StrategyA 策略A
type StrategyA struct{}

func (s *StrategyA) algorithm() {
	fmt.Println("执行了策略A")
}

type StrategyB struct {
}

func (s *StrategyB) algorithm() {
	fmt.Println("执行了策略B")
}
