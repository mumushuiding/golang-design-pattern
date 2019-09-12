package mediator

import (
	"testing"
)

func TestTest(t *testing.T) {
	// 初始化中间者
	mediator := new(mediatorA)
	// 同事A
	ca := &ColleagueA{}
	ca.mediator = mediator

	mediator.ca = ca

	// 同事B
	cb := &ColleagueB{}
	cb.mediator = mediator

	mediator.cb = cb

	// 更改同事A
	ca.modify()
	// 更改同事B
	cb.modify()
}
