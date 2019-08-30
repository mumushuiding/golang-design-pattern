package proxy

// IRealClass 实际类和代理要实现的接口
type IRealClass interface {
	getContent() string
}

// RealClass 实现类
type RealClass struct {
}

func (r *RealClass) getContent() string {
	return "Data from realClass"
}

// Proxy 代理
type Proxy struct {
	realClass IRealClass // 通过实际类的接口关联实际类
}

func (p *Proxy) getContent() string {
	if p.realClass == nil {
		p.realClass = new(RealClass) // 生成实际类
	}
	return p.realClass.getContent()
}
