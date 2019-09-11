package strategy

import "testing"

func TestTest(t *testing.T) {
	context := &Context{}

	context.istrategy = new(StrategyA)
	context.operation()

	context.istrategy = new(StrategyB)
	context.operation()
}
