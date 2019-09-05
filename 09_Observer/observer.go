package observer

import "fmt"

// IObserver 观察者接口
type IObserver interface {
	update(iObservable IObservable)
}

// IObservable 可观察者接口
type IObservable interface {
	registry(iObserver IObserver)
	remove(iObserver IObserver)
	getState() string
}

// Subscriber 观察者，实现了观察者接口
type Subscriber struct {
	name string
}

func (s *Subscriber) update(iObservable IObservable) {
	fmt.Printf("给%s的新消息！\n", s.name)
	fmt.Printf("消息：%s\n", iObservable.getState())
}

// Newsletter 可观察者，实现了可观察者接口
type Newsletter struct {
	observers []IObserver
	message   string
}

func (n *Newsletter) changeMessage(message string) {
	n.message = message
	n.notify()
}
func (n *Newsletter) remove(iObserver IObserver) {
	for i, observer := range n.observers {
		if observer == iObserver {
			n.observers = append(n.observers[:i], n.observers[i+1:]...)
			break
		}
	}
}
func (n *Newsletter) registry(iObserver IObserver) {
	n.observers = append(n.observers, iObserver)
}
func (n *Newsletter) getState() string {
	return n.message
}
func (n *Newsletter) notify() {
	for _, observer := range n.observers {
		observer.update(n)
	}
}
