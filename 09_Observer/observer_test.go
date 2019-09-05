package observer

import (
	"fmt"
	"testing"
)

func TestTest(t *testing.T) {
	// 新建可观察对象
	newsletter := &Newsletter{observers: make([]IObserver, 0)}
	// 新建观察对象
	s1 := &Subscriber{name: "s1"}
	s2 := &Subscriber{name: "s2"}
	// 注册对象
	newsletter.registry(s1)
	newsletter.registry(s2)
	// 可观察对象更新状态
	newsletter.changeMessage("状态1")
	fmt.Printf("------------\n\n")
	newsletter.changeMessage("状态2")
	fmt.Printf("------------\n\n")

	// 移除可观察对象s1
	newsletter.remove(s1)
	newsletter.changeMessage("状态3")
}
