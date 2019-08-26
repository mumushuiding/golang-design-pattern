package decorator

import (
	"fmt"
	"testing"
)

func TestTest(t *testing.T) {
	// 实例化Car1
	var auto IAuto
	auto = new(Car1)
	// 汽车添加空调
	dec := &AirConditioner{}
	auto = dec.new(auto)
	auto.showDetails()
	fmt.Printf("\ntotal: %d\n\n", auto.getPrice())

	// 添加导航
	dec1 := &Navigation{}
	auto = dec1.new(auto)
	auto.showDetails()
	fmt.Printf("\ntotal: %d\n\n", auto.getPrice())

	// 添加安全气囊
	dec2 := &Airbag{}
	auto = dec2.new(auto)
	auto.showDetails()
	fmt.Printf("\ntotal: %d\n\n", auto.getPrice())

}
