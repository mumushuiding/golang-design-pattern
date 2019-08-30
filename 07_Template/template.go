package template

import (
	"fmt"
)

// ICard 可变操作接口
type ICard interface {
	someWriting()
}

// holidayCard 基类
type holidayCard struct {
	ICard // 通过聚集接口实现泛型
}

// cardWriting 是模板方法
func (h *holidayCard) cardWriting() {
	h.someWriting()
	h.textWriting()
}

// textWriting 具体操作
func (*holidayCard) textWriting() {
	fmt.Println("顺利到达目的地")
}

// holidayCardToFriend 基类派生类
type holidayCardToFriend struct {
	holidayCard // 继承基类
}

// someWriting 覆写了基类方法
func (h *holidayCardToFriend) someWriting() {
	fmt.Print("给朋友：")
}

// holidayCardToCompany 基类派生类
type holidayCardToCompany struct {
	holidayCard // 继承基类
}

// someWriting 覆写了基类方法
func (h *holidayCardToCompany) someWriting() {
	fmt.Print("给公司：")
}
