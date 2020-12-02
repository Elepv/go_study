package adapter

//Target 适配的目标接口
type Target interface {
	Request() string
}

//Adaptee 被适配的目标接口
type Adaptee interface {
	SpecificRequest() string
}

//NewAdaptee 是被适配接口的工厂函数
func NewAdaptee() Adaptee {
	return &AdapteeImpl{}
}

//AdapteeImpl 是被适配的目标类
type AdapteeImpl struct{}

//SpecificRequest 是一个目标类的方法
func (*AdapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

//NewAdapter 是一个Adapter的工厂函数
func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}

//Adapter 是转换Adaptee为Target接口的适配器
type adapter struct {
	Adaptee
}

//Request 实现Target接口
func (a *adapter) Request() string {
	return a.SpecificRequest()
}
