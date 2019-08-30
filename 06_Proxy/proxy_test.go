package proxy

import (
	"fmt"
	"testing"
)

func TestTest(t *testing.T) {
	proxy := &Proxy{}
	fmt.Println(proxy.getContent())

}
