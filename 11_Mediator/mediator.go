package mediator

import "fmt"

// IMediator 中间者接口
type IMediator interface {
	modifyOccurred(ic IColleague)
}

// IColleague 同事接口
type IColleague interface {
	modify()
}

// Colleague 同事
type Colleague struct {
	mediator IMediator
}

// func (c *Colleague) modify() {
// 	c.mediator.modifyOccurred(c)
// }

// ColleagueA 同事A
type ColleagueA struct {
	Colleague
}

func (c *ColleagueA) methodA() {
	fmt.Println("由于同事B的变化，同事A执行了 methodA")
	fmt.Println("---------------------")
	fmt.Println()
}
func (c *ColleagueA) modify() {
	fmt.Println("更改同事A")
	c.mediator.modifyOccurred(c)
}

// ColleagueB 同事B
type ColleagueB struct {
	Colleague
}

func (c *ColleagueB) methodB() {
	fmt.Println("由于同事A的修改，同事B执行了 methodB")
	fmt.Println("---------------------")
	fmt.Println()
}
func (c *ColleagueB) modify() {
	fmt.Println("更改同事B")
	c.mediator.modifyOccurred(c)
}

// mediatorA 具体中间者A
type mediatorA struct {
	ca *ColleagueA
	cb *ColleagueB
}

func (m *mediatorA) modifyOccurred(ic IColleague) {
	fmt.Println("通知中间者")
	switch ic.(type) {
	case *ColleagueA:
		m.cb.methodB()
	case *ColleagueB:
		m.ca.methodA()
	}
}
